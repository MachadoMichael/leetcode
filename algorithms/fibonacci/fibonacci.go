package main

import (
	"fmt"
)

func main() {

	fmt.Println(fibonacciIterative(8))
	fmt.Println(fibonacciRecursive(8))
}

func fibonacciIterative(n int) int {
	var a, b int = 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}
