package main

import "fmt"

func filter(slice []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	even := filter(slice, func(n int) bool {
		return n%2 == 0
	})

	odd := filter(slice, func(n int) bool {
		return n%2 != 0
	})

	multiple3 := filter(slice, func(n int) bool {
		return n%3 == 0
	})

	fmt.Println(even)
	fmt.Println(odd)
	fmt.Println(multiple3)
}
