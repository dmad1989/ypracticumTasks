package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("😂😂😂")
	b := []byte("😎😎😎")
	// допишите код
	fmt.Println(bytes.Compare(a, b))
}
