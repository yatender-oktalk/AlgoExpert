package main

import "fmt"

var arr = []int{3, 5, -4, 8, 11, 1, -1, 6}

func main() {
	fmt.Println(BubbleSort(arr))
}

func BubbleSort(arr []int) []int {
	// code here
	for i := len(arr) - 1; i > 1; i-- {
		for j := i; j < i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
