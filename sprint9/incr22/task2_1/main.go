package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

// PrimeToFile записывает в файл fname простые числа,
// которые меньше или равны n.
func PrimeToFile(n int, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	k := 0
	for i := 1; i <= n; i += 2 {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			fmt.Fprintf(w, "%d ", i)
			k++
			if k == 10 {
				fmt.Fprintf(w, "\r\n")
				k = 0
			}

		}
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := PrimeToFile(10000, "prime.txt"); err != nil {
		panic(err)
	}
}
