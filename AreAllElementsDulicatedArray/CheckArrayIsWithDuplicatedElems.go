package main

import "fmt"

var arr = []int{4, 3, 3, 2, 4, 2}

func main() {
	fmt.Println(IsArrayDuplicated(arr))
}

// Here second approach can be just to store
// the values into the hashmap and check the values

// But in this alternative approach we are doing the xor
// so that final result will be the non duplicate number
func IsArrayDuplicated(arr []int) bool {
	var result int
	for i := 0; i < len(arr); i++ {
		result = result ^ arr[i]
	}
	return result == 0
}
