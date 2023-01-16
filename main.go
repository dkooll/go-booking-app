package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to, %v booking application\n", conferenceName)
	fmt.Printf("We have a total of, %v, tickets and, %v, are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thanks you %v %v for booking %v tickets. We have sent a confirmation email to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets remaining\n", remainingTickets)

	fmt.Printf("These are the people who have booked so far: %v\n", bookings)
}
