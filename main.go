package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

//slice is an abstraction of an array
//slices are more flexible and powerful
//they are of variable length or get an sub array of its own
// slices are also index based and have a size, but is resized when needed

// declare a slice
var bookings = make([]userData, 50)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()
	for {
		//declaring a variable
		//arrays in go have fixed size and same data type has to be used

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		fmt.Printf("%v tickets are remaining for this conference\n", remainingTickets)

		if isValidName && isValidEmail && isValidNumber {
			bookTicket(firstName, lastName, email, conferenceName, remainingTickets, userTickets)
			firstNames := getFirstNames()
			fmt.Printf("thje first names of bookings are : %v\n", firstNames)
			go sendTicket(userTickets, firstName, lastName, email)
		}
		if remainingTickets == 0 {
			fmt.Println("The conference is sold out")
			break
		}
		if !isValidName {
			fmt.Println("You have entered an invalid name")
		}
		if !isValidEmail {
			fmt.Println("You have entered an invalid email")
		}
		if !isValidNumber {
			fmt.Println("You have entered an invalid number of tickets")
		}

	}
}

func greetUsers() {
	fmt.Printf("welcome to %v booking applications\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // in go _ are used to specify unsused variables
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var userTickets uint
	var firstName string
	var lastName string
	var email string
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("enter the number of tickets you want to purchase")
	fmt.Scan(&userTickets)
	fmt.Printf("user %v %vhas booked %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf(" your email id is %v \n", email)
	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, conferenceName string, remainingTickets uint, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	//create a map for the user
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("The list of bookins : %v", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets remaining for %v \n", remainingTickets, conferenceName)
}

//tasks to work on next -- slices, structs, pointers, go routines
// to store data of different types we use structs
// in my opinion structs is analogous to a json object of javascript

//structs stand for structure

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	fmt.Printf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("##################")
}
