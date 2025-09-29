package main

import "fmt"

func main() {
	// klasszikus for
	for i := 1; i <= 5; i++ {
		fmt.Println("i:", i)
	}

	// while-szerű
	j := 3
	for j > 0 {
		fmt.Println("j:", j)
		j--
	}

	// végtelen ciklus (break-kel kilépünk)
	k := 0
	for {
		k++
		if k == 3 {
			fmt.Println("break at k=3")
			break
		}
	}
}