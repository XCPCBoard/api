package http

import (
	"fmt"
	"github.com/XCPCBoard/api/http/middleware"
	"github.com/XCPCBoard/api/http/token"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

var MsgCode = map[string]int{
	"success": 20000,
	"fail":    errors.ERROR.Code,
	"unKnow":  errors.INNER_ERROR.Code,
}

// SuccessResponse 响应成功
func SuccessResponse(ctx *gin.Context, data map[string]interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": MsgCode["success"],
		"msg":  "ok",
		"data": data,
	})
}

// SuccessResponseAddToken 响应成功，并添加token
func SuccessResponseAddToken(ctx *gin.Context, data map[string]interface{}) {
	id, ok := ctx.Get(middleware.TokenIDStr)
	name, ok2 := ctx.Get(middleware.TokenAccountStr)
	if !ok || !ok2 {
		e := errors.CreateError(MsgCode["unKnow"], "获取用户token中的id和name失败", ctx.Keys)
		logger.L.Err(e, 0)
		ctx.Error(e)
	}
	token, err := token.GenerateToken(fmt.Sprintf("%v", id),
		fmt.Sprintf("%v", name))
	if err != nil {
		e := errors.CreateError(MsgCode["unKnow"], "生产token失败", err)
		logger.L.Err(e, 0)
		ctx.Error(e)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  MsgCode["success"],
		"msg":   "ok",
		"data":  data,
		"token": token,
	})
}

// FailResponse 响应失败
func FailResponse(msg string, data map[string]interface{}) gin.H {
	res := gin.H{
		"code": MsgCode["fail"],
		"msg":  msg,
		"data": data,
	}
	return res
}
