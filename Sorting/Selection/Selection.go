package main

import "fmt"

var arr = []int{3, 5, -4, 8, 11, 1, -1, 6}

func main() {
	fmt.Println(SelectionSort(arr))
}

func SelectionSort(arr []int) []int {
	//outer loop to iterate
	for i := 0; i < len(arr)-1; i++ {
		// inner loop
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				// swap
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
