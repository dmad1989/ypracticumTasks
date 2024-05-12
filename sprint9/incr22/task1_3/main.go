package main

import (
	"bytes"
	"fmt"
)

var dict = map[string]string{
	"😅": "a",
	"😎": "e",
	"😂": "g",
	"🥶": "h",
	"🤓": "j",
	"🥰": "o",
	"😶": "p",
	"🐱": "r"}

func main() {
	b := []byte("😂🥰😶🥶😎🐱🌚")
	for k, v := range dict {
		b = bytes.ReplaceAll(b, []byte(k), []byte(v))
	}
	fmt.Printf("%q", b)
	// fmt.Println(b)
}
