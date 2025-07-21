package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	// single-line comments start with "//"
	fmt.Println("Starting Textio server")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Nice to meet you, %s!\n", name)

	fmt.Println("Do you want to (1) hear a joke or (2) exit?")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "1" {
		tellJoke()
	} else {
		fmt.Println("Bye!")
	}
}

func tellJoke() {
	fmt.Println("Why don't scientist trust atoms?")
	fmt.Println("Because they make up everything!")
}