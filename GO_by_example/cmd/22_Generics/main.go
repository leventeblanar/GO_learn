package main

import (
	"fmt"
)


// GO-ban ha akarsz írni egy függvényt,ami többféle típuson is működik
//			- írsz külön függvényt minden típusra
//			- vagy interface{}-et használsz, de az kényelmetlen és nem típusbitztos

//  Generics-szel írhatod ugyanazt a logikát általánosan, és a GO majd kitölti a konkrét típust


//  alább a külön megírt, generics nélküli
func sumInts(nums []int) int {
	total := 0
	for _,n := range nums {
		total += n
	}
	return total
}

func sumFloats(nums []float64) float64 {
	total := 0.0
	for _, n := range nums {
		total += n
	}
	return total
}

// Generics-el
func sumSlice[T int | float64](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	ints := []int{1, 2, 3}
	floats := []float64{1.5, 2.5, 3.5}

	fmt.Println(sumSlice(ints))
	fmt.Println(sumSlice(floats))
}