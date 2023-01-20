package token

import (
	"github.com/XCPCBoard/common/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(userName string, userId string) (string, error) {
	expiresTime := time.Now().Unix() + 1200
	claims := jwt.StandardClaims{
		Audience:  userName,          // 受众(用户名)
		ExpiresAt: expiresTime,       // 失效时间
		Id:        userId,            // 用户id
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    "gin hello",       // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   "login",           // 主题
	}

	//TokenSecret在config.yml中定义
	var jwtSecret = []byte(config.Conf.Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}
