package main

import (
	"fmt"
	"sync"
)

func main() {
	inCh := gen(2, 3)
	ch1 := square(inCh)
	ch2 := square(inCh)
	for n := range fanIn(ch1, ch2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for _, n := range nums {
			outCh <- n
		}
	}()

	return outCh
}

func square(inCh chan int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for n := range inCh {
			outCh <- n * n
		}
	}()

	return outCh
}

// эту функцию нужно реализовать
func fanIn(chs ...chan int) chan int {
	finalChan := make(chan int)
	var wg sync.WaitGroup

	out := func(ch chan int) {
		for data := range ch {
			finalChan <- data
		}
		wg.Done()
	}
	wg.Add(len(chs))

	for _, ch := range chs {
		go out(ch)
	}

	go func() {
		wg.Wait()
		close(finalChan)
	}()
	return finalChan
}
