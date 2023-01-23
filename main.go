package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) //empty list of UserData struct. 0 will be expanded dynamically

type UserData struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicket {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

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
		wg.Wait()
	}
}

func greetUsers() {
	fmt.Printf("Welcome to the, %v booking application\n", conferenceName)
	fmt.Printf("We have a total of, %v, tickets and, %v, are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
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

	firstName = strings.TrimSpace(firstName)

	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)

	lastName = strings.TrimSpace(lastName)

	fmt.Println("Please enter your email address")
	fmt.Scan(&email)

	email = strings.TrimSpace(email)

	fmt.Println("Enter the number of tickets you want to book")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	userData := UserData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of blookings is %v\n", bookings)

	fmt.Printf("Thanks %v %v for booking %v tickets. We have sent a confirmation email to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// simulate
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \nto mail address %v\n", ticket, email)
	fmt.Println("################")
	wg.Done()
}
