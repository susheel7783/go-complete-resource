package main

import (
	"fmt" // Imports fmt package for formatted I/O operations (printing, scanning)
)

func main() { // The main function - program execution starts here
	
	// Gets user's first name
	firstName := getUserData("Please enter your first name: ")
	// Calls getUserData function with prompt text
	// getUserData returns a string (the user's input)
	// firstName stores the returned value
	// := shorthand declaration (type string inferred from return value)
	
	// Gets user's last name
	lastName := getUserData("Please enter your last name: ")
	// Reuses the same function with different prompt text
	// lastName stores the second input
	
	// Gets user's birthdate
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	// Same function, different prompt
	// birthdate stores the third input
	// Note: Input is stored as string, not validated as actual date format
	
	// ... do something awesome with that gathered data!
	// Comment indicating future functionality could be added here
	// Could validate data, save to file, create user profile, etc.
	
	// Prints all collected data
	fmt.Println(firstName, lastName, birthdate)
	// Println automatically adds spaces between values
	// Example output: "John Doe 05/15/1990"
}

// Reusable function to get user input with a custom prompt
func getUserData(promptText string) string {
	// Parameter: promptText - the message to display to the user
	// Returns: string - the user's input
	
	fmt.Print(promptText)
	// Displays the prompt text
	// Print (not Println): cursor stays on same line for user to type
	// Example: "Please enter your first name: _" (cursor here)
	
	var value string
	// Declares a variable to store user input
	// Type: string (can hold text of any length)
	// Initialized to empty string "" by default
	
	fmt.Scan(&value)
	// Reads user input from console
	// &value: passes memory address so Scan can modify the variable
	// Waits for user to type and press Enter
	// Stops reading at first whitespace (space, tab, newline)
	// IMPORTANT: Only reads FIRST WORD if user enters multiple words
	
	return value
	// Returns the captured user input back to the caller
	// This value will be assigned to firstName, lastName, or birthdate
}
