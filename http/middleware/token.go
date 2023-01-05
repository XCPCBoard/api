package middleware

import (
	"github.com/XCPCBoard/api/errors"
	"github.com/XCPCBoard/common/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		//TokenSecret在config.yml中定义
		return []byte(config.Conf.TokenSecret), nil
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
		}
		//将id和Name写入ctx
		ctx.Set("xcpc_user_id", token.Id)
		ctx.Set("xcpc_user_name", token.Audience)

		ctx.Next()
	}
}
