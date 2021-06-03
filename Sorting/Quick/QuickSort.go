package main

import (
	"fmt"
)

var arr = []int{8, 5, 2, 9, 5, 6, 3}

func main() {
	fmt.Println(QuickSort(arr))
}

func QuickSort(a []int) []int {
	quickSort(a, 0, len(a)-1)
	return a
}

func Partition(arr []int, low, high int) int {
	if low >= high {
		return arr
	}
	fmt.Println(arr)
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i < high {
			i++
		}

		for arr[j] > pivot && j > low {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[j] = arr[j], pivot
	return j
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	var pivot = PartitionO(arr, low, high)
	quickSort(arr, low, pivot)
	quickSort(arr, pivot+1, high)

}

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := 5 % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left+1:])

	return a
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

func PartitionO(arr []int, low, high int) int {
	var pivot = arr[low]
	var i = low
	var j = high

	for i < j {
		for arr[i] <= pivot && i < high {
			i++
		}
		for arr[j] > pivot && j > low {
			j--
		}

		if i < j {
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[low], arr[j] = arr[j], pivot

	return j
}
