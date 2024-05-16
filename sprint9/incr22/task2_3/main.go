package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrimeToFile(n int, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	s := make([]byte, n+1)

	fmt.Fprintf(w, "%d ", 2)
	count := 1 // уже записали число 2
	// будем перебирать только нечётные числа
	for i := 3; i <= n; i += 2 {
		if s[i] == 0 {
			fmt.Fprintf(w, "%d ", i)
			for k := 2 * i; k <= n; k += i {
				s[k] = 1
			}
			count++
			if count == 10 {
				fmt.Fprintf(w, "\r\n")
				count = 0
			}
		}
	}
	w.Flush()

	return err
}

// SearchPrime находит строки в файле fname с указанными в primes простыми числами.
func SearchPrime(fname string, primes []int) (lines []int, err error) {
	f, err := os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	var lineNum, pCount int
	var primeStr string
	for s.Scan() {
		lineNum++
		line := s.Text()

		for _, v := range strings.Fields(line) {
			if primeStr == "" {
				primeStr = strconv.FormatInt(int64(primes[pCount]), 10)
			}

			fmt.Printf("prime: %s, v: %s", primeStr, v)
			fmt.Println()
			if v == primeStr {
				lines = append(lines, lineNum)
				pCount++
				if pCount == len(primes) {
					return
				}
				primeStr = ""
			}
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return
}
