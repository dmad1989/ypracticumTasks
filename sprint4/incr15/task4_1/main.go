package main

import (
	"fmt"
	"math"
)

func main() {
	c := gen(2, 3)
	out := square(c)

	for res := range out {
		fmt.Println(res)
	}
}

// реализация генератора gen здесь
func gen(nums ...int) chan int {
	input := make(chan int)
	go func() {
		defer close(input)
		for _, i := range nums {
			input <- i
		}
	}()
	return input
}

// реализация square здесь
func square(in chan int) chan int {
	squareRes := make(chan int)

	go func() {
		defer close(squareRes)
		for data := range in {
			squareRes <- int(math.Pow(float64(data), 2))
		}
	}()
	return squareRes
}
