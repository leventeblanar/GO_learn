package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func readInt(prompt string) (int, error) {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	line = strings.TrimSpace(line)
	return strconv.Atoi(line)
}


func readYN(prompt string) (bool, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader((os.Stdin))
	line, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	line = strings.TrimSpace(line)
	return line == "y" || line == "yes", nil

}

func play_game() {
	random_num := rand.IntN(100)
	attempts := 0

	for attempts < 10 {
		user_guess, err := readInt("Enter a guess: ")
		attempts ++
		if err != nil {
			fmt.Printf("Error: %s", err)
		}

		if int(user_guess) > random_num {
			fmt.Println("Too High")
		} else if int(user_guess) < random_num {
			fmt.Println("Too Low")
		} else {
			fmt.Println("YOU GOT IT!")
			fmt.Printf("It only took you %d to get it", attempts)
			fmt.Println("")
			break
		}
	}
}

func main() {
	fmt.Println("***** Number Guessing Game *****")
	fmt.Println("***** You'll get a random number and have 10 attempts to guess it. *****")

	for {
		play_game()
		user_resp, err := readYN("Do you want to play again? (y/n) ")
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		if user_resp == false {
			fmt.Println("Goodbye!")
			break
		}
	}
}
