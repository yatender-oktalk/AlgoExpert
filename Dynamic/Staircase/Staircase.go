package main

import "fmt"

var stairMemoized = make(map[string]int)

func main() {
	var n = 10
	var maxSteps = 2
	fmt.Println(Staircase(n, maxSteps))
}

func Staircase(n int, maxSteps int) int {
	key := getKey(n, maxSteps)
	if keyAvailableInMap(key) {
		return stairMemoized[key]
	}
	// user can take 1 step or 2 step not more than that
	if n <= 1 {
		return 1
	}

	var x int
	for i := 1; i <= min(maxSteps, n); i++ {
		x += Staircase(n-i, maxSteps)
	}

	stairMemoized[key] = x
	return x
}

func getKey(n, m int) string {
	return fmt.Sprintf("%d-%d", n, m)
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func keyAvailableInMap(key string) bool {
	_, isKeyAvailable := stairMemoized[key]
	return isKeyAvailable
}
