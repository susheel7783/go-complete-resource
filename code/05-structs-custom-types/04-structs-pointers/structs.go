package main

import (
	"fmt"   // Package for formatted I/O (printing and scanning)
	"time"  // Package for working with dates and times
)

// Define a custom type (struct) to represent a user
// A struct is like a blueprint that groups related data together
type user struct {
	firstName string    // User's first name
	lastName  string    // User's last name
	birthDate string    // User's birth date as a string
	createdAt time.Time // Timestamp when user was created (uses time.Time type)
}

func main() {
	// Collect user input by calling getUserData function for each field
	// Each call displays a prompt and waits for user to type input
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Declare a variable 'appUser' of type 'user' (our custom struct)
	var appUser user
	
	// Initialize the struct by assigning values to all fields
	// This creates a complete user object with all the collected data
	appUser = user{
		firstName: userFirstName,  // Assign collected first name
		lastName:  userLastName,   // Assign collected last name
		birthDate: userBirthdate,  // Assign collected birthdate
		createdAt: time.Now(),     // Automatically set current date/time
	}

	// Process the user data (placeholder comment for future functionality)
	// ... do something awesome with that gathered data!
	
	// Pass the address of appUser (pointer) to outputUserDetails function
	// Using & gives us a pointer, which is more efficient than copying the entire struct
	outputUserDetails(&appUser)
}

// Function that takes a pointer to a user struct and displays their details
// Using a pointer (*user) means we're passing a reference, not a copy
// This is more efficient for larger structs
func outputUserDetails(u *user) {
	// Print the user's information to the console
	// u.firstName accesses the firstName field through the pointer
	fmt.Println(u.firstName, u.lastName, u.birthDate)
	// Note: createdAt is not printed here, but could be added
}

// Reusable function to get user input with a custom prompt
// Takes a string parameter (the prompt to display)
// Returns the string value entered by the user
func getUserData(promptText string) string {
	fmt.Print(promptText)  // Display the prompt (without newline)
	
	var value string       // Declare variable to store user input
	
	fmt.Scan(&value)       // Read user input from console
	                       // & gives the memory address where input should be stored
	
	return value           // Return the collected input to the caller
}
```

**Key Concepts Explained:**

1. **Structs**: Custom data types that group related fields together (like a class in other languages)

2. **Pointers (`*` and `&`)**: 
   - `&appUser` creates a pointer (memory address) to appUser
   - `*user` means the function accepts a pointer to a user struct
   - This avoids copying the entire struct, making it more efficient

3. **time.Now()**: Gets the current date and time automatically

4. **Function Organization**: Code is split into reusable functions for better organization

**Sample Output When Running:**
```
Please enter your first name: John
Please enter your last name: Doe
Please enter your birthdate (MM/DD/YYYY): 01/15/1990
John Doe 01/15/1990
