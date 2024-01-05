package main

import (
	"fmt"
	"net/http"

	resty "github.com/go-resty/resty/v2"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	var users []User
	url := "https://jsonplaceholder.typicode.com/users"
	client := resty.New()

	resp, err := client.R().SetResult(&users).EnableTrace().Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode() != http.StatusOK {
		fmt.Printf("status %s\n", resp.Status())
		fmt.Println()
	}

	for _, user := range users {
		fmt.Printf("User: %s (%d) %s\n", user.Username, user.ID, user.Email)
	}
}
