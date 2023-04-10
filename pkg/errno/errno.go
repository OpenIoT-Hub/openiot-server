package errno

var MsgFlags = map[int]string{
	Success: "ok",

	InvalidParams: "请求参数错误",

	Error:                       "fail",
	ErrorDatabaseQuery:          "数据库查询错误",
	ErrorDatabaseRecordNotFound: "数据库返回空值",
	ErrorUserNotExist:           "用户不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
