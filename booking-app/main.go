package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && len(email) > 5
		isValidTicketNo := userTickets > 0 && userTickets <= remainingTickets

		// isValidCity := city == "Rio De Janeiro" || city == "Stockholm"

		if isValidName && isValidEmail && isValidTicketNo {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v.\n\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n\n\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out, come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email format entered is incorrect")
			}
			if !isValidTicketNo {
				fmt.Println("Number of tickets enterd is invalid")
			}

		}

	}

	city := "Cali"

	switch city {
	case "Cali", "Rio de Janeiro":
		// execute code for booking Cali and Rio conference tickets
	case "Doha":
		// execute code for booking London conference tickets
	case "Nairobi":
		// some code here
	case "Stockholm":
		// some cold here
	case "Tokyo", "Taiwan", "Beijin":
		// some code here
	case "Ottawa", "Denver":
		// some code here
	default:
		fmt.Println("No valid city selected")
	}

}
