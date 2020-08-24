package utils

import (
	"fmt"
	"time"
	
	"github.com/dgrijalva/jwt-go"
)

const (
	KEY                    string = "JWT-ARY-STARK"
	DEFAULT_EXPIRE_SECONDS int    = 86400
)

type User struct {
	UserId   uint64 `json:"user_id"`
	Username string `json:"username"`
}

type MyCustomClaims struct {
	User
	jwt.StandardClaims
}

func GenerateToken(userId uint64, username string) (tokenString string, err error) {
	mySigningKey := []byte(KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(DEFAULT_EXPIRE_SECONDS)).Unix()
	user := User{userId, username}
	claims := MyCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    user.Username,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("生成令牌失败 !! error :%v", err)
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (user *User, err error) {
	var token *jwt.Token
	token, err = jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	if err != nil {
		return
	}
	
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		user = &User{
			UserId:   claims.User.UserId,
			Username: claims.User.Username,
		}
		return
	}
	return
}
