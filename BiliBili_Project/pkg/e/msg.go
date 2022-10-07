package e

var MsgFlags = map[int]string{
	SUCCESS:      "OK",
	ERROR:        "fail",
	InvaliParams: "请求参数错误",

	ErrorAuthCheckTokenFail:        "Token验证失败",
	ErrorAuthCheckTokenTimeout:     "Token验证超时",
	ErrorAuthInsufficientAuthority: "无权限",

	ErrorUploadFile: "上传文件失败",
	ErrorLikeExist:  "已收藏",
	ErrorFavorExist: "以点赞",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[500]
}
