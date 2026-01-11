package main // Main package - entry point for executable

import (
	"fmt"  // For printing and scanning input
	"time" // For working with dates and times
)

// Struct definition - blueprint for creating user objects
type user struct {
	// Custom data type that groups related fields together
	// Like a template/class that defines what a "user" contains
	
	firstName string    // Field: stores user's first name
	lastName  string    // Field: stores user's last name
	birthDate string    // Field: stores birthdate as text
	createdAt time.Time // Field: stores when user was created (timestamp)
	// time.Time is a built-in type from the "time" package
}

func main() {
	// Collect user input
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	
	// ... do something awesome with that gathered data!
	// Currently just passes data to output function
	// TODO: Could create a user struct instance here
	
	outputUserDetails(lastName, firstName, birthdate)
	// Calls output function (note: parameters are SWAPPED - lastName first!)
}

// Function to display user details
func outputUserDetails(firstName, lastName, birthdate string) {
	// Parameters: three strings (firstName, lastName, birthdate)
	// firstName, lastName string - shorthand when types are same
	// ...
	fmt.Println(firstName, lastName, birthdate)
	// Just prints the values
	// TODO: Could create and display user struct here
}

// Reusable function to get user input
func getUserData(promptText string) string {
	fmt.Print(promptText)  // Display prompt
	var value string       // Variable to store input
	fmt.Scan(&value)       // Read input (stops at whitespace!)
	return value           // Return the input
}
