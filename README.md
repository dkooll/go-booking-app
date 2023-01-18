# Go Booking App

This is a Go program that simulates a ticket booking application for a conference. The program starts by  
greeting the user and displaying the number of available tickets.

Then, it enters a loop that prompts the user for their first name, last name, email, and number of tickets  
they want to book. The user input is then validated, and if it is valid, the tickets are booked and a  
confirmation email is sent to the user in a separate goroutine.  

If the user input is not valid, an error message is displayed. The program also keeps track of all bookings  
and the first names of the bookers in a slice of UserData structs. The program ends when all the tickets  
have been sold out.

## Overview
### Function Flow

![Function Flow](/images/functionflow.png "function flow")

## Details

The code starts by initializing package-level variables. This includes a constant for the total number of tickets and a couple of variables to keep track of the remaining tickets, the conference name, and an empty list of
UserData structs.

The UserData struct stores information about the user booking tickets such as their first name, last name, email address, and the number of tickets they are purchasing.

The main() function is where the application logic is. It starts with a greetUsers() function, which prints out a welcome message to the user. It then enters a loop to get user input from the command line.

The getUserInput() function gets the user's first name, last name, email address, and the number of tickets they want to book.

The validateUserInput() function checks if the user's input is valid. It checks if the first and last name are not too short, if the email address is valid, and if the number of tickets requested is valid.

The bookTicket() function is used to book the tickets. It subtracts the number of tickets requested from the remaining tickets and adds the user's information to a list of UserData
structs. It also prints out a confirmation message and the remaining tickets.

The getFirstNames() function is used to get the first names of all the users who have booked tickets.

The sendTicket() function is used to simulate sending a ticket to the user. It prints out a message and waits 10 seconds before finishing.

Finally, a WaitGroup is used to add concurrency to the application, so other threads can run while the sendTicket() function is running.

## Concepts Learned

- Using package-level variables and constants
- Using the sync package to add concurrency to an application
- Using the time package to add a delay to a function
- Making use of structs and structs fields to store user data
- Creating dynamic slices
- Using waitgroups to keep track of the number of goroutines running, allowing the main thread to wait for them to finish
- Process function inputs and return values