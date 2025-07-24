package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}


	// %T is a type placeholder
	fmt.Printf("conferenceTickets is %T, remaningsTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// %v only works with Printf
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// setting an array - size, type

	for {
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
		
		if userTickets <= int(remainingTickets) {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaning for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings { 				// underscore identifies unused variable
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// function to end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
		}
	}
}