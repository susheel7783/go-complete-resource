package main

import (
	"fmt"   // Package for formatted I/O (printing and scanning)
	"time"  // Package for working with dates and times
)

// Define a custom type (struct) to represent a user
type user struct {
	firstName string    // User's first name
	lastName  string    // User's last name
	birthDate string    // User's birth date as a string
	createdAt time.Time // Timestamp when user was created
}

// ==================== METHOD (attached to user struct) ====================
// This is a METHOD, not just a regular function
// (u user) is called a "receiver" - it attaches this function to the user type
// This means any user struct can call this method like: myUser.outputUserDetails()
// The method receives a COPY of the user struct (value receiver)
func (u user) outputUserDetails() {
	// Access the struct fields directly through 'u' (the receiver)
	// No need to pass parameters - the method already has access to the struct data
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func main() {
	// Collect user input for each field
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Declare a variable of type user
	var appUser user
	
	// Initialize the user struct with collected data
	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(), // Automatically capture current timestamp
	}

	// ... do something awesome with that gathered data!
	
	// ==================== CALLING THE METHOD ====================
	// Call the method directly on the appUser object
	// Syntax: objectName.methodName()
	// This is more object-oriented - the user "knows how to display itself"
	appUser.outputUserDetails()
}

// Reusable helper function to get user input with a custom prompt
func getUserData(promptText string) string {
	fmt.Print(promptText)  // Display the prompt
	var value string       // Variable to store input
	fmt.Scan(&value)       // Read input from console
	return value           // Return the value to caller
}
