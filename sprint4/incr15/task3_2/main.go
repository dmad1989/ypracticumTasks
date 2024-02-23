package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := generator("Hello")
	for msg := range ch {
		fmt.Println(msg)
	}
}

// Тут ваш генератор
func generator(msg string) chan string {
	in := make(chan string)

	go func() {
		defer close(in)
		for i := 0; i < 5; i++ {
			in <- msg + " " + strconv.Itoa(i)
		}
	}()
	return in
}
