package main

import (
	"fmt"
	"time"
)

func main() {
	// допишите код здесь
	birthday := time.Date(1993, time.November, 26, 0, 0, 0, 0, time.Local)

	days := time.Until(birthday.AddDate(100, 0, 0)).Hours() / 24
	fmt.Println(days)
}
