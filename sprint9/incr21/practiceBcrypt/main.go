package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

const (
	Salt = `bcryptsalt`
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Salt+password), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(Salt+password)) == nil
}

func main() {
	password := "mypass"
	hash, err := HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hash, CheckPasswordHash(password, hash))
}
