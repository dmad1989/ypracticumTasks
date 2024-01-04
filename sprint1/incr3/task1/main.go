package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://practicum.yandex.ru")
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}

	for k, v := range body {
		if k == 513 {
			break
		}
		fmt.Printf("byte %d, value %d", k, v)
		fmt.Println()
	}
}
