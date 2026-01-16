package main

import (
	"fmt" // Standard library for formatted I/O
	
	// Import the custom user package
	// This package now contains both User and Admin types
	"example.com/structs/user"
)

func main() {
	// ==================== CREATING A REGULAR USER ====================
	// Collect user input for creating a regular User
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Declare appUser as a pointer to user.User type
	var appUser *user.User
	
	// Create a new User with validation
	// Returns (*user.User, error)
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	
	// Check if user creation failed (validation error)
	if err != nil {
		fmt.Println(err) // Display error message
		return           // Exit program early
	}

	// ==================== CREATING AN ADMIN USER ====================
	// Create an Admin user using the NewAdmin constructor
	// 
	// NewAdmin is a separate constructor function that creates an Admin type
	// Admin likely embeds/contains User and adds additional fields (email, password)
	// 
	// Parameters:
	// - "test@example.com" = admin email
	// - "test123" = admin password
	// 
	// This returns a *user.Admin (pointer to Admin struct)
	admin := user.NewAdmin("test@example.com", "test123")
	
	// ==================== USING ADMIN METHODS ====================
	// Call OutputUserDetails on the admin object
	// 
	// This works because Admin likely:
	// - Embeds User struct (has all User fields and methods)
	// - OR implements the same OutputUserDetails method
	admin.OutputUserDetails()
	
	// Clear the admin's name fields
	// Admin has access to ClearUserName because it embeds/extends User
	admin.ClearUserName()
	
	// Display admin details again (name fields now cleared)
	admin.OutputUserDetails()

	// ... do something awesome with that gathered data!
	
	// ==================== USING REGULAR USER METHODS ====================
	// Call methods on the regular User object
	appUser.OutputUserDetails() // Display: firstName lastName birthDate
	appUser.ClearUserName()     // Clear name fields
	appUser.OutputUserDetails() // Display: (empty) (empty) birthDate
}

// ==================== HELPER FUNCTION ====================
// Reusable function to prompt user for input
// Remains in main package (unexported, only used here)
func getUserData(promptText string) string {
	fmt.Print(promptText) // Display the prompt without newline
	var value string      // Variable to store user input
	fmt.Scanln(&value)    // Read entire line of input (including spaces)
	return value          // Return the collected value
}

// ----------------
📁 THE ADMIN IMPLEMENTATION (user/admin.go)
Here's what the user/admin.go file would likely contain:
gopackage user

import (
	"fmt"
	"time"
)

// ==================== ADMIN STRUCT WITH EMBEDDING ====================
// Admin struct embeds User and adds admin-specific fields
// 
// Embedding (also called composition) in Go:
// - Admin "has-a" User (not "is-a" - Go doesn't have inheritance)
// - Admin automatically gets all User fields and methods
// - Can override methods if needed
type Admin struct {
	User                // Embedded User struct (anonymous field)
	email    string     // Admin-specific field: email address
	password string     // Admin-specific field: password
}

// ==================== ADMIN CONSTRUCTOR ====================
// NewAdmin creates a new Admin with hardcoded user info
// 
// In a real application, you might:
// - Accept firstName, lastName, birthdate as parameters
// - Hash the password before storing
// - Validate email format
// - Store admin flag in database
func NewAdmin(email, password string) *Admin {
	return &Admin{
		User: User{
			firstName: "ADMIN",           // Hardcoded admin first name
			lastName:  "USER",            // Hardcoded admin last name
			birthDate: "01/01/2000",      // Hardcoded birthdate
			createdAt: time.Now(),        // Current timestamp
		},
		email:    email,    // Admin's email
		password: password, // Admin's password (should be hashed in production!)
	}
}

// ==================== OPTIONAL: OVERRIDE METHOD ====================
// You can override embedded methods to customize behavior
// 
// If you want Admin to display differently:
// func (a *Admin) OutputUserDetails() {
//     fmt.Printf("ADMIN: %s %s (Email: %s)\n", 
//         a.firstName, a.lastName, a.email)
// }

// ==================== ADMIN-SPECIFIC METHODS ====================
// Methods that only Admin has (not User)

// Email getter method
func (a *Admin) Email() string {
	return a.email
}

// Method to check if password matches
func (a *Admin) CheckPassword(password string) bool {
	return a.password == password
	// In production: use bcrypt.CompareHashAndPassword()
}

🔑 KEY CONCEPTS:
1. Struct Embedding (Composition)
gotype Admin struct {
    User              // Embedded struct - Admin has all User fields/methods
    email    string   // Additional Admin-specific field
    password string   // Additional Admin-specific field
}

// Admin automatically inherits:
// - firstName, lastName, birthDate, createdAt (fields)
// - OutputUserDetails(), ClearUserName() (methods)
2. Method Inheritance Through Embedding
goadmin := user.NewAdmin("test@example.com", "test123")

// These work because Admin embeds User:
admin.OutputUserDetails()  // ✅ Inherited from User
admin.ClearUserName()      // ✅ Inherited from User

// Admin-specific methods:
admin.Email()              // ✅ Only Admin has this
admin.CheckPassword("...")  // ✅ Only Admin has this
3. Accessing Embedded Fields
go// Inside Admin methods, you can access User fields:
func (a *Admin) SomeMethod() {
    fmt.Println(a.firstName)  // Access embedded User field
    fmt.Println(a.email)      // Access Admin's own field
    
    // Or explicitly:
    fmt.Println(a.User.firstName)
}
```

---

**📊 Sample Output:**
```
Please enter your first name: John
Please enter your last name: Doe
Please enter your birthdate (MM/DD/YYYY): 05/15/1990
ADMIN USER 01/01/2000
 01/01/2000
John Doe 05/15/1990
 05/15/1990
```

**Explanation:**
1. Admin is created with hardcoded values: "ADMIN USER 01/01/2000"
2. After clearing: " 01/01/2000" (names cleared, birthdate remains)
3. Regular user displays entered values: "John Doe 05/15/1990"
4. After clearing: " 05/15/1990" (names cleared)

---

**🎯 Complete Project Structure:**
```
project/
├── go.mod              # module example.com/structs
├── main.go             # Entry point (creates User and Admin)
└── user/
    ├── user.go         # User struct and methods
    └── admin.go        # Admin struct (embeds User)

🚀 Composition vs Inheritance:
Go doesn't have traditional inheritance. Instead it uses composition:
go// Traditional inheritance (other languages):
// class Admin extends User { ... }

// Go composition (embedding):
type Admin struct {
    User  // Has-a relationship, not is-a
    // ... additional fields
}
Benefits:

✅ More flexible than inheritance
✅ Can embed multiple structs
✅ Clear which methods come from which type
✅ Avoids complex inheritance hierarchies


💡 Production Improvements:
go// Better Admin constructor with validation:
func NewAdmin(email, password, firstName, lastName, birthdate string) (*Admin, error) {
    // Validate email format
    if !isValidEmail(email) {
        return nil, errors.New("invalid email format")
    }
    
    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    
    // Create User first (with validation)
    u, err := New(firstName, lastName, birthdate)
    if err != nil {
        return nil, err
    }
    
    return &Admin{
        User:     *u,
        email:    email,
        password: string(hashedPassword),
    }, nil
}
