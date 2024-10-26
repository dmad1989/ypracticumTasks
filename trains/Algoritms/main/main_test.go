package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type inParams struct {
	arr []int
	num int
}

var tests = []struct {
	name     string
	inParams inParams
	expected int
}{{
	name:     "10 el's",
	inParams: inParams{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, num: 10},
	expected: 9,
},
	{
		name: "100 el's",
		inParams: inParams{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
			41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
			51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
			61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
			71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
			81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
			91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
		}, num: 30},
		expected: 29,
	},
}

func TestFindElementB(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findElementB(tt.inParams.arr, tt.inParams.num)
			assert.Equal(t, res, tt.expected)
		})
	}
}

func TestFindElementL(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findElementL(tt.inParams.arr, tt.inParams.num)
			assert.Equal(t, res, tt.expected)
		})
	}
}

func BenchmarkFindElementB(b *testing.B) {
	b.StopTimer()
	arr := makeArray(0, 1000)
	for i := 0; i < b.N; i++ {
		findElementB(arr, rand.Intn(1000))
	}
}

func BenchmarkFindElementL(b *testing.B) {
	b.StopTimer()
	arr := makeArray(0, 1000)
	for i := 0; i < b.N; i++ {
		findElementL(arr, rand.Intn(1000))
	}
}

func makeArray(start int, end int) (result []int) {
	result = make([]int, 0, end-start)
	for value := start; value <= end; value++ {
		result = append(result, value)
	}

	return result
}
