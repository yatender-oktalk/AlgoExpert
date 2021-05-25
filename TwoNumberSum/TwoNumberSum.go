package main

import "fmt"

var arr = []int{3, 5, -4, 8, 11, 1, -1, 6}

func main() {
	targetSum := 10
	x := TwoNumSum(arr, targetSum)
	fmt.Println(x)
}

func TwoNumSum(arr []int, targetSum int) []int {
	convertedMap := ConvertedMap(arr)
	for _, v := range arr {
		requiredNum := targetSum - v
		if convertedMap[requiredNum] && requiredNum != v {
			return []int{requiredNum, v}
		}
	}

	return []int{}
}

func ConvertedMap(arr []int) map[int]bool {
	m := make(map[int]bool)

	for _, v := range arr {
		m[v] = true
	}

	return m
}
