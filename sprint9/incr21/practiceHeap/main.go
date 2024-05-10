package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Name     string
	Priority int // приоритет элемента в очереди
}

type Queue []*Item

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	// сортируем очередь в порядке убывания
	return q[i].Priority > q[j].Priority
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x.(*Item))
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // присваиваем последнему элементу nil, чтобы избежать утечек памяти
	*q = old[0 : n-1]
	return item
}

func main() {
	q := Queue{
		{"John", 1},
		{"Caroline", 5},
		{"Alex", 4},
	}
	heap.Init(&q)

	// добавляем в коллекцию ещё один элемент
	heap.Push(&q, &Item{"Marta", 3})

	for q.Len() > 0 {
		// извлекаем из очереди все элементы
		item := heap.Pop(&q).(*Item)
		fmt.Printf("%.2d:%s\n", item.Priority, item.Name)
	}
}
