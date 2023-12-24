package main

import (
	"fmt"
	"sync"
	"time"
)

const availableTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserInfo, 0)

var wg = sync.WaitGroup{}

type UserInfo struct {
	firstName       string
	lastName        string
	email           string
	numberoftickets uint
}

func main() {
	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isvalidName, isvalidEmail, isvalidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isvalidName && isvalidEmail && isvalidTicketNumber {

		noremainingTickets := remainingTickets == 0
		if noremainingTickets {
			fmt.Println("Tickets are sold out ! Please try later")
		}

		bookTickets(firstName, lastName, userTickets, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

	} else {
		if !isvalidEmail {
			fmt.Print("Your email is invalid, try again \n")
		}
		if !isvalidName {
			fmt.Print("Your name is too short, try again \n")
		}
		if !isvalidTicketNumber {
			fmt.Print("Your number of tickets is invalid, try again \n")
		}

		//break
		//continue
	}

	wg.Wait()
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func greetUsers() {
	fmt.Printf("Welcome to the %v application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available to book\n", availableTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter the firstName :")
	fmt.Scan(&firstName)
	fmt.Println("Enter the lastName :")
	fmt.Scan(&lastName)
	fmt.Println("Enter the Email address :")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets needed:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, userTickets uint, email string) {

	remainingTickets = remainingTickets - userTickets

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = firstName
	// userData["email"] = email
	// userData["numberoftickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserInfo{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberoftickets: userTickets,
	}

	bookings = append(bookings, userData)

	firstNames := getFirstNames()

	fmt.Printf("Thank you %v %v for booking with us , a total of %v Tickets are booked, the confirmation is send to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets are remaining for the conference\n", remainingTickets)
	fmt.Printf("List of Booking are %v\n", bookings)
	fmt.Printf("These are first name of all the people who have bookings %v\n", firstNames)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Print("#########################################################\n")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Print("#########################################################\n")

	wg.Done()
}
