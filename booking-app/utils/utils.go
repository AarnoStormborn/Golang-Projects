package utils

import "fmt"

var ConferenceName = "Go Conference"

func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nPlease enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("\nPlease enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("\nPlease enter your email ID: ")
	fmt.Scan(&email)
	fmt.Println("\nHow many tickets do you want? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
