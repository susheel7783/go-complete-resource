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

// ==================== METHOD WITH POINTER RECEIVER ====================
// (u *user) is a POINTER RECEIVER - notice the asterisk *
// This means the method receives a reference to the original struct, not a copy
// Used here for reading data (doesn't modify, but pointer receivers are common practice)
func (u *user) outputUserDetails() {
	// Access struct fields through the pointer (Go automatically dereferences)
	// We don't need to write (*u).firstName - Go does this for us
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// ==================== METHOD THAT MODIFIES THE STRUCT ====================
// (u *user) is a POINTER RECEIVER - REQUIRED for modifying the struct
// Without the pointer (*), changes would only affect a copy, not the original
func (u *user) clearUserName() {
	// Modify the actual struct fields by setting them to empty strings
	// This changes the ORIGINAL struct because we have a pointer receiver
	u.firstName = ""  // Clear first name
	u.lastName = ""   // Clear last name
	// Note: birthDate and createdAt are NOT cleared
}

func main() {
	// Collect user input for each field
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Declare and initialize a user struct
	var appUser user
	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!
	
	// ==================== DEMONSTRATING METHOD CALLS ====================
	
	// Call 1: Display user details BEFORE clearing
	// Output: Will show the firstName, lastName, and birthDate we entered
	appUser.outputUserDetails()
	
	// Call 2: Clear the user's name fields
	// This MODIFIES the original appUser struct because clearUserName uses pointer receiver
	appUser.clearUserName()
	
	// Call 3: Display user details AFTER clearing
	// Output: Will show empty strings for firstName and lastName, but birthDate remains
	appUser.outputUserDetails()
}

// Reusable helper function to get user input with a custom prompt
func getUserData(promptText string) string {
	fmt.Print(promptText)  // Display the prompt (no newline)
	var value string       // Variable to store user input
	fmt.Scan(&value)       // Read input from console into value
	return value           // Return the collected input
}
