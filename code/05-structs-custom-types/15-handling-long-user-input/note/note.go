package note

import (
	"errors" // Package for creating error values
	"fmt"    // Package for formatted I/O (printing)
	"time"   // Package for working with dates and times
)

// ==================== NOTE STRUCT DEFINITION ====================
// Note represents a note with title, content, and creation timestamp
// 
// ENCAPSULATION:
// - All fields are UNEXPORTED (lowercase) - private to this package
// - External packages cannot directly access or modify these fields
// - Must use New() to create and Display() to view
//
// This ensures:
// - Notes are always created with valid data (via New() validation)
// - createdAt is always set automatically (users can't forget or set wrong value)
// - No external code can corrupt the Note's state
type Note struct {
	title     string    // Private: the note's title
	content   string    // Private: the note's content
	createdAt time.Time // Private: timestamp when note was created
}

// ==================== DISPLAY METHOD ====================
// Display shows the note's information in a formatted way
// 
// METHOD SIGNATURE BREAKDOWN:
// - (note Note) = VALUE RECEIVER (receives a copy of the Note)
// - Display = method name (capitalized = EXPORTED, callable from other packages)
// - () = no parameters
//
// VALUE RECEIVER vs POINTER RECEIVER:
// - Using (note Note) instead of (note *Note)
// - This creates a COPY of the Note when method is called
// - OK here because:
//   * We're only READING data, not modifying it
//   * Note struct is relatively small (copying is cheap)
//   * We don't need to modify the original Note
//
// If you needed to modify Note, you'd use pointer receiver: (note *Note)
func (note Note) Display() {
	// Print formatted note information
	// 
	// Printf format verbs:
	// - %v = default format (prints the value as-is)
	// - \n = newline character
	//
	// Output example:
	// Your note titled My First Note has the following content:
	//
	// This is the content of my note
	//
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", 
		note.title,   // Access private field (we're inside the package)
		note.content) // Access private field
	
	// Note: createdAt is not displayed, but could be added:
	// fmt.Printf("Created at: %v\n", note.createdAt.Format("2006-01-02 15:04:05"))
}

// ==================== CONSTRUCTOR FUNCTION ====================
// New creates and returns a new Note with validation
// 
// NAMING CONVENTION:
// - Named "New" (not "NewNote") because it's in the note package
// - Called as: note.New(...) from other packages
// - This is the standard Go convention for constructors in dedicated packages
//
// RETURN TYPE:
// - Returns (Note, error) - a Note VALUE, not a pointer
// - Different from common pattern of (*Note, error)
// - Both approaches are valid:
//   * Returning value: simpler, good for small structs
//   * Returning pointer: more efficient for large structs, enables mutation
//
// Parameters:
// - title (string): the note's title
// - content (string): the note's content
//
// Returns:
// - Note: the created Note (or empty Note{} if validation fails)
// - error: nil if successful, error object if validation fails
func New(title, content string) (Note, error) {
	// ==================== VALIDATION ====================
	// Check if either required field is empty
	// Using || (OR operator) - if ANY field is empty, validation fails
	//
	// This ensures:
	// - No note can exist with empty title
	// - No note can exist with empty content
	// - Data integrity is maintained at creation time
	if title == "" || content == "" {
		// Return empty Note struct and error
		// Note{} creates a zero-value Note (all fields are zero/empty)
		return Note{}, errors.New("Invalid input.")
		
		// Could provide more specific error message:
		// return Note{}, errors.New("Title and content are required.")
	}
	
	// ==================== NOTE CREATION ====================
	// If validation passes, create and return the Note
	//
	// Note the explicit field initialization:
	// - title and content come from parameters
	// - createdAt is automatically set to current time
	//   (users don't provide this - it's controlled internally)
	//
	// This ensures CONSISTENCY - all Notes have a proper creation timestamp
	return Note{
		title:     title,      // Set from parameter
		content:   content,    // Set from parameter
		createdAt: time.Now(), // Automatically set to current date/time
	}, nil // Return nil error (no error occurred)
	
	// Note: We return the Note value directly, not &Note{...}
	// This means the struct is copied when returned
	// For small structs like this, copying is fine
}

// ==================== WHAT THIS PACKAGE PROVIDES ====================
//
// PUBLIC API (Exported - accessible from other packages):
// - Note type (can declare variables, use as parameter types)
// - New() function (create Notes)
// - Display() method (view Note contents)
//
// PRIVATE IMPLEMENTATION (Unexported - hidden from other packages):
// - title field (cannot access directly)
// - content field (cannot access directly)  
// - createdAt field (cannot access directly)
//
// USAGE FROM MAIN.GO:
// import "example.com/note/note"
//
// // Create note (only way to create a valid Note)
// myNote, err := note.New("Title", "Content")
// if err != nil {
//     // Handle validation error
// }
//
// // Display note (only way to see Note data)
// myNote.Display()
//
// // CANNOT DO (fields are private):
// // fmt.Println(myNote.title)     // ❌ Compile error
// // myNote.content = "Changed"    // ❌ Compile error
// // myNote.createdAt = time.Now() // ❌ Compile error

