package main // Main package - entry point for executable

import (
	"fmt"  // For printing and scanning input
	"time" // For working with dates and times
)

// Struct definition - blueprint for user objects
type user struct {
	// Custom type that groups user-related data together
	firstName string    // User's first name
	lastName  string    // User's last name
	birthDate string    // User's birthdate (stored as string)
	createdAt time.Time // Timestamp when user was created
}

func main() { // Main function - program starts here
	
	// Step 1: Collect user input
	userFirstName := getUserData("Please enter your first name: ")
	// Calls getUserData, stores result in userFirstName
	
	userLastName := getUserData("Please enter your last name: ")
	// Gets last name from user
	
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	// Gets birthdate from user
	
	// Step 2: Create struct instance
	var appUser user
	// Declares a variable of type 'user' (our custom struct)
	// appUser is currently empty (zero values: "", "", "", zero time)
	
	appUser = user{
		// Creates a new user struct and assigns it to appUser
		// Using struct literal syntax with field names
		
		firstName: userFirstName,
		// Sets firstName field to the input value
		// Field name: value syntax
		
		lastName:  userLastName,
		// Sets lastName field
		
		birthDate: userBirthdate,
		// Sets birthDate field
		
		createdAt: time.Now(),
		// Sets createdAt to current timestamp
		// time.Now() returns current date and time
	}
	
	// ... do something awesome with that gathered data!
	
	// ❌ BUG: This line has undefined variables!
	outputUserDetails(lastName, firstName, birthdate)
	// lastName, firstName, birthdate are NOT defined in this scope
	// Should be: userLastName, userFirstName, userBirthdate
	// OR better: pass the struct instead
	
	// ✅ Should be one of these:
	// Option 1: outputUserDetails(userFirstName, userLastName, userBirthdate)
	// Option 2: outputUserDetails(appUser) // Pass entire struct
}

// Function to display user details (currently accepts individual strings)
func outputUserDetails(firstName, lastName, birthdate string) {
	// Parameters: three separate strings
	// firstName, lastName string - shorthand notation (both are string type)
	
	// ...
	fmt.Println(firstName, lastName, birthdate)
	// Prints the three values separated by spaces
}

// Reusable function to get user input with a prompt
func getUserData(promptText string) string {
	// Parameter: prompt message to display
	// Returns: user's input as string
	
	fmt.Print(promptText)
	// Displays the prompt (cursor stays on same line)
	
	var value string
	// Variable to store user input
	
	fmt.Scan(&value)
	// Reads input from console
	// &value passes memory address so Scan can modify it
	// WARNING: Stops at first whitespace!
	
	return value
	// Returns the captured input
}
