package main

import (
    "fmt"
    "strings"
)

func Validating (firstName string, lastName string, userEmail string, userTickets uint, RemainingTicket uint) {
	// Validating user input
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicket := userTickets > RemainingTicket
	isValidData := isValidName && isValidEmail && isValidTicket

	if !isValidData {
		if !isValidName {
			fmt.Print("You are providing wrong name either in first name or in last name\n")
		}
		if !isValidEmail {
			fmt.Print("You are provding wrong email address also. It must be contains @ sign\n")
		}
		if isValidTicket {
			fmt.Printf("You want more than %d tickets but we only have %d tickets for this conference\n", userTickets, ConferenceTicket)
		}
	}
}