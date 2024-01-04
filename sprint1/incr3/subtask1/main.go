package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.URL)
			return nil
		},
	}
	response, err := client.Get("http://ya.ru")

	_, err = io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}

	// if len(body) > 512 {
	// 	body = body[:512]
	// }
	// fmt.Print(string(body))
}
