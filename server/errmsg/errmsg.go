package errmsg

var (
	HTTPSUCCESS = 200
	HTTPERROR = 500
	HTTP_REQUEST_FORBIDDEN = 403
)

var ErrMsg = map[int]string {
	HTTPERROR: "访问异常",
	HTTPSUCCESS: "访问成功",
	HTTP_REQUEST_FORBIDDEN: "禁止访问",
}

func ErrMsgGet(code int) string{
	if _,ok:=ErrMsg[code];!ok {
		return "未知异常"
	}
	return ErrMsg[code]
}