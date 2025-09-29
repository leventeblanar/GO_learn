package main

import "fmt"

func devisible_by_3(n int) string {
	if n%3 == 0 {
		return fmt.Sprintf("The number %d is divisible with 3.", n)
	} else {
		return fmt.Sprintf("The number %d is not divisible with 3.", n)
	}
}

func main() {
	n := 7

	if n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}


	if n := 10; n < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("non-negative")
	}


	//  calling the above function
	fmt.Println(devisible_by_3(6))
	fmt.Println(devisible_by_3(5))
}