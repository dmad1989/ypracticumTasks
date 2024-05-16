package main

import (
	"fmt"
	"testing"
)

func TestPrime(t *testing.T) {
	if err := PrimeToFile(50000, "prime.txt"); err != nil {
		t.Error(err)
	}
	lines, err := SearchPrime("prime.txt", []int{2, 137, 977, 2239, 2293, 9941,
		16693, 16699, 26647, 37579, 48337})
	if err != nil {
		t.Error(err)
	}
	if fmt.Sprint(lines) != "[1 4 17 34 34 123 193 194 293 398 498]" {
		t.Errorf("unexpected lines %v", lines)
	}
}
