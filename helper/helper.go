package helper

import "strings"

func ValidateUserInput(userFirstName string, userLastName string, userEmail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	return len(userFirstName) >= 2 && len(userLastName) >= 2, strings.Contains(userEmail, "@"), userTickets > 0 && userTickets <= remainingTickets
}
