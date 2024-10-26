package main

import "fmt"

func main() {
	res := findElementB([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 10)
	fmt.Println(res)
}

func findElementL(arr []int, num int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}

func findElementB(arr []int, num int) int {
	maxi := len(arr) - 1
	mi := maxi / 2
	mini := 0
	for mi != 0 {
		if arr[mi] == num {
			return mi
		}

		if arr[mi] < num {
			mini = mi
		}

		if arr[mi] > num {
			maxi = mi
		}

		mi = ((maxi - mini) / 2) + mini
	}
	return -1
}
