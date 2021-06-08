package main

import "fmt"

func main() {
	var n = 10
	var maxSteps = 3
	fmt.Println(Staircase(n, maxSteps))
}

func Staircase(n int, maxSteps int) int {
	// user can take 1 step or 2 step not more than that
	if n <= 1 {
		return 1
	}

	var x int
	for i := 1; i <= min(maxSteps, n); i++ {
		x += Staircase(n-i, maxSteps)
	}

	return x
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
