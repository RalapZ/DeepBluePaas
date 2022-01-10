package config

import "github.com/RalapZ/DeepBluePaas/server/errmsg"

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