// ==================== POTENTIAL ENHANCEMENTS ====================
//
// Additional methods you might add:
//
// // Getter methods (Go convention: no "Get" prefix)
// func (n Note) Title() string {
//     return n.title
// }
//
// func (n Note) Content() string {
//     return n.content
// }
//
// func (n Note) CreatedAt() time.Time {
//     return n.createdAt
// }
//
// // Save note to file
// func (n Note) Save(filename string) error {
//     data := fmt.Sprintf("Title: %s\nContent: %s\nCreated: %v",
//         n.title, n.content, n.createdAt)
//     return os.WriteFile(filename, []byte(data), 0644)
// }
//
// // Update methods (would need pointer receiver)
// func (n *Note) UpdateTitle(newTitle string) error {
//     if newTitle == "" {
//         return errors.New("Title cannot be empty")
//     }
//     n.title = newTitle
//     return nil
// }
//
// func (n *Note) UpdateContent(newContent string) error {
//     if newContent == "" {
//         return errors.New("Content cannot be empty")
//     }
//     n.content = newContent
//     return nil
// }

// --------------
🔑 KEY DESIGN DECISIONS:
1. Value Return vs Pointer Return
This Package (Returns Value):
gofunc New(title, content string) (Note, error) {
    return Note{...}, nil  // Returns a copy of the struct
}

// Usage:
note, err := note.New("Title", "Content")
Alternative (Returns Pointer):
gofunc New(title, content string) (*Note, error) {
    return &Note{...}, nil  // Returns a pointer to the struct
}

// Usage:
note, err := note.New("Title", "Content")  // note is *Note
When to use each:

Value: Small structs, immutable data, simpler API
Pointer: Large structs, need mutation, consistent with other methods

2. Value Receiver vs Pointer Receiver
Current (Value Receiver):
gofunc (note Note) Display() {  // Receives a copy
    fmt.Println(note.title)   // Read-only access
}
Alternative (Pointer Receiver):
gofunc (note *Note) Display() {  // Receives a pointer
    fmt.Println(note.title)    // Can read AND modify
}
Best Practice:

Use pointer receivers for ALL methods on a type (for consistency)
Or use value receivers for small, immutable types
Don't mix - be consistent within a type


📊 Complete Usage Example:
gopackage main

import (
    "fmt"
    "example.com/note/note"
)

func main() {
    // ==================== SCENARIO 1: Valid Note ====================
    myNote, err := note.New("Shopping List", "Milk, Eggs, Bread")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    myNote.Display()
    // Output:
    // Your note titled Shopping List has the following content:
    //
    // Milk, Eggs, Bread
    //
    
    // ==================== SCENARIO 2: Empty Title ====================
    invalidNote, err := note.New("", "Some content")
    if err != nil {
        fmt.Println("Error:", err)  // Output: Error: Invalid input.
        return
    }
    
    // ==================== SCENARIO 3: Empty Content ====================
    invalidNote2, err := note.New("Title", "")
    if err != nil {
        fmt.Println("Error:", err)  // Output: Error: Invalid input.
        return
    }
    
    // ==================== CANNOT ACCESS PRIVATE FIELDS ====================
    // fmt.Println(myNote.title)      // ❌ Compile error
    // myNote.content = "New content" // ❌ Compile error
    // myNote.createdAt = time.Now()  // ❌ Compile error
}

🎯 Encapsulation Benefits:
go// ✅ VALID - Using public API
note, _ := note.New("Title", "Content")
note.Display()

// ❌ INVALID - Cannot break encapsulation
note.title = ""              // Compile error - field is private
note.createdAt = pastTime    // Compile error - field is private

// ✅ ENFORCED INVARIANTS
// - Notes always have title and content (validated in New)
// - createdAt always set correctly (can't be manually set wrong)
// - No way to create invalid Note from outside package

💡 Improved Version with Pointer Consistency:
gopackage note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

// Pointer receiver for consistency (even though read-only)
func (n *Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", 
		n.title, n.content)
}

// Return pointer for consistency with pointer receivers
func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("Invalid input.")
	}
	
	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

// Additional methods with pointer receivers
func (n *Note) Title() string {
	return n.title
}

func (n *Note) Content() string {
	return n.content
}

func (n *Note) CreatedAt() time.Time {
	return n.createdAt
}

🚀 Summary:
FeatureImplementationBenefitPrivate FieldsLowercase namesData protection, encapsulationConstructorNew() functionValidation at creation, consistencyAuto-timestampcreatedAt: time.Now()No manual errors, always accurateValue Return(Note, error)Simple for small structsValue Receiver(note Note)Safe for read-only operationsValidationCheck in New()Ensures all Notes are valid
This is a well-designed, encapsulated Go package following best practices! 🎉
