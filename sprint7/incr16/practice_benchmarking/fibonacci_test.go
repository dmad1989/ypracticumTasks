package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func BenchmarkFibo(b *testing.B) {
	count := 20
	b.Run("recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FiboRecursive(count)
		}
	})

	b.Run("optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FiboOptimized(count)
		}
	})
}

func BenchmarkSortSlice(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, 10000)
	b.ResetTimer() // сбрасываем счётчик, чтобы инициализация слайса не посчиталась

	var cmps int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := 0; i < len(slice); i++ {
			slice[i] = rand.Intn(1000)
		}
		b.StartTimer()
		// сортируем
		sort.Slice(slice, func(i, j int) bool {
			cmps++ // увеличиваем счётчик
			return slice[i] < slice[j]
		})
	}
	b.ReportMetric(float64(cmps)/float64(b.N), "compares/op")
}
