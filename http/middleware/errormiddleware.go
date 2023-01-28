package middleware

import (
	"github.com/XCPCBoard/api/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 先调用c.Next()执行后面的中间件
		// 所有中间件及router处理完毕后从这里开始执行
		// 检查c.Errors中是否有错误

		if err := c.Errors.Last(); err != nil {
			// 若是自定义的错误则将code、msg返回
			if myErr, ok := err.Err.(*errors.MyError); ok {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, myErr)
			} else {
				// 若非自定义错误则返回详细错误信息err.Error()
				c.AbortWithStatusJSON(http.StatusBadRequest, errors.GetError(errors.INNER_ERROR, err.Error()))
			}
			logger.Logger.Warn(err.Error(), 0, "")
		}

		return // 检查最后一个错误就行

	}
}
