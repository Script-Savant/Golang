package main

import (
	"booking-app/helper"
	"fmt"
	"strconv" 
)

var conferenceName = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {
	

	// Greet users
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// Booking loop
	for remainingTickets > 0 && len(bookings) < conferenceTickets {
		firstName, lastName, email, userTickets := getUserInput()

		// Validate input
		if helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets) {
			remainingTickets = bookTicket(firstName, lastName, email, userTickets, remainingTickets)

			// Print first names of all bookings
			firstNames := printFirstNames()
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
	fmt.Println("Get your tickets here to attend.")
	fmt.Println("")
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
func bookTicket(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (uint) {
	remainingTickets -= userTickets

	var userData = make(map[string] string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is: %v\n", bookings)

	fmt.Printf("\nThank you %v %v for booking %v ticket(s).\n", firstName, lastName, userTickets)
	fmt.Printf("A confirmation email will be sent to %v.\n", email)
	fmt.Printf("%v ticket(s) remaining.\n\n", remainingTickets)

	return remainingTickets
}

// Extracts first names from bookings
func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
		
	}
	return firstNames
}
