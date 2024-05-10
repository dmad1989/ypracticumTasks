package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("🌝🌖🌗🌘🌚🌒🌓🌔🌝")
	// допишите код
	fmt.Println(bytes.IndexRune(b, '🌚'))
}
