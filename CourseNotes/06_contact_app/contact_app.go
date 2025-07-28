package main

import (
	"fmt"
)

func main() {

	fmt.Println("*** Contact Book application ***")

	contactInfos := map[string]string{
		"Anna": "+36 30 123 4567",
		"Dani": "+36 70 7654 321",
		"Lajos": "+36 70 6767 345",
	}

	for {
		var name string
		
		fmt.Println("Please enter a name you wish to look up for contact info.")
		fmt.Println("To quit, enter 'exit'.")
		fmt.Scan(&name)

		if name == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		val, ok := contactInfos[name]
		if ok {
			fmt.Printf("%v: %v \n", name, val)
		} else {
			fmt.Printf("There is no contact in our application with the name: %v \n", name)
		}
	}
}