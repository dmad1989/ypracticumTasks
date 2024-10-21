package main

import "fmt"

func main() {
	res := findElement([]int{1, 2, 3, 4, 5}, 5)
	fmt.Println(res)
}

func findElement(arr []int, num int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}
