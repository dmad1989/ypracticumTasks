package main

import (
	"fmt"
	"time"
)

func main() {
	currentTimeStr := "2021-09-19T15:59:41+03:00"
	t, err := time.Parse(time.RFC3339, currentTimeStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
