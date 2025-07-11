package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	// Greet users
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// Booking loop
	for remainingTickets > 0 && len(bookings) < conferenceTickets {
		firstName, lastName, email, userTickets := getUserInput()

		// Validate input
		if validateUserInput(firstName, lastName, email, userTickets, remainingTickets) {
			remainingTickets, bookings = bookTicket(firstName, lastName, email, userTickets, remainingTickets, bookings)

			// Print first names of all bookings
			firstNames := printFirstNames(bookings)
			fmt.Printf("The first names of the bookings are: %v\n\n", firstNames)

			// If tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out, come back next year.")
				break
			}
		}
	}
}

// Greets the user
func greetUsers(confName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application.\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.\n")
}

// Gets user input
func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// Books the ticket and returns updated values
func bookTicket(firstName string, lastName string, email string, userTickets uint, remainingTickets uint, bookings []string) (uint, []string) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("\nThank you %v %v for booking %v ticket(s).\n", firstName, lastName, userTickets)
	fmt.Printf("A confirmation email will be sent to %v.\n", email)
	fmt.Printf("%v ticket(s) remaining.\n\n", remainingTickets)

	return remainingTickets, bookings
}

// Extracts first names from bookings
func printFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		if len(names) > 0 {
			firstNames = append(firstNames, names[0])
		}
	}
	return firstNames
}
