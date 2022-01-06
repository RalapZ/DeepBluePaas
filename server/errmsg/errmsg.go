package errmsg

var (
	HTTPSUCCESS = 200
	HTTPERROR = 500
	HTTP_REQUEST_FORBIDDEN = 403

	LOG_PATH_NOT_IS_EXIST= 800   //1800-1850   日志异常码
	LOG_FILENAME_NOT_IS_EXIST=801
	LOG_LOGTYPE_NOT_IS_EXIST=802
	LOG_LEVEL_NOT_IS_EXIST=803
	LOG_NORMAL= 888

)

var ErrMsg = map[int]string {
	HTTPERROR: "访问异常",
	HTTPSUCCESS: "访问成功",
	HTTP_REQUEST_FORBIDDEN: "禁止访问",
	LOG_PATH_NOT_IS_EXIST: "日志路径配置不存在",

}

func ErrMsgGet(code int) string{
	if _,ok:=ErrMsg[code];!ok {
		return "未知异常"
	}
	return ErrMsg[code]
}