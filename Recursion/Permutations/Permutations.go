package main

var arr []int = []int{1, 2, 3}
var result = make([][]int, 0)

func main() {
	GetPermutations(arr)
}

func GetPermutations(array []int) [][]int {
	GetPermutationsHelper(array, result)
	return nil
}

func GetPermutationsHelper(array []int, results [][]int) [][]int {
	if len(array) == 0 {
		return results
	}

	for i := 0; i < len(array); i++ {

	}
}
