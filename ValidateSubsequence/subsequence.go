package main

import "fmt"

var arr = []int{5, 1, 22, 25, 6, -1, 8, 10}

var sub = []int{1, 6, -1, 8}

func main() {
	var resp bool = IsValidSubsequence(arr, sub)
	fmt.Println(resp)
}

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	var i, j int
	for i < len(array) && j < len(sequence) {
		if array[i] == sequence[j] {
			j++
		}
		i++
	}
	return j == len(sequence)
}
