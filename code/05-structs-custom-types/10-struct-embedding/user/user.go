package user

import (
	"errors" // Package for creating error values
	"fmt"    // Package for formatted I/O (printing)
	"time"   // Package for working with dates and times
)

// ==================== BASE USER STRUCT ====================
// User struct represents a basic user in the system
// 
// All fields are UNEXPORTED (lowercase):
// - Cannot be accessed directly from other packages
// - Must use exported methods to interact with data
// - Provides encapsulation and data protection
type User struct {
	firstName string    // Private: user's first name
	lastName  string    // Private: user's last name
	birthDate string    // Private: user's birth date
	createdAt time.Time // Private: timestamp when user was created
}

// ==================== ADMIN STRUCT WITH EMBEDDING ====================
// Admin struct represents an administrator user
// 
// Field order matters for struct literals:
// - email and password are Admin-specific fields
// - User is an EMBEDDED (anonymous) field
// 
// Embedding means:
// - Admin "has-a" User (composition, not inheritance)
// - Admin automatically gets ALL User fields (firstName, lastName, etc.)
// - Admin automatically gets ALL User methods (OutputUserDetails, ClearUserName)
// - Can access User fields like: admin.firstName (Go promotes them up)
type Admin struct {
	email    string // Private: admin's email address
	password string // Private: admin's password (should be hashed in production!)
	User           // Embedded User struct - Admin inherits User's fields and methods
	               // This is an ANONYMOUS field (no name, just the type)
}

