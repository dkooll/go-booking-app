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

	fmt.Printf("Welcome to, %v booking application\n", conferenceName)
	fmt.Printf("We have a total of, %v, tickets and, %v, are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email address")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you want to book")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicket {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thanks %v %v for booking %v tickets. We have sent a confirmation email to %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("We have %v tickets remaining\n", remainingTickets)

			var firstNames []string
			for _, x := range bookings {
				names := strings.Fields(x) // Split the string into a slice of strings
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry we are sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email is invalid")
			}
			if !isValidTicket {
				fmt.Println("You have entered an invalid number of tickets")
			}
		}
	}
}
