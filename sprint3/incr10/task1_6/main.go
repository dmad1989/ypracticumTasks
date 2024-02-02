package main

import (
	"fmt"
	"time"
)

func main() {
	// допишите код здесь
	// ...
	start := time.Now()
	t := time.NewTicker(2 * time.Second)
	for i := 0; i < 10; i++ {
		z := <-t.C
		fmt.Println(int(z.Sub(start).Seconds()))
	}
}
