package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	fmt.Println("Welcome to", conferenceName)

	// var bookings [50]string //array
	var bookings []string //slice
	
	var userName string
	var tickets int
	
	fmt.Println("What is your name?")
	fmt.Scan(&userName)
	fmt.Println("How many tickets are you looking for?")
	fmt.Scan(&tickets)
	
	fmt.Printf("Dear, %v, you have booked %v tickets\n", userName, tickets)
	
	// bookings[0] = userName
	bookings = append(bookings, userName)

	fmt.Printf("Your name from array elements is %v\n", bookings[0])
	fmt.Printf("Slice length: %v\n", len(bookings))
}
