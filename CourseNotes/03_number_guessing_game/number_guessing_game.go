package main

import (
	"fmt"
	"math/rand"
)

func main() {

	target := rand.Intn(100) + 1

	var guess int

	attempts := 0

	fmt.Println("*** The magical number guessing game ***")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++

		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Printf("Congrats! You guessed it in %d attempts.\n", attempts)
			break
		}
	}
}
