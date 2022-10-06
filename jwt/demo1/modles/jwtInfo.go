package modles

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 2 //定义JWT过期时间为2小时

var MySecret = []byte("夏天悄悄过去")

type MyClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

// 生成JWT
func GenToken(username string) (string, error) {
	C := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //设置超时时间,输出过期的时间，按照格式
			Issuer:    "my-project",                               //签发人
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, C)
	//使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*MyClaims); ok && token.Valid { //valid:正确，有效  校验token
		return claim, err //返回JWT中的字段
	}
	return nil, errors.New("invalid token")
}
