package jwt

import (
	"time"

	"github.com/chccc1994/bilibili/models"
	"github.com/chccc1994/bilibili/utils"
	"github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte(utils.JwtSecret)

type MyClaims struct {
	UserId    uint
	Authority int
	jwt.StandardClaims
}

func GenToken(user models.User) (string, error) {

	c := MyClaims{
		UserId:    user.ID,
		Authority: user.Authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "FanOne",
			Subject:   "TOKEN",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
	// return token.SignedString(MySecret)
}

// 解析Token
func ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	// 解析token

	// token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
	// 	func(token *jwt.Token) (i interface{}, err error) {
	// 		return jwtkey, nil
	// 	})
	// if err != nil {
	// 	return nil, nil, err
	// }
	// if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
	// 	return token, claims, nil
	// }
	// return nil, nil, errors.New("invalid token")
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtkey, nil
	})
	return token, claims, err
}
