package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 1)
	go func() {
		ch <- 7
	}()
	v := <-ch
	fmt.Println(v)
}
