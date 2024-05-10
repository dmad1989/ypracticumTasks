package main

import (
	"fmt"
	"math/rand"
	"time"
)

const Range = 6

var rnd *rand.Rand

func Scores(n int) []byte {
	scores := make([]byte, n)
	for i := 0; i < n; i++ {
		scores[i] = byte(rnd.Int31n(Range) + 1)
	}
	return scores
}

func Prob(scores []byte) [Range]float64 {
	var counts [Range]int
	n := len(scores)
	for i := 0; i < n; i++ {
		counts[scores[i]-1]++
	}
	var av [Range]float64
	for i := 0; i < 6; i++ {
		av[i] = float64(counts[i]) / float64(n)
	}
	return av
}

func main() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 100; i < 10001; i *= 10 {
		fmt.Println("n=", i, Prob(Scores(i)))
	}
}
