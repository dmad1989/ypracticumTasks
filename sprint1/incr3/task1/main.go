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

	if len(body) > 512 {
		body = body[:512]
	}
	fmt.Print(string(body))
}
