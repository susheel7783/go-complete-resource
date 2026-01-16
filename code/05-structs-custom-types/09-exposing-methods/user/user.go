// ==================== PACKAGE DECLARATION ====================
// Package name MUST match the directory name
// This file should be in a directory named "user"
// 
// Directory structure:
// project/
// ├── go.mod
// ├── main.go
// └── user/
//     └── user.go (this file)
package user

import (
	"errors" // Package for creating error values
	"fmt"    // Package for formatted I/O (printing)
	"time"   // Package for working with dates and times
)

// ==================== STRUCT DEFINITION ====================
// User struct represents a user in the system
// 
// IMPORTANT CAPITALIZATION:
// - "User" (uppercase U) = EXPORTED (public, visible outside this package)
// - "firstName", "lastName", etc. (lowercase) = UNEXPORTED (private, only visible within user package)
// 
// This creates ENCAPSULATION:
// - Other packages can create and use User objects
// - But cannot directly access or modify the internal fields
// - Must use the exported methods to interact with User data
type User struct {
	firstName string    // Private field - cannot be accessed from other packages
	lastName  string    // Private field - cannot be accessed from other packages
	birthDate string    // Private field - cannot be accessed from other packages
	createdAt time.Time // Private field - timestamp when user was created
}

// ==================== EXPORTED METHOD - OUTPUT ====================
// OutputUserDetails displays the user's information
// 
// Method signature breakdown:
// - (u *User) = receiver (pointer to User struct)
// - OutputUserDetails = method name (capitalized = EXPORTED)
// - () = no parameters
// 
// Because it's EXPORTED (capitalized), it can be called from other packages:
// appUser.OutputUserDetails() ✅
func (u *User) OutputUserDetails() {
	// Access private fields directly (we're inside the same package)
	// Print first name, last name, and birth date
	fmt.Println(u.firstName, u.lastName, u.birthDate)
	// Note: createdAt is not displayed here
}

// ==================== EXPORTED METHOD - MUTATION ====================
// ClearUserName clears the user's name fields
// 
// Pointer receiver (*User) is REQUIRED because:
// - This method MODIFIES the struct fields
// - Without pointer, it would only modify a copy
// 
// Because it's EXPORTED (capitalized), it can be called from other packages:
// appUser.ClearUserName() ✅
func (u *User) ClearUserName() {
	u.firstName = "" // Set first name to empty string
	u.lastName = ""  // Set last name to empty string
	// birthDate and createdAt remain unchanged
}

// ==================== EXPORTED CONSTRUCTOR FUNCTION ====================
// New creates and returns a new User instance with validation
// 
// Naming convention:
// - In Go, constructor functions in their own package are typically named "New"
// - Called as: user.New(...) when imported
// - If in a shared package, would be named "NewUser"
// 
// Parameters:
// - firstName: user's first name
// - lastName: user's last name
// - birthdate: user's birth date as string
// 
// Returns:
// - *User: pointer to the newly created User (or nil if validation fails)
// - error: nil if successful, error object if validation fails
// 
// This is Go's standard error handling pattern - return both result and error
func New(firstName, lastName, birthdate string) (*User, error) {
	// ==================== VALIDATION ====================
	// Check if any required field is empty
	// Using || (OR operator) - if ANY field is empty, validation fails
	if firstName == "" || lastName == "" || birthdate == "" {
		// Return nil for User pointer (no object created)
		// Return error with descriptive message
		return nil, errors.New("First name, last name and birthdate are required.")
	}
	
	// ==================== OBJECT CREATION ====================
	// If validation passes, create and return the User
	// 
	// The & operator returns the memory address (pointer) of the struct
	// This is efficient - we return a pointer, not a copy of the entire struct
	return &User{
		firstName: firstName,   // Assign parameter to private field
		lastName:  lastName,    // Assign parameter to private field
		birthDate: birthdate,   // Assign parameter to private field
		createdAt: time.Now(),  // Automatically set to current timestamp
		                        // Users don't provide this - it's set internally
	}, nil // Return nil for error (no error occurred)
}

// ==================== WHAT'S NOT INCLUDED (But Could Be) ====================
//
// Additional methods you might add:
//
// // Getter methods (Go convention: no "Get" prefix)
// func (u *User) FirstName() string {
//     return u.firstName
// }
//
// func (u *User) LastName() string {
//     return u.lastName
// }
//
// func (u *User) BirthDate() string {
//     return u.birthDate
// }
//
// func (u *User) CreatedAt() time.Time {
//     return u.createdAt
// }
//
// // Setter methods
// func (u *User) SetFirstName(name string) error {
//     if name == "" {
//         return errors.New("first name cannot be empty")
//     }
//     u.firstName = name
//     return nil
// }
//
// // Business logic methods
// func (u *User) Age() (int, error) {
//     // Calculate age from birthDate
// }
//
// func (u *User) FullName() string {
//     return u.firstName + " " + u.lastName
// }


// -------------
🔑 KEY CONCEPTS:
1. Package-Level Encapsulation
go// ✅ CAN DO (from main.go):
import "example.com/structs/user"

u, _ := user.New("John", "Doe", "01/01/1990")  // Create user
u.OutputUserDetails()                          // Call exported method
u.ClearUserName()                              // Call exported method

// ❌ CANNOT DO (from main.go):
u.firstName = "Jane"        // Error: firstName is unexported
fmt.Println(u.birthDate)    // Error: birthDate is unexported
2. Exported vs Unexported Summary
ItemExported?Accessible from main?User✅ Yes (capital U)✅ YesfirstName❌ No (lowercase f)❌ NoNew()✅ Yes (capital N)✅ YesOutputUserDetails()✅ Yes (capital O)✅ YesClearUserName()✅ Yes (capital C)✅ Yes

3. Why This Design?
Encapsulation Benefits:

✅ Data Protection: Fields can't be modified incorrectly from outside
✅ Validation Control: All changes go through methods that can validate
✅ Flexibility: Can change internal implementation without breaking external code
✅ API Clarity: Only exported items are the public interface

Example of Protection:
go// If fields were exported:
user.firstName = ""  // ❌ Could break invariants
user.createdAt = someOldTime  // ❌ Could corrupt data

// With private fields:
// Must use methods which can enforce rules
user.ClearUserName()  // ✅ Controlled way to clear name

4. Constructor Pattern Benefits
gofunc New(...) (*User, error) {
    // ✅ Validation in one place
    // ✅ Ensures all Users are created properly
    // ✅ Auto-sets createdAt (users can't forget or set wrong)
    // ✅ Returns pointer for efficiency
    // ✅ Returns error for proper error handling
}

🎯 Real-World Usage from main.go:
gopackage main

import (
    "fmt"
    "example.com/structs/user"
)

func main() {
    // Create user with validation
    u, err := user.New("Alice", "Smith", "05/15/1990")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    // Use exported methods
    u.OutputUserDetails()  // ✅ Works - exported
    u.ClearUserName()      // ✅ Works - exported
    
    // Cannot access private fields
    // fmt.Println(u.firstName)  // ❌ Compile error
    // u.birthDate = "wrong"     // ❌ Compile error
}

