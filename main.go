package main

import (
	"fmt"
	"strconv"
	"strings"
)

// package level variables, available to all functions
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) //create a empty list of maps. 0 will be expanded dynamically

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicket {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of booking are %v\n", firstNames)

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

func greetUsers() {
	fmt.Printf("Welcome to the, %v booking application\n", conferenceName)
	fmt.Printf("We have a total of, %v, tickets and, %v, are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string { //outputs contains only the type
	var firstNames []string
	for _, booking := range bookings {
		//names := strings.Fields(x) // Split the string into a slice of strings
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func validateUserInput(firstName, lastName, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a empty map per user. We cannot mix data types in a map
	userData := make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) //convert uint to string

	bookings = append(bookings, userData)
	fmt.Printf("List of blookings is %v\n", bookings)

	fmt.Printf("Thanks %v %v for booking %v tickets. We have sent a confirmation email to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
