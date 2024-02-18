package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	//структурa, в которую входит набор утверждений JWT c зарегистрированными именами утверждений из спецификации JWT
	jwt.RegisteredClaims
	UserID int
}

const TOKEN_EXP = time.Hour * 3
const SECRET_KEY = "supersecretkey"

func main() {
	tokenString, err := BuildJWTString()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tokenString)

	fmt.Println(GetUserID(tokenString))
}

func BuildJWTString() (string, error) {
	//создаёт новый токен с указанными методом подписи и утверждениями
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
	},
		UserID: 1})
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	// возвращаем строку токена
	return tokenString, nil
}

func GetUserID(tokenString string) int {
	// создаём экземпляр структуры с утверждениями
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(SECRET_KEY), nil
		})
	if err != nil {
		return -1
	}
	if !token.Valid {
		fmt.Println("Token is not valid")
		return -1
	}

	fmt.Println("Token is valid")
	return claims.UserID
}
