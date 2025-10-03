package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// generikus maxValue
func maxValue[T constraints.Ordered](nums []T) T {
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}


func main() {
	ints := []int{1, 7, 3, 5}
	floats := []float64{1.2, 5.4, 3.3}

	fmt.Println(maxValue(ints))
	fmt.Println(maxValue(floats))
}