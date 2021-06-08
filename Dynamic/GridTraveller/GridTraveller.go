package main

import "fmt"

var gridMemoized = make(map[string]int)

func main() {
	// Check brute force first
	fmt.Println(GridTravellerMemoized(5, 5))

}

// Brute Force
func GridTravellerBF(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}

	return GridTravellerBF(m-1, n) + GridTravellerBF(m, n-1)
}

func GridTravellerMemoized(m, n int) int {
	key := fmt.Sprintf("%d-%d", m, n)

	if keyAvailableInMap(key) {
		return gridMemoized[key]
	}

	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}

	gridMemoized[key] = GridTravellerMemoized(m-1, n) + GridTravellerMemoized(m, n-1)
	return gridMemoized[key]
}

func keyAvailableInMap(key string) bool {
	_, isKeyAvailable := gridMemoized[key]
	return isKeyAvailable
}
