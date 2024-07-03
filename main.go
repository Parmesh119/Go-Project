package main

import (
	"fmt"
)

// Welcoming user
var ConferenceName string = "Go Conference"

// Initializing tickets
const ConferenceTicket uint = 50
var RemainingTicket uint = ConferenceTicket

// Declaring Slice
var bookings = []string{}

func main() {

	// GREETING
	fmt.Print(greet())

	// showing details
	fmt.Printf("We have %d number of total ticket and remaining tickets are %v\n", ConferenceTicket, RemainingTicket)

	// Loop for booking tickets
	for i := 0; i < int(ConferenceTicket); i++ {

		// User input
		firstName, lastName, userEmail, userTickets := User_input()

		// Remaining Tickets
		Remaining_Ticket := Remaining_Ticket(userTickets)

		// Validating user
		Validating(firstName, lastName, userEmail, userTickets, Remaining_Ticket)

		// Printing details
		fmt.Printf("The name of the user is %s and %s is asking for %d tickets of this conference\n", firstName, firstName, userTickets)

		// Slicing
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Println(bookings)

		// Firstname
		Firstname()

		fmt.Printf("Thanks %s for booking the tickets. You will soon receive a confirmation mail on provided email address %s\n", firstName, userEmail)
		fmt.Printf("The %d tickets are remaining in %s\n", RemainingTicket, ConferenceName)
	}

}
