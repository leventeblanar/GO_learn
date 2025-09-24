package main

import (
	"fmt"
)

func makeCounter() func(int) int {
	sum := 0
	return func() int {
		sum += 1
		return sum
	}
}


func main() {
	counterA := makeCounter()
	counterB := makeCounter()

	fmt.Println(counterA())
	fmt.Println(counterA())
	fmt.Println(counterA())
	
	fmt.Println(counterB())
	fmt.Println(counterB())
}