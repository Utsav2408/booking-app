package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {
	greetUsers()
	for remainingTickets > 0 && len(bookings) < 50 {

		userFirstName, userLastName, userEmail, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userFirstName, userLastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, userFirstName, userLastName, userEmail)
			// call function for preprocess
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			// noTicketsRemaining  := remainingTickets == 0
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Comeback next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email Address entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of Tickets you entered is Invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	//For Each Loop
	firstNames := []string{}
	for _, booking := range bookings { // _ is a blank identifier. It marks those variables that are not used
		firstNames = append(firstNames, booking["userFirstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userFirstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&userLastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets required: ")
	fmt.Scan(&userTickets)

	return userFirstName, userLastName, userEmail, userTickets
}

func bookTicket(userTickets uint, userFirstName string, userLastName string, userEmail string) {
	remainingTickets -= userTickets

	// create a map for a user
	var userData = make(map[string]string)
	userData["userFirstName"] = userFirstName
	userData["userLastName"] = userLastName
	userData["userEmail"] = userEmail
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
