# Booking App

This code is a booking application for a Go conference. It starts by greeting the user and then prompts them for
input. It validates the user input and if it is valid, it books the ticket for the user, adds it to the list
of bookings, and prints out a confirmation.

It also prints out the remaining tickets and the list of booked tickets. Finally, it sends a ticket to the user's
email address.

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