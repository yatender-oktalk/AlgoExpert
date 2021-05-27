package main

import "sort"

var array = []int{1, 2, 3, 5, 6, 8, 9}

func main() {
	SortedSquaredArray(array)
}

func SortedSquaredArray(array []int) []int {
	a := []int{}

	for _, v := range array {
		a = append(a, v*v)
	}
	sort.Ints(a)
	return a
}
