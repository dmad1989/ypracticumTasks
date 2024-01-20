package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	// допишите код
	// 1) создайте переменную типа *log.Logger
	// 2) запишите в неё нужные строки

	// ...

	fmt.Print(&buf)
	// должна вывести
	// mylog: Hello, world!
	// mylog: Goodbye
}
