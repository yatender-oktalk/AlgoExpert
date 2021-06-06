package main

import "fmt"

var arr = []int{4, 3, 3, 2, 4}

func main() {
	fmt.Println(NonDuplicateNumber(arr))
}

// Here second approach can be just to store
// the values into the hashmap and check the values

// But in this alternative approach we are doing the xor
// so that final result will be the non duplicate number
func NonDuplicateNumber(arr []int) int {
	var result int
	for i := 0; i < len(arr); i++ {
		result = result ^ arr[i]
		fmt.Println(result)
	}
	return result
}
