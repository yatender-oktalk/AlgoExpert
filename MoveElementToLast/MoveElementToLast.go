package main

import "fmt"

var arr []int = []int{2, 1, 2, 2, 4, 3, 5, 2}

func main() {
	fmt.Println(MoveElementToEnd(arr, 2))
}

func MoveElementToEnd(arr []int, toMove int) []int {
	// We'll use 2 pointers 1 from start and 1 from end
	// if starting pointer match value we'll swap from end pointer and decrease the end pointer
	// if starting pointer doesn't match then just increase the start pointer
	// break when start pointer reaches to end pointer our value will be sorted

	startPointer := 0
	endPointer := len(arr) - 1

	for i := 0; startPointer < endPointer && i < len(arr); i++ {
		if arr[endPointer] == toMove {
			endPointer--
		} else if arr[startPointer] == toMove {
			arr[startPointer], arr[endPointer] = arr[endPointer], arr[startPointer]
			startPointer++
			endPointer--
		} else {
			startPointer++
		}
	}

	return arr
}
