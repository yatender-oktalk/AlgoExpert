package main

import (
	"fmt"
)

var arr = []int{2, 1, 5, 2, 3, 3, 4}

func main() {
	fmt.Println(FirstDuplicateValueSecond(arr))
}

func FirstDuplicateValue(array []int) int {
	arrMap := make(map[int]int)
	for _, v := range array {
		if arrMap[v] == 0 {
			arrMap[v] = 1
		} else {
			return v
		}
	}
	return -1
}

func FirstDuplicateValueSecond(array []int) int {
	repeating := -1
	for _, v := range array {
		// Here we will just make the value of that index -1*(value)
		val := getAbsolute(v) - 1

		if array[val] > 0 {
			array[val] = array[val] * -1
		} else {

			return v
		}
	}
	return repeating
}

func getAbsolute(x int) int {
	if x >= 0 {
		return x
	} else {
		return -1 * x
	}
}
