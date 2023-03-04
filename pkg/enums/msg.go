package enums

var MsgFlags = map[int]string{
	SUCCESS:           "ok",
	ERROR:             "fail",
	INVALID_PARAMS:    "请求参数错误",
	ERROR_AUTH_FAILED: "认证错误",
	SERVER_ERROR:      "服务器错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
