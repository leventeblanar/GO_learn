package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Person struct {
	Name string
	Phone string
	Email string
}

func main() {

	reader := bufio.NewReader((os.Stdin))


	contacts := map[string]Person{
		"Anna": {Name: "Anna", Phone: "+36 30 123 4567", Email: "ana@gmail.com"},
		"Béla": {Name: "Béla", Phone: "+36 30 123 9876", Email: "Bela@gmail.com"},
		"Lajos": {Name: "Lajos", Phone: "+36 70 657 8493", Email: "laliba@gmail.com"},
	}

	fmt.Println("––– ContactBook application - version 2 –––")

	for {
		var choice string
		fmt.Println("Please choose whether you would like to query a name or add a new contact.")
		fmt.Println("To select the functionality, enter 'query' or 'add'. To quit, enter 'exit'.")
		fmt.Scan(&choice)

		if choice == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		switch choice {
		case "query":
			var name string

			fmt.Println("Enter a name you would like to look up.")
			fmt.Println("To quit, enter 'exit'.")
			fmt.Scan(&name)


			val, ok := contacts[name]
			if ok {
				fmt.Println("Contact found:")
				fmt.Println("---------------")
				fmt.Println("Name: ", val.Name)
				fmt.Println("Phone:", val.Phone)
				fmt.Println("Email:", val.Email)
			} else {
				fmt.Printf("There is no contact in our application with the name: %v \n", name)
			}
		case "add":
			var name, phone, email string
			fmt.Println("Enter the name of the contact: ")
			fmt.Scan(&name)

			fmt.Println("Phone number: ")
			phoneRaw, _ := reader.ReadString('\n')
			phone = strings.TrimSpace(phoneRaw)

			fmt.Println("Email: ")
			fmt.Scan(&email)

			contacts[name] = Person{Name: name, Phone: phone, Email: email}
			fmt.Printf("Contact %v has been added. \n", name)
		default:
			fmt.Println("Unkown option. Please type 'query', 'add', or 'exit'.")
		}
	}
}