package main

import (
	"fmt"
)

var arr = []int{8, 5, 2, 9, 5, 6, 3}

func main() {
	fmt.Println(MyQuickSort(arr))
}

func MyQuickSort(a []int) []int {
	// if length is less than 2 then return the array to break the recursion
	if len(a) < 2 {
		return a
	}

	// if not then let's go ahead and choose our points

	leftIndex := 0
	rightIndex := len(a) - 1

	pivotIndex := rightIndex

	// after assigning the index we need to interate the array
	for i := range a {
		// swap with left and i value if you find any value less than pivot index
		// as we need to push the lesser values to the left
		if a[i] < a[pivotIndex] {
			a[leftIndex], a[i] = a[i], a[leftIndex]
			leftIndex++
		}
	}
	// left is the correct place for the pivot index or right index
	a[leftIndex], a[pivotIndex] = a[pivotIndex], a[leftIndex]

	MyQuickSort(a[:leftIndex])
	MyQuickSort(a[leftIndex+1:])

	return a
}
