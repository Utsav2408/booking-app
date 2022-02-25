package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	//Array
	// var bookings = [50]string{}
	// var bookings [50]string
	//Slice
	var bookings []string
	// var bookings = []string{}
	// bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	//Infinite loop
	for remainingTickets > 0 && len(bookings) < 50 {
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

		if userTickets <= remainingTickets {
			remainingTickets -= userTickets
			// bookings[0] = userFirstName + " " + userLastName
			bookings = append(bookings, userFirstName+" "+userLastName)

			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The array type: %T\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("The array length: %v\n", len(bookings))

			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The slice type: %T\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("The slice length: %v\n", len(bookings))

			fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			//For Each Loop
			firstNames := []string{}
			for _, booking := range bookings { // _ is a blank identifier. It marks those variables that are not used
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			// noTicketsRemaining  := remainingTickets == 0
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Comeback next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you cannot book %v tickets\n", remainingTickets, userTickets)
			// continue
		}
	}
}
