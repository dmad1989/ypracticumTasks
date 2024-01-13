package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("TASK_DURATION"))

	envList := os.Environ()
	// выводим первые пять элементов
	for i := 0; i < 5 && i < len(envList); i++ {
		fmt.Println(envList[i])
	}
}
