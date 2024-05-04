package main

// реализуйте функцию Reverse
// ...

func Reverse[T any](slice []T) []T {
	i := 0
	j := len(slice) - 1
	for i != j && j > i {
		b := slice[i]
		e := slice[j]
		slice[i] = e
		slice[j] = b
		i++
		j--
	}
	return slice
}

// func Reverse[T any](s []T) []T {
//     count := len(s)
//     for i := 0; i < count/2; i++ {
//         s[i], s[count-i-1] = s[count-i-1], s[i]
//     }
//     return s
// }
