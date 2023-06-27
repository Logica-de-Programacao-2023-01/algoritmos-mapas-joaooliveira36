package main

import (
	"fmt"
)

func fibonacciMap(n int) map[int]int {
	fibMap := make(map[int]int)

	for i := 0; i <= n; i++ {
		fibMap[i] = fibonacci(i)
	}

	return fibMap
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	n := 10
	fibMap := fibonacciMap(n)

	for i, num := range fibMap {
		fmt.Printf("Fibonacci(%d) = %d\n", i, num)
	}
}
