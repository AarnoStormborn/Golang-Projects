package main

import (
	"booking-app/utils"
	"fmt"
	"sync"
	"time"
)

var bookings = make([]User, 0)

type User struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var waitGroup = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := utils.GetUserInput()

		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicket {

			remainingTickets = remainingTickets - userTickets

			var userData = User{
				firstName: firstName,
				lastName:  lastName,
				email:     email,
				tickets:   userTickets,
			}

			bookings = append(bookings, userData)
			fmt.Printf("\nList of Bookings: %v\n", bookings)

			waitGroup.Add(1)
			go sendTicket(userTickets, firstName, lastName)

			firstNames := printFirstNames()
			fmt.Printf("\nAll bookings: %v\n", firstNames)

			fmt.Printf("\nUser %v %v booked %v tickets\n", firstName, lastName, userTickets)
			fmt.Printf("\nRemaining Tickets: %v\n", remainingTickets)

		} else {

			if !isValidName {
				fmt.Println("Error in Name! Too short")
			}
			if !isValidEmail {
				fmt.Println("Invalid email!")
			}
			if !isValidTicket {
				fmt.Println("Invalid number of Tickets!")
			}
			continue
		}

		if remainingTickets == 0 {
			fmt.Printf("\nAll tickets are sold!\n")
			break
		}

		waitGroup.Wait()
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v Ticket Booking!\n", utils.ConferenceName)
	fmt.Printf("We have %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTicket(userTickets uint, firstName string, lastName string) {
	time.Sleep(10 * time.Second)

	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n############################")
	fmt.Printf("Sending ticket: %v\n", ticket)
	fmt.Println("############################")
	waitGroup.Done()
}
