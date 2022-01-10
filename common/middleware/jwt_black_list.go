package middleware

import (
	"context"
	//"github.com/RalapZ/DeepBluePaas/server/config"

	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type JwtService struct {
}

//@author:RalapZ
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = JwtDB.Create(&jwtList).Error
	if err != nil {
		return
	}
	BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@author: RalapZ
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

//func (jwtService *JwtService) IsBlacklist(jwt string) bool {
//	_, ok := config.BlackCache.Get(jwt)
//	return ok
//	//err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
//	//isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
//	//return !isNotFound
//}

//@author: Ralap
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = JwtRedis.Get(context.Background(), userName).Result()
	return err, redisJWT
}

//@author: RalapZ
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(JwtConfig.ExpiresTime) * time.Second
	err = JwtRedis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
//
//func LoadAll() {
//	var data []string
//	err := config.GVA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
//	if err != nil {
//		middleware.LogGlobal.Error("加载数据库jwt黑名单失败!", zap.Error(err))
//		return
//	}
//	for i := 0; i < len(data); i++ {
//		config.BlackCache.SetDefault(data[i], struct{}{})
//	} // jwt黑名单 加入 BlackCache 中
//}
