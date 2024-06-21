// Implement a fibonacci function that returns a function (a closure) that returns succesive
// fibonacci numbers (0, 1, 1, 2, 3, 5, .....).

package main

import "fmt"

func fibonacci() func() int {
	a := 0
	b := 1
	i := 0
	return func() int {
		if i == 0 {
			i++
			return a
		} else if i == 1 {
			i++
			return b
		} else {
			i++
			a, b = b, a+b
			return b
		}
	}
}

func main() {
	fib_num := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib_num())
	}
}
