package main

import "fmt" // Standard library for formatted I/O

// ==================== CUSTOM TYPE DEFINITION ====================
// Define a NEW type called "str" based on the built-in "string" type
// 
// This is called a "type definition" or "type alias with a new type"
// Syntax: type NewTypeName ExistingType
// 
// IMPORTANT:
// - "str" is a DIFFERENT type from "string" (not just an alias)
// - You cannot directly assign string to str without conversion
// - BUT you can add methods to "str" (cannot add methods to built-in "string")
// 
// This is useful when you want to:
// - Add behavior (methods) to simple types
// - Create domain-specific types (Email, Username, etc.)
// - Add type safety (Age, Distance with specific units)
type str string

// ==================== METHOD ON CUSTOM TYPE ====================
// Add a method to the str type
// 
// Method signature breakdown:
// - (text str) = receiver (the str value this method operates on)
// - log = method name
// - () = no parameters
// 
// This method can ONLY be called on str type, not on regular string
// 
// Note: Uses VALUE receiver (str, not *str) because:
// - Strings are immutable in Go (cannot be modified)
// - Strings are efficiently passed by value
// - No need for pointer receiver
func (text str) log() {
	// Print the str value to the console
	// "text" contains the str value that log() was called on
	fmt.Println(text)
}

func main() {
	// ==================== CREATING AND USING CUSTOM TYPE ====================
	
	// Declare a variable of type "str" (our custom type)
	// Type annotation "str" is required here for clarity
	var name str = "Max"
	
	// Alternative shorter syntax (type inference):
	// name := str("Max")  // Explicit conversion from string to str
	
	// ==================== CALLING THE METHOD ====================
	// Call the log() method on the str value
	// This works because we defined log() with str as the receiver
	name.log()  // Output: Max
	
	// ==================== WHAT DOESN'T WORK ====================
	
	// Regular strings don't have the log() method:
	// var regularString string = "Hello"
	// regularString.log()  // ❌ ERROR: string type has no field or method log
	
	// Type safety - str and string are different types:
	// var name str = "Max"
	// var text string = name  // ❌ ERROR: cannot use name (type str) as type string
	
	// Need explicit conversion:
	// var text string = string(name)  // ✅ OK: explicit conversion
}


// ----------
🔑 KEY CONCEPTS:
1. Custom Type vs Type Alias
Custom Type (New Type):
gotype str string  // Creates a NEW type based on string

var s str = "hello"
var x string = "world"
// s = x  // ❌ ERROR: different types
s = str(x)  // ✅ OK: explicit conversion needed
Type Alias (Same Type):
gotype str = string  // Creates an ALIAS (notice the =)

var s str = "hello"
var x string = "world"
s = x  // ✅ OK: same type, no conversion needed
2. Why Create Custom Types?
Benefits:

✅ Add methods to simple types
✅ Add type safety and semantic meaning
✅ Prevent mixing up similar values
✅ Create domain-specific abstractions

Real-World Examples:
go// Email type with validation
type Email string

func (e Email) IsValid() bool {
    return strings.Contains(string(e), "@")
}

// Temperature with unit conversion
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
    return float64(c)*9/5 + 32
}

// UserID for type safety
type UserID int

func (id UserID) IsAdmin() bool {
    return id < 100  // First 100 IDs are admins
}

3. Value Receiver vs Pointer Receiver
go// Value receiver - receives a COPY of str
func (text str) log() {
    fmt.Println(text)
}

// For strings, value receiver is preferred because:
// - Strings are immutable (cannot be changed anyway)
// - Small memory overhead
// - Simpler to use

📊 Complete Working Example:
gopackage main

import "fmt"

// Custom string type
type str string

// Method to print the string
func (text str) log() {
	fmt.Println(text)
}

// Method to get length
func (text str) length() int {
	return len(text)  // Need to convert to string? No! Go handles it
}

// Method to uppercase (returns new str)
func (text str) upper() str {
	// strings.ToUpper requires string, so convert
	// Then convert back to str
	// import "strings" needed
	return str(text + "!")  // Simple example
}

func main() {
	// Create custom type value
	var name str = "Max"
	
	// Call methods
	name.log()                    // Output: Max
	fmt.Println(name.length())    // Output: 3
	
	// Method chaining style
	newName := name.upper()
	newName.log()                 // Output: Max!
	
	// ==================== TYPE CONVERSION ====================
	
	// str to string
	regularString := string(name)
	fmt.Println(regularString)    // Output: Max
	
	// string to str
	anotherStr := str("Hello")
	anotherStr.log()              // Output: Hello
}
```

**Output:**
```
Max
3
Max!
Max
Hello

🎯 Practical Real-World Example:
gopackage main

import (
	"fmt"
	"strings"
)

// Custom type for email addresses
type Email string

// Validate email format
func (e Email) IsValid() bool {
	s := string(e)
	return strings.Contains(s, "@") && strings.Contains(s, ".")
}

// Get domain part
func (e Email) Domain() string {
	s := string(e)
	parts := strings.Split(s, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

// Log the email (with privacy)
func (e Email) LogSecure() {
	if len(e) > 3 {
		fmt.Printf("%s***@%s\n", string(e[:2]), e.Domain())
	}
}

func main() {
	// Create Email type
	email := Email("john.doe@example.com")
	
	// Use methods
	fmt.Println("Valid?", email.IsValid())    // Output: Valid? true
	fmt.Println("Domain:", email.Domain())    // Output: Domain: example.com
	email.LogSecure()                         // Output: jo***@example.com
	
	// Type safety prevents mixing up with regular strings
	// var name string = email  // ❌ ERROR: type mismatch
	var name string = string(email)  // ✅ OK: explicit conversion
	fmt.Println(name)
}

💡 Common Use Cases:

Domain Types:

gotype Username string
type Password string
type ProductID int

// Prevents accidentally using password as username
func Login(user Username, pass Password) { }

Units and Measurements:

gotype Kilometers float64
type Miles float64

func (km Kilometers) ToMiles() Miles {
    return Miles(float64(km) * 0.621371)
}

String Enhancements:

gotype JSONString string

func (j JSONString) Parse() (map[string]interface{}, error) {
    // JSON parsing logic
}
