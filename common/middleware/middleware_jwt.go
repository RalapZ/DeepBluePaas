package middleware

import (
	"github.com/RalapZ/DeepBluePaas/common/http/dpresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/global"

	//"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var (
	jwtService = JwtService{}
	JwtRedis *redis.Client
	JwtDB *gorm.DB
	JwtConfig JWT
	BlackCache *local_cache.Cache
)


type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                  // 签发者
}

//
//func JWTGetConfig(Config viper.Viper)(errorcode int){
//	JwtConfig.ExpiresTime = Config.GetInt64("auth.jwt.expires-time")
//	JwtConfig.BufferTime = Config.GetInt64("auth.jwt.buffer-time")
//	JwtConfig.Issuer = Config.GetString("auth.jwt.signing-key")
//	return errmsg.GLOBAL_STATS_OK
//}



//中间件jwt认证方式
/*待做事项
1.获取jwt的token；
2.校验有效性;
*/
func JWTAuth(redis *redis.Client,db *gorm.DB,cache *local_cache.Cache) gin.HandlerFunc {
	JwtRedis =  redis
	JwtDB = db
	BlackCache = cache
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			LogGlobal.Error("token is not exists")
			dpresponse.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}else{
			if token != "myzone"{
				LogGlobal.Error("token is error")
				dpresponse.FailWithDetailed(gin.H{"reload": true}, "token is error", c)
				c.Abort()
				return
			}else{
				c.Next()
			}
		}
		//if jwtService.IsBlacklist(token) {
		//	dpresponse.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
		//	c.Abort()
		//	return
		//}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				dpresponse.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			dpresponse.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开
		//if err, _ = userService.FindUserByUuid(claims.UUID.String()); err != nil {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + JwtConfig.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.GVA_CONFIG.System.UseMultipoint {
				err, RedisJwtToken := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					LogGlobal.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
