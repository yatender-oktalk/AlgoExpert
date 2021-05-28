package main

import "fmt"

var arr = []int{1, 5, 8, 90, 99, 190}
var target = 90

func main() {
	result := BinarySearchIterative(arr, target)
	fmt.Println(result)
}

func BinarySearch(arr []int, target, low, hi int) int {
	if hi >= low {
		guess := (low + hi) / 2
		if target == arr[guess] {
			return guess
		} else {
			if arr[guess] > target {
				hi = guess - 1
			} else {
				low = guess + 1
			}
			return BinarySearch(arr, target, low, hi)
		}
	}
	return -1
}

func BinarySearchNumberAvailable(arr []int, target int) int {
	low := 0
	hi := len(arr) - 1

	if hi >= low {
		guess := low + (hi-low)/2
		if target == arr[guess] {
			return 1
		} else {
			if arr[guess] > target {
				arr = arr[:guess]
			} else {
				low = guess + 1
				arr = arr[low:]
			}
			return BinarySearchNumberAvailable(arr, target)
		}
	}
	return -1
}

func BinarySearchIterative(arr []int, target int) int {
	low := 0
	hi := len(arr) - 1

	for hi >= low {
		mid := low + (hi-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
