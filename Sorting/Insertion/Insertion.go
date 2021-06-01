package main

import "fmt"

var arr = []int{3, 5, -4, 8, 11, 1, -1, 6}

func main() {
	fmt.Println(InsertionSortSol1(arr))
}

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		item := arr[i]
		j := i - 1
		for ; j >= 0 && item < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		fmt.Println(j)
		arr[j+1] = item
	}

	return arr
}

func InsertionSortSol1(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}
