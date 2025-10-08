package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) 	//standard inputból olvasás
	fmt.Print("Enter your full name: ")

	name, err := reader.ReadString('\n')		// addig olvas, amíg ENTER-t nem nyomsz
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	name = strings.TrimSpace(name) // levágjuk a \n és egyéb whitespace karaktereket

	if name == "" {
		fmt.Println("You didn't enter a name!")
		return
	}

	fmt.Println("Hello,", name)
}