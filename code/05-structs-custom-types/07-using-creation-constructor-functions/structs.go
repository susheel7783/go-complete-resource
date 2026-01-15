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

// Method with pointer receiver to display user details
// Pointer receiver (*user) allows working with the original struct
func (u *user) outputUserDetails() {
	// Print user's basic information
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// Method with pointer receiver to clear name fields
// MUST use pointer receiver to modify the original struct
func (u *user) clearUserName() {
	u.firstName = ""  // Set first name to empty string
	u.lastName = ""   // Set last name to empty string
}

// ==================== CONSTRUCTOR FUNCTION ====================
// This is a "constructor" or "factory" function (Go doesn't have built-in constructors)
// Convention: name it "new" + TypeName (newUser, newCar, newProduct, etc.)
// 
// Takes individual parameters for each field (except createdAt which is auto-set)
// Returns a POINTER to a newly created user struct (*user)
func newUser(firstName, lastName, birthdate string) *user {
	// The & operator returns the ADDRESS (pointer) of the struct
	// This creates the struct and immediately returns a pointer to it
	return &user{
		firstName: firstName,   // Assign parameter to struct field
		lastName:  lastName,    // Assign parameter to struct field
		birthDate: birthdate,   // Assign parameter to struct field
		createdAt: time.Now(),  // Automatically set to current time
	}
	
	// Note: createdAt is handled internally - users don't need to provide it
	// This ensures consistency - all users get a proper timestamp
}

func main() {
	// Collect user input for each field
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ==================== USING THE CONSTRUCTOR ====================
	// Declare appUser as a POINTER to user (notice the asterisk)
	var appUser *user
	
	// Call the constructor function to create and initialize a new user
	// newUser returns a pointer, so appUser stores the memory address of the struct
	// This is cleaner than manually creating and initializing the struct
	appUser = newUser(userFirstName, userLastName, userBirthdate)
	
	// ALTERNATIVE SYNTAX (shorter, more common):
	// appUser := newUser(userFirstName, userLastName, userBirthdate)

	// ... do something awesome with that gathered data!
	
	// Call methods on the user pointer (Go automatically handles dereferencing)
	appUser.outputUserDetails()  // Display: firstName lastName birthDate
	appUser.clearUserName()      // Clear the name fields
	appUser.outputUserDetails()  // Display: (empty) (empty) birthDate
}

// Reusable helper function to get user input with a custom prompt
func getUserData(promptText string) string {
	fmt.Print(promptText)  // Display the prompt
	var value string       // Variable to store user input
	fmt.Scan(&value)       // Read input from console
	return value           // Return the collected value
}


// ------------------------------------------------------------------

🔑 KEY CONCEPTS:
1. Constructor Function Pattern
Before (Manual Creation):
govar appUser user
appUser = user{
	firstName: userFirstName,
	lastName:  userLastName,
	birthDate: userBirthdate,
	createdAt: time.Now(),
}
After (Using Constructor):
govar appUser *user
appUser = newUser(userFirstName, userLastName, userBirthdate)
// Or shorter:
// appUser := newUser(userFirstName, userLastName, userBirthdate)

2. Benefits of Constructor Functions
✅ Encapsulation: Hides implementation details (users don't set createdAt manually)
✅ Consistency: Ensures all users are created with proper initialization
✅ Validation: Can add validation logic in one place (future enhancement)
✅ Cleaner Code: One line instead of multi-line struct literal
✅ Default Values: Can set fields automatically (like createdAt)

3. Why Return a Pointer?
gofunc newUser(...) *user {  // Returns POINTER
	return &user{...}      // & creates pointer to the struct
}
Reasons:

Efficiency: Avoid copying large structs
Consistency: Methods use pointer receivers, so return pointer
Mutability: Allows methods to modify the struct
Go Convention: Constructor functions typically return pointers


4. Memory Management
goreturn &user{...}  // Creates struct on the heap, returns its address
```

- Go's garbage collector automatically manages memory
- The struct lives as long as something references it
- When `appUser` is no longer used, Go cleans up the memory
- No manual `free()` or `delete` needed!

---

**📊 Sample Output:**
```
Please enter your first name: Alice
Please enter your last name: Johnson
Please enter your birthdate (MM/DD/YYYY): 07/10/1992
Alice Johnson 07/10/1992
 07/10/1992

🎯 Real-World Example with Validation:
gofunc newUser(firstName, lastName, birthdate string) (*user, error) {
	// Add validation
	if firstName == "" || lastName == "" {
		return nil, fmt.Errorf("name fields cannot be empty")
	}
	
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
