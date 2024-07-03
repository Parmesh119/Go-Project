package main

import (
    "fmt"
    "strings"
)

func Firstname () {
	// Adding only first name of the person who booking the tickets
	Firstname_booking := []string{}
	for _, name := range bookings {
		var name = strings.Fields(name)
		Firstname_booking = append(Firstname_booking, name[0])
	}
	fmt.Print(Firstname_booking, "\n")
}