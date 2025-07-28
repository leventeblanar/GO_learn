package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	for {
		var guess string
		guess = strings.ToLower(guess)

		fmt.Println("Guess if the number is even or odd. (type even, odd, or exit)")
		fmt.Scan(&guess)

		if guess == "exit" {
			fmt.Println("Thanks for playing.")
			break
		}

		if guess != "even" && guess != "odd" {
			fmt.Println("Invalid input. Please type 'even' , 'odd' or 'exit'.")
			continue
		}

		num := rand.Intn(100) + 1
		isEven := num%2 == 0

		if (guess == "even" && isEven) || (guess == "odd" && !isEven) {
			fmt.Printf("The number was %d. Good guess!\n", num)
		} else {
			fmt.Printf("The number was %d. Wrong guess!\n", num)
		}

	}
	
}
