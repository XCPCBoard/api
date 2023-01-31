package token

import (
	"context"
	"errors"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/common/dao"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const tokenExpiresTime = 60 * 15

//GenerateToken 生成token
func GenerateToken(userName string, userId string) (string, error) {
	expiresTime := time.Now().Unix() + tokenExpiresTime
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
	//if err != nil {
	//	return "", err
	//}
	//token = "Bearer " + token//不需要
	return token, err
}

//GenerateTokenAndSetRedis 生成token并存入redis
func GenerateTokenAndSetRedis(userName string, userId string) (string, error) {
	token, err := GenerateToken(userName, userId)

	if e := dao.RedisClient.Set(context.Background(), userId, token, tokenExpiresTime).Err(); e != nil {
		return "", errors.New("token 写入redis时出错：\n" + e.Error())
	}

	return token, err
}
