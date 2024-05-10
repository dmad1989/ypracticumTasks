package main

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func main() {
	// соль лучше генерировать хорошим ГПСЧ
	// здесь для наглядности просто строка 8 байт
	salt := []byte("someSalt")
	// пароль и соль передаются функции слайсом байт
	// 1<<15, 8, 1 — факторы требуемой вычислительной сложности
	// по последним данным, за 2017 год, эксперты криптозащиты рекомендуют
	// для интерактивной аутентификации параметры сложности 32768, 8, 1
	// 32 — требуемая длина полученного ключа
	key, err := scrypt.Key([]byte("very secret password"), salt, 1<<15, 8, 1, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	// для читаемости перекодируем ключ в Base64 символами ASCII
	fmt.Println(base64.StdEncoding.EncodeToString(key))
}
