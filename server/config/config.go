//###############################################################################################//
//                                    ┌─┐       ┌─┐ + +                                          //
//                                 ┌──┘ ┴───────┘ ┴──┐++                                         //
//                                 │       ───       │++ + + +                                   //
//                                 ███████───███████ │+                                          //
//                                 │       ─┴─       │+                                          //
//                                 └───┐         ┌───┘                                           //
//                                     │         │   + +                                         //
//                                     │         └──────────────┐                                //
//                                     │                        ├─┐                              //
//                                     │                        ┌─┘                              //
//                                     └─┐  ┐  ┌───────┬──┐  ┌──┘  + + + +                       //
//                                       │ ─┤ ─┤       │ ─┤ ─┤                                   //
//                                       └──┴──┘       └──┴──┘  + + + +                          //
//                                  神兽出没               永无BUG                                 //
//  Author: ralap                                                                                //
//  Date  : 2022/01/10                                                                           //
//###############################################################################################//
package config

import (
	"fmt"

	"github.com/RalapZ/DeepBluePaas/server/errmsg"
	"github.com/go-redis/redis/v8"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/RalapZ/DeepBluePaas/common/middleware"
)


var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	LocalBlackCache local_cache.Cache

	JwtConfig middleware.JWT

	IPAddr  = ":18081"
	Config *viper.Viper
	LogConfig middleware.LogConfig
	ServerInfo Server
	NacosCliConfig constant.ClientConfig
	NacosSerConfig constant.ServerConfig
	NacosInsConfig vo.RegisterInstanceParam
	LocalAddr string
)


type Server struct{
	Name string
	Grpc uint64
	Http string
}

func ServiceGetConfig()(errorcode int){
	if ok:= Config.IsSet("service.grpc");!ok{
		return errmsg.GRPC_PORT_NOT_IS_EXIST
	}
	ServerInfo.Name=Config.GetString("service.name")
	ServerInfo.Grpc=Config.GetUint64("service.grpc")
	//ServicePort.http=Config.GetString("service.http")
	return errmsg.GLOBAL_STATS_OK
}



//解析配置文件  可优化，变更为map循环;还有各个业务正常码
func ParserConfig() {
	//fmt.Println(os.Getwd())
	Config=viper.New()
	Config.AddConfigPath("/Users/ralap/code/golang/deepblue/DeepBluePaas/etc")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	if err := Config.ReadInConfig(); err!=nil {
		panic(err)
	}
	//fmt.Printf("%##v",config)
	//fmt.Printf("%##v", Config.Get("db.datasource"))
	LogConfig.LogPath=Config.GetString("logger.path1")
	//if ok := Config.IsSet("logger.path");!ok {
	//	fmt.Println(ok)
	//}

	if errcode := JWTGetConfig();errcode != errmsg.GLOBAL_STATS_OK{
		panic(errcode)
	}
	if errcode := RedisGetConfig();errcode != errmsg.GLOBAL_STATS_OK{
		panic(errcode)
	}
	if errcode := getLocalIpaddr();errcode != errmsg.GLOBAL_STATS_OK{
		panic(errcode)
	}


	if errcode :=  NacosGetInstanceConfig();errcode !=errmsg.GLOBAL_STATS_OK{
		panic(errcode)
	}
	fmt.Printf("%##v",NacosInsConfig)

	if errorcode := ServiceGetConfig(); errorcode != errmsg.GLOBAL_STATS_OK {
		panic(errorcode)
	}

	if errorcode := NacosGetConfig();errorcode != errmsg.GLOBAL_STATS_OK {
		panic(errorcode)
	}

	if errorcode := LogGetConfig();errorcode != errmsg.GLOBAL_STATS_OK {
		panic(errorcode)
	}
	fmt.Println("logconfig",LogConfig.LogPath)
}
