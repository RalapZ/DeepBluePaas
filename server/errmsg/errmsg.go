package errmsg

var (
	/*
	错误码规划：
	1.550-800    自定义http异常码
	1.1800-1899  日志异常码
	2.1900-1999  grpc异常码
	3.2000-2099  注册中心/配置中心错误码
	*/
	HTTP_RESPONSE_SUCCESS = 200
	HTTP_RESPONSE_ERROR = 500
	HTTP_RESPONSE_FORBIDDEN = 403

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
	HTTP_RESPONSE_ERROR: "访问异常",
	HTTP_RESPONSE_SUCCESS: "访问成功",
	HTTP_RESPONSE_FORBIDDEN: "禁止访问",
	LOG_PATH_NOT_IS_EXIST: "日志路径配置不存在",


	GLOBAL_STATS_OK: "状态正常",

}

func ErrMsgGet(code int) string{
	if _,ok:=ErrMsg[code];!ok {
		return "未知异常"
	}
	return ErrMsg[code]
}