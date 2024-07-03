package main


func Remaining_Ticket(userTickets uint) uint {
	// remaining tickets
	RemainingTicket = RemainingTicket - uint(userTickets)
	return RemainingTicket
}