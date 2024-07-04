package UserInput

import (
	"fmt"
)

func User_input() (string, string, string, uint) {
	// Taking user's name
	var firstName string
	fmt.Print("Enter your name: ")
	fmt.Scan(&firstName)

	var lastName string
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	// Taking email
	var userEmail string
	fmt.Print("Enter your email: ")
	fmt.Scan(&userEmail)

	// Taking tickets that user wants
	var userTickets uint
	fmt.Print("How many tickets that you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}
