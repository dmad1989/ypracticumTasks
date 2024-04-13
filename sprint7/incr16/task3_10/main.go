package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

// реализуем интерфейс sort.Interface для сортировки по возрасту

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

// Sort сортирует слайс ByAge, так он реализует интерфейс sort.Interface.
func (a ByAge) Sort() {
	sort.Sort(a)
}

func main() {
	people := ByAge{
		{"Bob", 31},
		{"John", 48}, // John старший
		{"Michael", 17},
		{"John", 26}, // John младший
	}

	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		if people[i].Name == people[j].Name {
			return people[i].Age < people[j].Age
		}
		return people[i].Name < people[j].Name
	})
	fmt.Println(people)
}
