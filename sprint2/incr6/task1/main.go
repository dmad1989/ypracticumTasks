package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer
	mylog := log.New(&buf, "", 0)

	mylog.Println("Hello world!")
	mylog.Println("Goodbye")

	fmt.Print(&buf)
}
