package main

import (
	"fmt"
)

func main() {
	fmt.Println(pow(5, 3))
}

func pow(num, power int) int {
	multiplySum := num

	if power == 0 {
		return 1
	} else {
		for i := 0; i < power; i++ {
			multiplySum = multiply(multiplySum, num)
		}
	}
	return multiplySum
}

func multiply(num, multiplier int) int {
	sum := 0
	for i := 0; i < multiplier; i++ {
		sum += num
	}
	return sum
}
