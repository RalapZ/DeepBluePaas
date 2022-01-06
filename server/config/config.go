package config

import (
	"fmt"
	"github.com/RalapZ/DeepBluePaas/common/middleware"
	"github.com/RalapZ/DeepBluePaas/server/errmsg"
	"os"

	"github.com/spf13/viper"
)


var Config *viper.Viper

var LogConfig middleware.LogConfig


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

	return errmsg.LOG_NORMAL
}

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
	errorcode := LogGetConfig()
	if errorcode != errmsg.LOG_NORMAL {
		panic(errorcode)
	}
	fmt.Println("logconfig",LogConfig.LogPath)

}
