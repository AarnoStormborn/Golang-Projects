package main

import "strings"

var remainingTickets uint = 50

const conferenceTickets = 50

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}
