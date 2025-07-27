package main // this is a program that can be run in solo

import (
	"fmt"
	"strings"
)

func main() {  // this is the enter point of the program
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{} // empty string slice

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// %T is a type placeholder
	fmt.Printf("conferenceTickets is %T, remaningsTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// setting an array - size, type

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets int
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)
		
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaning for %v\n", remainingTickets, conferenceName)

			// print first names
			printFirstNames((bookings))

			if remainingTickets == 0 {
				// function to end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The First name or Last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered is in a wrong format.")
			}
			if !isValidTicketNumber{
				fmt.Println("The number of tickets you entered is invalid.")
			}
		}
	}

	city := "London"

	switch city {
	case "New York":
		/// code
	case "Singapore", "London":
		/// code
	case "Berlin", "Mexico City":

	case "Hong Kong":

	default:
		fmt.Print("No valid city selected.")
	}
}


func greetUsers(confName string, confTickets int, remainTickets uint) {
	fmt.Printf("Welcome to our %v booking application.\n", confName)
	// %v only works with Printf
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remainTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings { 				// underscore identifies unused variable
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("These are all our bookings: %v\n", firstNames)
}