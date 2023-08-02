package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/qppHUST/prepareForOffer/web/model"
)

var secret = []byte("qpphust")

func GetJWT(authRequest model.TokenRequest) (string, error) {
	claim := jwt.MapClaims{
		"username": authRequest.Username,
		"passqord": authRequest.Password,
		"Exp":      jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		"Issuer":   "qpphust",
	}
	//生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//生成签名字符串
	return token.SignedString(secret)
}

// 验证token是否合法
func ValidateJWT(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

func ParseJWT(tokenString string) (*jwt.Claims, error) {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	return &token.Claims, nil
}
