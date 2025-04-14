package main

import "fmt"

func visit(numbers []int, callback func(int2 int)) {
	for _, n := range numbers {
		callback(n * 3)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n, "- Printed withing the callback function.")
	})

	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n, "- Printedddd withing the callback function.")
	})
}
