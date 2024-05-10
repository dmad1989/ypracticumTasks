package main

import (
	"crypto/rand"
	"fmt"
)

const Alphabet = 26

func GenCryptoKey(n int) (string, error) {
	data := make([]byte, n)
	nrnd, err := rand.Read(data)
	if err != nil {
		return ``, err
	} else if nrnd != n {
		return ``, fmt.Errorf(`nrnd %d != n %d`, nrnd, n)
	}
	for i := range data {
		data[i] = 'A' + data[i]%Alphabet
	}
	return string(data), nil
}

func main() {
	for i := 16; i <= 64; i += 16 {
		key, err := GenCryptoKey(i)
		if err != nil {
			panic(err)
		}
		fmt.Println(key)
	}
}
