package config

import (
	"context"
	"fmt"
	"github.com/RalapZ/DeepBluePaas/common/middleware"
	"github.com/RalapZ/DeepBluePaas/server/errmsg"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

//
//cache:
//datasource: redis
//redis:
//addr: "127.0.0.1"
//port: 6379
//#    passwd: ""

var redisCfg Redis

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Port 	 string
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}


func RedisGetClientConn() {
	//redisCfg := GVA_REDIS.Redis
	fmt.Println("redisCfg:",redisCfg)
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr+":"+redisCfg.Port,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		middleware.LogGlobal.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		middleware.LogGlobal.Info("redis connect ping response:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}
}

func RedisGetConfig()(errorcode int){
	//var redisCfg Redis
	if ok:= Config.IsSet("registry.nacos.client.namespaceid");!ok{
		return errmsg.REGISTRY_CONFIG_NOT_ERROR
	}
	redisCfg.Port = Config.GetString("cache.redis.port")
	redisCfg.Addr = Config.GetString("cache.redis.addr")
	redisCfg.Password = Config.GetString("cache.redis.passwd")
	redisCfg.DB = Config.GetInt("cache.redis.db")

	//RedisGetClientConn()
	return errmsg.GLOBAL_STATS_OK
}
