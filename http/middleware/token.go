package middleware

import (
	"github.com/XCPCBoard/api/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

const (
	secret = "XCPCBoard2023_C7*1&123" //记得合并后改到config里
)

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}

//AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//错误

		auth := ctx.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			ctx.Error(errors.NewError(http.StatusForbidden, "认证失败"))
			ctx.Abort()
			return
		}
		auth = strings.Fields(auth)[1]
		// 校验token
		token, err := parseToken(auth)
		if err != nil {
			ctx.Error(errors.NewError(http.StatusOK, "token 错误"+err.Error()))
			ctx.Abort()
			return
		}
		//过期
		if !token.VerifyExpiresAt(time.Now().Unix(), false) {
			ctx.Error(errors.NewError(http.StatusOK, "token 超时"+err.Error()))
			ctx.Abort()
			return
		} else {

		}
		ctx.Writer.WriteString(token.Id)
		ctx.Next()
	}
}
