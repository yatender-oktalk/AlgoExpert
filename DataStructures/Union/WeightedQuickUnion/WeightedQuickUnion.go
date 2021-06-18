package main

import "fmt"

var sizeArr []int = make([]int, 0)

func main() {
	arr := Initialize(10)
	fmt.Println(arr)
}

func Initialize(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		sizeArr[i] = 1
	}
	return arr
}

func IsConnected(n, m int, arr []int) bool {
	return root(n, arr) == root(m, arr)
}

func WeightedUnion(a, b int, arr []int) []int {
	i := root(a, arr)
	j := root(b, arr)

	if i == j {
		return arr
	}

	// make smaller root point to larger one
	if sizeArr[i] < sizeArr[j] {
		arr[i] = j
		sizeArr[j] += sizeArr[i]
	} else {
		arr[j] = i
		sizeArr[i] += sizeArr[j]
	}

	return arr
}
func root(n int, arr []int) int {
	for arr[n] != n {
		n = arr[n]
	}
	return n
}
