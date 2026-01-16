package main

import (
	"fmt" // Standard library package for formatted I/O
	
	// ==================== CUSTOM PACKAGE IMPORT ====================
	// Import the custom "user" package from the local project
	// "example.com/structs/user" is the module path defined in go.mod
	// 
	// Path breakdown:
	// - "example.com/structs" is the module name (from go.mod)
	// - "user" is a subdirectory/package within the project
	// 
	// Project structure would look like:
	// project/
	// ├── go.mod (contains: module example.com/structs)
	// ├── main.go (this file)
	// └── user/
	//     └── user.go (contains User struct and methods)
	"example.com/structs/user"
)

func main() {
	// Collect user input for each field
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ==================== USING EXPORTED TYPES ====================
	// Declare appUser as a pointer to user.User
	// 
	// Note the capitalization:
	// - "user" (lowercase) = package name
	// - "User" (uppercase) = exported type from that package
	// 
	// In Go, capitalization determines visibility:
	// - Uppercase = Exported (public, accessible from other packages)
	// - Lowercase = Unexported (private, only accessible within same package)
	var appUser *user.User
	
	// ==================== CALLING EXPORTED FUNCTION ====================
	// Call user.New() - the constructor function from the user package
	// 
	// "New" is capitalized = exported (public)
	// If it were "new" (lowercase), it would be unexported (private)
	// and we couldn't call it from main package
	//
	// This function returns (*user.User, error) just like before
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	
	// ==================== ERROR HANDLING ====================
	// Check if user creation failed
	if err != nil {
		fmt.Println(err) // Print error message
		return           // Exit program early
	}

	// ... do something awesome with that gathered data!
	
	// ==================== CALLING EXPORTED METHODS ====================
	// Call exported methods on the user object
	// 
	// Note the capitalization of method names:
	// - OutputUserDetails (capitalized) = exported method
	// - ClearUserName (capitalized) = exported method
	// 
	// These were previously:
	// - outputUserDetails (lowercase) = unexported (private)
	// - clearUserName (lowercase) = unexported (private)
	appUser.OutputUserDetails() // Display user information
	appUser.ClearUserName()     // Clear name fields
	appUser.OutputUserDetails() // Display again (names cleared)
}

// Helper function to get user input
// This remains in the main package and is unexported (lowercase)
// because it's only used within main.go
func getUserData(promptText string) string {
	fmt.Print(promptText) // Display prompt
	var value string      // Variable to store input
	fmt.Scanln(&value)    // Read entire line of input
	return value          // Return collected input
}

// ------------------------
📁 THE USER PACKAGE (user/user.go)
Here's what the separate user/user.go file would contain:
gopackage user // Package declaration - must match directory name

import (
	"errors" // For creating error values
	"fmt"    // For printing
	"time"   // For timestamps
)

// ==================== EXPORTED STRUCT ====================
// User struct with CAPITALIZED name = exported (public)
// Fields are LOWERCASE = unexported (private to this package)
// 
// This is Go's encapsulation:
// - Other packages can use User type
// - But cannot directly access firstName, lastName, etc.
// - Must use exported methods to interact with the data
type User struct {
	firstName string    // unexported field
	lastName  string    // unexported field
	birthDate string    // unexported field
	createdAt time.Time // unexported field
}

// ==================== EXPORTED METHOD ====================
// OutputUserDetails is capitalized = exported (public)
// Can be called from other packages
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// ==================== EXPORTED METHOD ====================
// ClearUserName is capitalized = exported (public)
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// ==================== EXPORTED CONSTRUCTOR ====================
// New is capitalized = exported (public)
// Convention: constructor functions are often named just "New"
// when in their own package (user.New instead of user.NewUser)
func New(firstName, lastName, birthdate string) (*User, error) {
	// Validation
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required.")
	}
	
	// Return pointer to User and nil error
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
```

---

**🔑 KEY CONCEPTS:**

### **1. Package Organization**
```
project/
├── go.mod                    # Module definition
├── main.go                   # Entry point (package main)
└── user/                     # Custom package
    └── user.go               # User type and methods
2. Exported vs Unexported
NameVisibilityAccessUserExportedAccessible everywhereuserUnexportedOnly within same packageNew()ExportedCallable from other packagesnew()UnexportedOnly within same packageOutputUserDetails()ExportedPublic methodoutputUserDetails()UnexportedPrivate method
3. Encapsulation
go// In user package:
type User struct {
    firstName string  // PRIVATE - cannot access from main
}

// From main package:
appUser.firstName = "Bob"  // ❌ ERROR - cannot access unexported field

// Must use public methods:
appUser.OutputUserDetails()  // ✅ OK - exported method

4. Import Paths
go// go.mod file contains:
module example.com/structs

// So you import:
import "example.com/structs/user"

// And use:
user.New(...)
user.User
