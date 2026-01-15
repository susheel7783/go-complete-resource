package main

import (
	"errors" // Package for creating error values
	"fmt"    // Package for formatted I/O (printing and scanning)
	"time"   // Package for working with dates and times
)

// Define a custom type (struct) to represent a user
type user struct {
	firstName string    // User's first name
	lastName  string    // User's last name
	birthDate string    // User's birth date as a string
	createdAt time.Time // Timestamp when user was created
}

// Method with pointer receiver to display user details
func (u *user) outputUserDetails() {
	// Print user's basic information to console
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// Method with pointer receiver to clear name fields
// Pointer receiver is required to modify the original struct
func (u *user) clearUserName() {
	u.firstName = "" // Clear first name
	u.lastName = ""  // Clear last name
}

// ==================== CONSTRUCTOR WITH ERROR HANDLING ====================
// Constructor function that creates a new user with VALIDATION
// Returns TWO values: (*user, error) - this is Go's error handling pattern
// 
// Return values:
// - *user: pointer to the created user (nil if validation fails)
// - error: nil if successful, error message if validation fails
func newUser(firstName, lastName, birthdate string) (*user, error) {
	// ==================== VALIDATION LOGIC ====================
	// Check if any required field is empty
	// The || operator means "OR" - if ANY field is empty, validation fails
	if firstName == "" || lastName == "" || birthdate == "" {
		// Return nil for the user pointer (no user created)
		// Return an error with a descriptive message
		// errors.New() creates a new error value with the given message
		return nil, errors.New("First name, last name and birthdate are required.")
	}
	
	// If validation passes, create and return the user
	// Return the pointer to the new user AND nil for error (no error occurred)
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(), // Auto-set creation timestamp
	}, nil // nil means "no error"
}

func main() {
	// Collect user input for each field
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ==================== CALLING FUNCTION WITH ERROR HANDLING ====================
	// Declare appUser as a pointer to user
	var appUser *user
	
	// Call newUser and capture BOTH return values
	// appUser: will hold the user pointer (or nil if error)
	// err: will hold the error (or nil if successful)
	appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	
	// ==================== ERROR CHECKING ====================
	// Check if an error occurred during user creation
	// In Go, errors are values that need to be explicitly checked
	if err != nil {
		// If error exists (not nil), print the error message
		fmt.Println(err) // Outputs: "First name, last name and birthdate are required."
		
		// Exit the program early - don't continue with invalid data
		// return exits the main function, ending the program
		return
	}
	
	// If we reach here, user creation was successful and appUser is valid
	// ... do something awesome with that gathered data!
	
	// Display user details (will show the entered information)
	appUser.outputUserDetails()
	
	// Clear the user's name fields
	appUser.clearUserName()
	
	// Display user details again (names will be empty, birthdate remains)
	appUser.outputUserDetails()
}

// Reusable helper function to get user input with a custom prompt
func getUserData(promptText string) string {
	fmt.Print(promptText) // Display the prompt (no newline)
	var value string      // Variable to store user input
	
	// fmt.Scanln reads a line of input (including spaces) until newline
	// This is different from fmt.Scan which stops at whitespace
	// Scanln is better for names like "John Doe" (with spaces)
	fmt.Scanln(&value)
	
	return value // Return the collected input
}

// ----------------------

🔑 KEY CONCEPTS:
1. Go's Error Handling Pattern
Go doesn't use try-catch exceptions. Instead, functions return errors as values:
go// Function signature with error return
func newUser(...) (*user, error) {
    // Returns two values: result AND error
}

// Calling and checking
result, err := newUser(...)
if err != nil {
    // Handle error
}
// Use result

2. Multiple Return Values
gofunc newUser(...) (*user, error) {
    // Success case: return user pointer and nil error
    return &user{...}, nil
    
    // Error case: return nil pointer and error
    return nil, errors.New("error message")
}

3. Error Checking Flow
goappUser, err := newUser(firstName, lastName, birthdate)

if err != nil {           // If error exists
    fmt.Println(err)      // Display error message
    return                // Exit early (don't continue)
}

// Safe to use appUser here - we know it's valid
appUser.outputUserDetails()

4. Creating Errors
Three common ways:
go// 1. errors.New() - simple error message
return nil, errors.New("First name is required")

// 2. fmt.Errorf() - formatted error message
return nil, fmt.Errorf("invalid age: %d", age)

// 3. Custom error types (advanced)
type ValidationError struct { ... }
```

---

**📊 Sample Outputs:**

**Scenario 1: Valid Input**
```
Please enter your first name: Bob
Please enter your last name: Smith
Please enter your birthdate (MM/DD/YYYY): 12/25/1985
Bob Smith 12/25/1985
 12/25/1985
```

**Scenario 2: Empty Input (Error)**
```
Please enter your first name: 
Please enter your last name: Smith
Please enter your birthdate (MM/DD/YYYY): 12/25/1985
First name, last name and birthdate are required.

🎯 Important Changes from Previous Version:
1. fmt.Scanln() vs fmt.Scan()
go// OLD: fmt.Scan() - stops at whitespace
fmt.Scan(&value)  // "John Doe" → only captures "John"

// NEW: fmt.Scanln() - reads entire line
fmt.Scanln(&value)  // "John Doe" → captures "John Doe"
2. Constructor Returns Two Values
go// OLD: Returns only user pointer
func newUser(...) *user { ... }

// NEW: Returns user pointer AND error
func newUser(...) (*user, error) { ... }
3. Validation Before Creation
go// Check all required fields
if firstName == "" || lastName == "" || birthdate == "" {
    return nil, errors.New("...")  // Fail fast
}

// Only create user if validation passes
return &user{...}, nil
