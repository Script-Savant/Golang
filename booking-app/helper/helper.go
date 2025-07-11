package helper

import (
	"fmt"
	"strings"
	"time"
)

// Validates user input
func ValidateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint) bool {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && len(email) > 5
	isValidTicketNo := userTickets > 0 && userTickets <= remainingTickets

	if isValidName && isValidEmail && isValidTicketNo {
		return true
	}

	// Show error messages
	if !isValidName {
		fmt.Println("❌ First or last name entered is too short.")
	}
	if !isValidEmail {
		fmt.Println("❌ Email format entered is incorrect.")
	}
	if !isValidTicketNo {
		fmt.Println("❌ Number of tickets entered is invalid or exceeds availability.")
	}
	fmt.Println()

	return false
}

func SendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n %v \nTo email address: %v\n", ticket, email)
	fmt.Println("######################")

}
