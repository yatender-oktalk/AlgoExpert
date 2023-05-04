package array_algo

import "testing"

func TestArrayProduct(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		var arr []int = []int{5, 1, 4, 2}
		ArrayProduct(arr)
	})
}
