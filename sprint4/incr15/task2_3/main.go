package main

import "fmt"

func main() {
	chIn := make(chan int)
	chOut := make(chan int)
	quit := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			chIn <- i
		}
		close(chIn)
	}()
	go func() {
		for x := range chIn {
			chOut <- x * 2
		}
		close(chOut)
	}()
	go func() {
		for i := range chOut {
			fmt.Printf("%d ", i)
		}
		quit <- struct{}{}
	}()
	<-quit
}
