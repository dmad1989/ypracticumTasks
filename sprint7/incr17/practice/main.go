package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main() {
	a := 100 // левое слагаемое
	b := -1  // правое слагаемое
	fmt.Printf("Сумма: %d", add(a, b))
}
