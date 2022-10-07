package main

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var MySecret = []byte("夏天夏天悄悄过去")

type MyClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenToken(username, password string) (string, error) {
	c := MyClaim{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "jwt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// return token.SignedString([]byte("secret"))
	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*MyClaim, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaim{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaim); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
func main() {
	fmt.Println("hello jwt")
	tokenKey, err := GenToken("lii", "123456")
	if err != nil {
		panic(err)
	}
	fmt.Println("jet token key:", tokenKey)

	mc, err := ParseToken(tokenKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("UserName:", mc.Username)
	fmt.Println("Password:", mc.Password)
}
