package http

import "github.com/gin-gonic/gin"

var MsgCode = map[string]int{
	"success": 200,
	"fail":    400,
	"unKnow":  500,
}

//SuccessResponse 响应成功
func SuccessResponse(msg string, data map[string]interface{}) gin.H {
	res := gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	}
	return res
}

//FailResponse 响应失败
func FailResponse(msg string, data map[string]interface{}) gin.H {
	res := gin.H{
		"code": 400,
		"msg":  msg,
		"data": data,
	}
	return res
}
