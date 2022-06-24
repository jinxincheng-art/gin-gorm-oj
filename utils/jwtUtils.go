package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var key = []byte("gin-gorm-oj-key")

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(identity, name string) (string, error) {
	userClaims := &UserClaims{
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "",err
	}
	return tokenString,nil
}

func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaims := new(UserClaims)

	claims,err := jwt.ParseWithClaims(tokenString,userClaims, func(token *jwt.Token) (interface{},error){
		return key,nil
	})

	if err != nil {
		return nil,err
	}

	if !claims.Valid {
		return nil,fmt.Errorf("analyse Token Error:%v",err)
	}

	return userClaims,nil
}
