package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	res, err := encode(16)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func encode(size int) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b) // записываем байты в слайс b
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
