package main

import (
	"fmt"
	"strings"
	"test-app/helper"
)

func stringOperation() {
	var i int = 0

	fullnames := []string{}
	for {
		fmt.Println("Enter Name")
		var name string
		fmt.Scan(&name)

		fmt.Println("Enter Last Name")
		var lastName string
		fmt.Scan(&lastName)

		fullnames = append(fullnames, name+" "+lastName)
		printFirstName(fullnames)

		i++
		if i == 3 {
			break
		}
	}

}

func printFirstName(fullnames []string) {
	firstNames := []string{}
	for _, fullname := range fullnames {
		names := strings.Fields(fullname)
		firstNames = append(firstNames, names[0])
	}

	fmt.Printf("THe first names of bookings are: %v \n", firstNames)
}

func testBasics1() {

	fmt.Printf("Addition of 5 + 7 = %v", helper.Add(5, 7))

	//=======
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Printf("conferenceTickets is - %T, remainingTickets is %T, conferenceName is %T", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to ", conferenceName, "conference booking application")
	fmt.Printf("We have remaining %v Tickets \n", remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Println(conferenceName)

	var userName string
	var userTickets int

	//ask user their name
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tikcets.\n", userName, userTickets)

	var firstName string
	fmt.Println("Name please")
	fmt.Scan(&firstName)
	fmt.Println(firstName, "is good boy")

	//
	//var bookings = [50]string{"abc","cde"}
	var bookings [50]string

	bookings[0] = firstName
	bookings[1] = "Ajay"

	fmt.Printf("\nComplete Array is %v", bookings)
	fmt.Printf("\nComplete Array type %T", bookings)
	fmt.Printf("\nComplete Array length is %v", len(bookings))
	//fmt.Printf("\n2nd Name is %v", bookings[1])

	var number int
	//fmt.Println("Enter number")
	//fmt.Scan(&number)

	fmt.Printf("\nName is %v\n", bookings[number])

	bookings2 := []string{}
	bookings2 = append(bookings2, "taj")
	fmt.Print(bookings2)

}

func main() {
	testBasics1()
	//stringOperation()
}
