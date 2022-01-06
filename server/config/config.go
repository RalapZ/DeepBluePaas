package config

import (
	"fmt"
	"os"

	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/spf13/viper"

	"github.com/RalapZ/DeepBluePaas/common/middleware"
	"github.com/RalapZ/DeepBluePaas/server/errmsg"
)


var (
	IPAddr  = ":18081"
	Config *viper.Viper
	LogConfig middleware.LogConfig
	ServerInfo Server
	NacosCliConfig constant.ClientConfig
	NacosSerConfig constant.ServerConfig

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
//addr: 127.0.0.1
//port: 8848
//namespaceid: deeblue
//timeoutms: 5000
//notLoadcacheatstart: true
//logdir: "/tmp/nacos/log"
//cachedir: "/tmp/nacos/cache"
//rotatetime: "1h"
//maxage: 3
//loglevel: "debug"
func NacosGetConfig()(errorcode int){
	if ok:= Config.IsSet("registry.nacos.client.namespaceid");!ok{
		return errmsg.REGISTRY_CONFIG_NOT_ERROR
	}

	NacosSerConfig.IpAddr = Config.GetString("registry.nacos.server.addr")
	NacosSerConfig.Port = Config.GetUint64("registry.nacos.server.port")

	NacosCliConfig.NamespaceId = Config.GetString("registry.nacos.client.namespaceid")
	NacosCliConfig.TimeoutMs = Config.GetUint64("registry.nacos.client.timeoutms")
	NacosCliConfig.NotLoadCacheAtStart = Config.GetBool("registry.nacos.client.notLoadcacheatstart")
	NacosCliConfig.LogDir = Config.GetString("registry.nacos.client.logdir")
	NacosCliConfig.CacheDir = Config.GetString("registry.nacos.client.cachedir")
	NacosCliConfig.RotateTime = Config.GetString("registry.nacos.client.rotatetime")
	NacosCliConfig.MaxAge = Config.GetInt64("registry.nacos.client.maxage")
	NacosCliConfig.LogLevel =  Config.GetString("registry.nacos.client.loglevel")
	return errmsg.GLOBAL_STATS_OK
}



func LogGetConfig() (errorcode int){
	if ok:= Config.IsSet("logger.path");!ok{
		return errmsg.LOG_PATH_NOT_IS_EXIST
	}
	if ok:= Config.IsSet("logger.filename");!ok{
		return errmsg.LOG_FILENAME_NOT_IS_EXIST
	}
	if ok:= Config.IsSet("logger.logtype");!ok{
		return errmsg.LOG_LOGTYPE_NOT_IS_EXIST
	}
	if ok:= Config.IsSet("logger.level");!ok{
		return errmsg.LOG_LEVEL_NOT_IS_EXIST
	}
	LogConfig.LogLevel=Config.GetString("logger.level")
	LogConfig.LogName=Config.GetString("logger.filename")
	LogConfig.LogPath=Config.GetString("logger.path")
	LogConfig.LogType=Config.GetString("logger.logtype")
	//Config.IsSet()

	return errmsg.GLOBAL_STATS_OK
}


//解析配置文件  可优化，变更为map循环;还有各个业务正常码
func ParserConfig() {
	fmt.Println(os.Getwd())
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
