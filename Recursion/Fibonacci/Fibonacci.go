package main

import "fmt"

func main() {
	n := 6
	fmt.Println(fib(n))
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if n < 1 {
		return -1
	}

	return fib(n-1) + fib(n-2)
}
