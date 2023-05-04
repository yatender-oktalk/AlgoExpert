package array_algo

import "fmt"

var arr []int = []int{5, 1, 4, 2}

func ArrayProduct(arr []int) []int {
	// Final Product array

	productArr := make([]int, len(arr))
	leftArr := make([]int, len(arr))
	rightArr := make([]int, len(arr))

	// Initialize everything to 1

	for i := range arr {
		productArr[i] = 1
		leftArr[i] = 1
		rightArr[i] = 1
	}

	leftRunningProduct := 1
	// let's store the left product in left arr
	for i := range arr {
		leftArr[i] = leftRunningProduct
		leftRunningProduct *= arr[i]
	}

	rightRunningProduct := 1
	// similarly define the right array from reverse pointer
	for i := len(arr) - 1; i >= 0; i-- {
		rightArr[i] = rightRunningProduct
		rightRunningProduct *= arr[i]
	}

	for i := range arr {
		productArr[i] = rightArr[i] * leftArr[i]
	}
	return productArr
}
