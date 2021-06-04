package main

var arr = []int{1, 0, 0, -1, -1, 0, 1, 1}

var order = []int{0, 1, -1}

func main() {
	SortNew(arr)
}

func SortThreeNumList(arr []int, order []int) []int {
	// let's take 2 pointers
	onePointer := 0
	threePointer := len(arr) - 1

	for i := 0; onePointer <= threePointer && i < len(arr)-1; i++ {
		if arr[i] == order[0] {
			arr[onePointer], arr[i] = arr[i], arr[onePointer]
			onePointer++
		} else if arr[i] == order[2] {
			arr[threePointer], arr[i] = arr[i], arr[threePointer]
			threePointer--
		}
	}

	return arr
}

func SortNew(arr []int) []int {
	threePointer := len(arr) - 1
	onePointer := 0

	i := 0

	for onePointer < threePointer && i <= threePointer {
		if arr[i] == order[0] {
			//swap with the one pointer
			arr[i], arr[onePointer] = arr[onePointer], arr[i]
			onePointer++
			i++
		} else if arr[i] == order[2] {
			//swap with the one pointer
			if arr[i] == order[1] {
				arr[i], arr[threePointer] = arr[threePointer], arr[i]
				i++
			} else {
				arr[i], arr[threePointer] = arr[threePointer], arr[i]
			}
			threePointer--
		} else if arr[i] == order[1] {
			i++
		}
	}
	return arr
}
