package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"
)

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

// MaxPrimeInterval читает файл fname с простыми числами и возвращает
// максимальный интервал dif между числами.
// prime — первое число с найденным интервалом до следующего числа.
func MaxPrimeInterval(fname string) (prime int, dif int, err error) {
	f, err := os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var prev, num int
	var s, wnum string
	for err != io.EOF {
		s, err = r.ReadString(' ')
		if err != nil && !errors.Is(err, io.EOF) {
			return
		}
		wnum = strings.TrimSpace(s)
		if len(wnum) == 0 {
			continue
		}
		if num, err = strconv.Atoi(wnum); err != nil {
			return
		}
		if num-prev > dif {
			dif = num - prev
			prime = prev
		}
		prev = num
	}
	return prime, dif, nil
}

func TestPrime(t *testing.T) {
	if err := PrimeToFile(50000, "prime.txt"); err != nil {
		t.Error(err)
	}
	if prime, dif, err := MaxPrimeInterval("prime.txt"); err != nil {
		t.Error(err)
	} else if prime != 31397 || dif != 72 {
		t.Errorf("unexpected prime=%d / interval=%d", prime, dif)
	}
}
