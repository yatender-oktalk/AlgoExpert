package main

import "fmt"

func main() {
	arr := Initialize(10)
	fmt.Println(IsConnectedQuick(3, 4, arr))
	fmt.Println(arr)
	arr = UnionQuick(3, 4, arr)
	fmt.Println(arr)
	fmt.Println(IsConnectedQuick(3, 4, arr))
	arr = UnionQuick(3, 9, arr)
	fmt.Println(arr)
	fmt.Println(IsConnectedQuick(3, 9, arr))
	arr = UnionQuick(3, 8, arr)
	fmt.Println(IsConnectedQuick(3, 8, arr))
	fmt.Println(arr)

}

// This function will initialize the array where nodes are present
func Initialize(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

// Quick Union Improves
func UnionQuick(n, m int, arr []int) []int {
	i := root(n, arr)
	j := root(m, arr)
	arr[i] = j
	return arr
}

func IsConnectedQuick(n, m int, arr []int) bool {
	return root(n, arr) == root(m, arr)
}

// This private function will give back the root node
func root(n int, arr []int) int {
	for arr[n] != n {
		n = arr[n]
	}
	return n
}
