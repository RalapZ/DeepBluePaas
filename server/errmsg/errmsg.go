package errmsg

var (
	HTTPSUCCESS = 200
	HTTPERROR = 500
	HTTP_REQUEST_FORBIDDEN = 403

	LOG_PATH_NOT_IS_EXIST= 1800   //1800-1899   日志异常码
	LOG_FILENAME_NOT_IS_EXIST=1801
	LOG_LOGTYPE_NOT_IS_EXIST=1802
	LOG_LEVEL_NOT_IS_EXIST=1803
	LOG_NORMAL= 1899


	GRPC_PORT_NOT_IS_EXIST = 1900        //1900-1999   grpc异常码

	REGISTRY_CONFIG_NOT_ERROR= 2000      //2000-2099 配置中心错误码


	GLOBAL_STATS_OK = 8888   //全局正常

)

var ErrMsg = map[int]string {
	HTTPERROR: "访问异常",
	HTTPSUCCESS: "访问成功",
	HTTP_REQUEST_FORBIDDEN: "禁止访问",
	LOG_PATH_NOT_IS_EXIST: "日志路径配置不存在",


	GLOBAL_STATS_OK: "状态正常",

}

func ErrMsgGet(code int) string{
	if _,ok:=ErrMsg[code];!ok {
		return "未知异常"
	}
	return ErrMsg[code]
}