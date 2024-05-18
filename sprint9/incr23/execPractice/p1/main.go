package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("echo", "Hello, world!").Output()
	if err != nil {
		log.Printf("Ошибка: %v", err)
	}
	fmt.Println(string(out))
}
