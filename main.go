package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conf"
var remainingTickets uint = 50
var bookings = make([]Booking, 0)

type Booking struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(
			firstName, lastName, email, userTickets,
		)

		if isValidName && isValidEmail && isValidTicketNumber {
			booking := bookTicket(userTickets, firstName, lastName, email)
			printBookingsConference(booking)

			wg.Add(1)
			go sendTicket(booking)

			firstNames := getFirstNames()
			fmt.Printf("The FirstNames are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conf is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("FirstName or LastName is too short")
			}
			if !isValidEmail {
				fmt.Println("Email is incorrect")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets invalid")
			}
		}
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	var firstNames []string

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) Booking {
	remainingTickets -= userTickets

	var booking = Booking{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, booking)

	return booking
}

func sendTicket(booking Booking) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", booking.numberOfTickets, booking.firstName, booking.lastName)

	fmt.Println("##########################")
	fmt.Printf("Sending ticket %v\nTo email address: %v\n", ticket, booking.email)
	fmt.Println("##########################")

	defer wg.Done()
}

func printBookingsConference(booking Booking) {
	fmt.Printf("List of booking is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", booking.firstName, booking.lastName, booking.numberOfTickets, booking.email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
