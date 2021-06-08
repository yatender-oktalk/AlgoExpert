package main

import "fmt"

func main() {
	arr := Initialize(10)
	fmt.Println(IsConnected(3, 4, arr))
	fmt.Println(arr)
	arr = Union(3, 4, arr)
	fmt.Println(arr)
	fmt.Println(IsConnected(3, 4, arr))
	arr = Union(3, 9, arr)
	arr = Union(3, 8, arr)
	fmt.Println(arr)

}

// This function will tell that is
func IsConnected(n, m int, arr []int) bool {
	return arr[n] == arr[m]
}

// This function will initialize the array where nodes are present
func Initialize(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func Union(n, m int, arr []int) []int {
	valM := arr[m]
	valN := arr[n]

	for i := 0; i < len(arr); i++ {
		if arr[i] == valN {
			arr[i] = valM
		}
	}
	return arr
}
