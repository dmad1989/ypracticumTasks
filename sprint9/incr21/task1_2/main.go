package main

import (
	"fmt"
	"regexp"
)

const emailPattern = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

func main() {
	m := regexp.MustCompile(emailPattern)

	emails := []string{"$€§@yandex.com", "ivan@mail.ru", "john@gmailyahoo", "fedor@gmail.com", "stepan@yahoo.com", "commanderpike@gmail.com", "greta@abcd@gmail_yahoo.com"}
	for _, e := range emails {
		if b := m.MatchString(e); !b {
			fmt.Printf("Not valid email: %s", e)
			fmt.Println("")
		}

	}

}