// ==================== USER METHOD - OUTPUT ====================
// OutputUserDetails displays the user's information
// 
// Receiver: (u *User) - pointer to User
// - Works for both User and Admin (Admin embeds User)
// - When called on Admin, 'u' refers to the embedded User part
// 
// Method is EXPORTED (capitalized) so it can be called from other packages
func (u *User) OutputUserDetails() {
	// Access private fields directly (we're inside the user package)
	// Print: firstName lastName birthDate
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// ==================== USER METHOD - MUTATION ====================
// ClearUserName clears the user's name fields
// 
// Pointer receiver (*User) is REQUIRED because:
// - This method MODIFIES the struct
// - Without pointer, it would only modify a copy
// 
// Works for both User and Admin through embedding
func (u *User) ClearUserName() {
	u.firstName = "" // Set first name to empty string
	u.lastName = ""  // Set last name to empty string
	// Note: birthDate and createdAt remain unchanged
}

// ==================== ADMIN CONSTRUCTOR ====================
// NewAdmin creates and returns a new Admin instance
// 
// IMPORTANT DIFFERENCE from New():
// - Returns Admin (value, NOT pointer) - notice no asterisk *
// - No error return - no validation performed
// - Hardcodes the User fields to generic admin values
// 
// Parameters:
// - email: administrator's email address
// - password: administrator's password (plaintext - should be hashed!)
// 
// Returns:
// - Admin value (not a pointer like User's New function)
func NewAdmin(email, password string) Admin {
	// Return an Admin struct initialized with:
	// 1. Admin-specific fields (email, password)
	// 2. Embedded User with hardcoded values
	return Admin{
		email:    email,    // Set admin's email
		password: password, // Set admin's password
		
		// Initialize the embedded User struct
		// All admins get the same generic User values
		User: User{
			firstName: "ADMIN",      // Hardcoded first name
			lastName:  "ADMIN",      // Hardcoded last name
			birthDate: "---",        // Placeholder birthdate
			createdAt: time.Now(),   // Current timestamp
		},
	}
}

// ==================== USER CONSTRUCTOR ====================
// New creates and returns a new User instance with validation
// 
// This is the constructor for regular users (not admins)
// 
// Parameters:
// - firstName: user's first name
// - lastName: user's last name
// - birthdate: user's birth date
// 
// Returns:
// - *User: pointer to the newly created User (nil if validation fails)
// - error: nil if successful, error object with message if validation fails
// 
// IMPORTANT: Returns POINTER (*User) unlike NewAdmin which returns value
func New(firstName, lastName, birthdate string) (*User, error) {
	// ==================== VALIDATION ====================
	// Validate that all required fields are provided
	// If ANY field is empty, creation fails
	if firstName == "" || lastName == "" || birthdate == "" {
		// Return nil pointer and error
		return nil, errors.New("First name, last name and birthdate are required.")
	}
	
	// ==================== USER CREATION ====================
	// If validation passes, create and return the User
	// & operator returns a pointer to the newly created struct
	return &User{
		firstName: firstName,   // Set from parameter
		lastName:  lastName,    // Set from parameter
		birthDate: birthdate,   // Set from parameter
		createdAt: time.Now(),  // Automatically set to current time
	}, nil // No error occurred
}

// ==================== HOW ADMIN USES USER METHODS ====================
// 
// When you call methods on Admin:
// 
// admin := NewAdmin("test@example.com", "test123")
// admin.OutputUserDetails()  // Calls User's method on the embedded User
// admin.ClearUserName()      // Calls User's method on the embedded User
// 
// Go automatically "promotes" the embedded User's methods to Admin
// It's as if Admin has these methods, but they operate on the embedded User part
// 
// Behind the scenes, Go does:
// admin.User.OutputUserDetails()  // Explicit access to embedded struct
// 
// But you can just write:
// admin.OutputUserDetails()       // Go promotes it automatically


// ---------------
🔑 KEY CONCEPTS:
1. Struct Embedding (Anonymous Fields)
gotype Admin struct {
    email    string
    password string
    User           // Anonymous field - just the type name, no field name
}

// This is different from:
type Admin struct {
    email    string
    password string
    user User      // Named field - would need admin.user.firstName
}
2. Method Promotion
When a struct embeds another struct, the embedded struct's methods are promoted:
goadmin := NewAdmin("test@example.com", "test123")

// All of these work because User is embedded:
admin.OutputUserDetails()  // ✅ Promoted from User
admin.ClearUserName()      // ✅ Promoted from User

// Can also access explicitly:
admin.User.OutputUserDetails()  // ✅ Explicit access
3. Field Access with Embedding
go// Inside the user package, Admin can access User's private fields:
admin.firstName  // ✅ Promoted from embedded User
admin.lastName   // ✅ Promoted from embedded User
admin.email      // ✅ Admin's own field

// Explicit access also works:
admin.User.firstName  // ✅ Explicit path

🎯 Value vs Pointer Return Types:
FunctionReturn TypeWhy?New()(*User, error)Returns pointer for efficiency, error for validationNewAdmin()AdminReturns value (no validation needed, struct is small)
Different approaches:
go// User constructor - returns pointer and error
user, err := user.New("John", "Doe", "01/01/1990")
if err != nil { /* handle */ }

// Admin constructor - returns value, no error
admin := user.NewAdmin("test@example.com", "test123")
// No error checking needed

📊 Memory Layout:
goadmin := NewAdmin("test@example.com", "test123")

// In memory, admin contains:
// ┌─────────────────────────────────┐
// │ email: "test@example.com"       │  Admin's field
// │ password: "test123"             │  Admin's field
// │ ┌─────────────────────────────┐ │
// │ │ firstName: "ADMIN"          │ │  Embedded User
// │ │ lastName: "ADMIN"           │ │  Embedded User
// │ │ birthDate: "---"            │ │  Embedded User
// │ │ createdAt: <timestamp>      │ │  Embedded User
// │ └─────────────────────────────┘ │
// └─────────────────────────────────┘

💡 Usage Examples:
gopackage main

import "example.com/structs/user"

func main() {
    // Create regular user with validation
    u, err := user.New("Alice", "Smith", "05/15/1990")
    if err != nil {
        panic(err)
    }
    u.OutputUserDetails()  // Output: Alice Smith 05/15/1990
    
    // Create admin without validation
    admin := user.NewAdmin("admin@example.com", "securepass")
    admin.OutputUserDetails()  // Output: ADMIN ADMIN ---
    
    // Both can use the same methods
    u.ClearUserName()
    admin.ClearUserName()
    
    // Cannot access private fields from main package
    // fmt.Println(u.firstName)     // ❌ Compile error
    // fmt.Println(admin.email)     // ❌ Compile error
}

🚀 Why Return Value for Admin but Pointer for User?
NewAdmin returns Admin (value):

No validation needed (always succeeds)
No error to return
Admin struct is relatively small
Simpler API (no error handling)

New returns (*User, error) (pointer):

Validation can fail (needs error)
Pointer is more efficient for larger structs
Consistent with Go conventions for fallible operations
Allows nil return on error
