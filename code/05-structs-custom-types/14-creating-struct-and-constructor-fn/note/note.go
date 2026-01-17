package note

import (
	"errors" // Package for creating error values
	"time"   // Package for working with dates and times
)

// ==================== NOTE STRUCT DEFINITION ====================
// Note represents a note with title, content, and creation timestamp
// 
// ENCAPSULATION THROUGH UNEXPORTED FIELDS:
// - All fields are lowercase (unexported/private)
// - Cannot be accessed from other packages
// - This protects data integrity
//
// Benefits of private fields:
// ✅ External code cannot create invalid Notes (empty title/content)
// ✅ External code cannot tamper with createdAt timestamp
// ✅ Only way to create Note is through New() which validates
// ✅ Fields can only be accessed/modified within this package
type Note struct {
	title     string    // Private: the note's title
	content   string    // Private: the note's content  
	createdAt time.Time // Private: timestamp when note was created
}

// ==================== CONSTRUCTOR FUNCTION ====================
// New creates and returns a new Note with validation
// 
// This is the ONLY way external packages should create Notes
// Ensures all Notes are created in a valid state
//
// FUNCTION SIGNATURE BREAKDOWN:
// - New = function name (capitalized = EXPORTED/public)
// - (title, content string) = parameters (both are strings)
// - (Note, error) = TWO return values (the Note and any error)
//
// RETURN TYPE - VALUE vs POINTER:
// Returns (Note, error) - a Note VALUE, not a pointer (*Note)
// 
// Why return value instead of pointer?
// ✅ Simpler API - no need to worry about nil pointers
// ✅ Note is small (3 fields) - copying is cheap
// ✅ Immutable design - once created, Note doesn't change
// ✅ Value semantics - each variable gets its own copy
//
// Alternative would be (*Note, error) returning a pointer
// Both approaches are valid - this is a design choice
//
// Parameters:
// - title: the note's title (cannot be empty)
// - content: the note's content (cannot be empty)
//
// Returns:
// - Note: the created Note struct (or empty Note{} if validation fails)
// - error: nil if successful, error object with message if validation fails
func New(title, content string) (Note, error) {
	// ==================== VALIDATION LOGIC ====================
	// Validate that both required fields are non-empty
	// 
	// The || operator means OR:
	// - Returns true if EITHER condition is true
	// - If title is empty OR content is empty, validation fails
	//
	// This enforces business rules:
	// ✅ Every Note must have a title
	// ✅ Every Note must have content
	// ✅ No half-complete Notes can exist
	if title == "" || content == "" {
		// Validation failed - return empty Note and error
		//
		// Note{} creates a zero-value Note:
		// - title: "" (empty string)
		// - content: "" (empty string)
		// - createdAt: zero time (January 1, year 1, 00:00:00 UTC)
		//
		// errors.New() creates a new error with the given message
		// The caller will receive this error and can handle it
		return Note{}, errors.New("Invalid input.")
		
		// Could be more specific:
		// return Note{}, errors.New("Title and content are required.")
		// Or even more detailed:
		// if title == "" {
		//     return Note{}, errors.New("Title cannot be empty.")
		// }
		// if content == "" {
		//     return Note{}, errors.New("Content cannot be empty.")
		// }
	}
	
	// ==================== NOTE CREATION ====================
	// If validation passes, create the Note with provided values
	//
	// Struct literal syntax:
	// Note{ fieldName: value, ... }
	//
	// Key points:
	// 1. title and content come from function parameters
	// 2. createdAt is AUTOMATICALLY set to current time
	//    - Users cannot set this themselves
	//    - Guarantees accurate timestamp
	//    - time.Now() returns current date and time
	// 3. All fields are explicitly initialized
	//
	// Return the Note VALUE (a copy) and nil error
	// nil means "no error occurred" - everything is OK
	return Note{
		title:     title,      // Assign parameter to field
		content:   content,    // Assign parameter to field
		createdAt: time.Now(), // Auto-set to current timestamp
	}, nil // nil error = success
	
	// Note: The Note is returned by VALUE, not by reference
	// This means the caller gets a COPY of the Note
	// Changes to the copy won't affect the original (immutability)
}

// ==================== WHAT'S MISSING (Intentionally Minimal) ====================
//
// This package has NO methods for displaying, saving, or modifying Notes
// This is a MINIMAL design - just defines the type and constructor
//
// The calling code (main.go) would handle:
// - Displaying notes
// - Saving notes to files
// - Any other operations
//
// Benefits of minimal design:
// ✅ Single Responsibility - Note just represents data
// ✅ Flexibility - Calling code decides how to use Note
// ✅ Simplicity - Easy to understand and maintain
//
// If you wanted to add methods, you could add:
// - Display() method to show the note
// - Save() method to save to file
// - Title(), Content(), CreatedAt() getters
// - UpdateTitle(), UpdateContent() setters (would need pointer receivers)

// ==================== USAGE FROM MAIN.GO ====================
//
// import "example.com/note/note"
//
// // Create a valid note
// myNote, err := note.New("Shopping", "Buy milk and eggs")
// if err != nil {
//     fmt.Println("Error:", err)
//     return
// }
// // myNote is now a valid Note with all fields set
//
// // Try to create invalid note (empty title)
// badNote, err := note.New("", "Content")
// if err != nil {
//     fmt.Println("Error:", err)  // Prints: Error: Invalid input.
//     return
// }
//
// // CANNOT access private fields from main package:
// // fmt.Println(myNote.title)     // ❌ Compile error
// // myNote.content = "new"        // ❌ Compile error  
// // myNote.createdAt = time.Now() // ❌ Compile error
//
// // To display Note, main.go would need to implement its own logic
// // Or this package would need to add a Display() method

// ==================== DESIGN PATTERNS DEMONSTRATED ====================
//
// 1. ENCAPSULATION
//    - Private fields protect data
//    - Public constructor controls creation
//
// 2. VALIDATION AT CREATION
//    - Invalid Notes cannot exist
//    - Fail fast - errors detected immediately
//
// 3. IMMUTABILITY (by design)
//    - No setter methods
//    - No pointer receiver methods that modify
//    - Once created, Note cannot change
//
// 4. CONSTRUCTOR PATTERN
//    - New() function is the "constructor"
//    - Standard Go idiom for object creation
//
// 5. ERROR HANDLING
//    - Returns (result, error) tuple
//    - Caller must check error explicitly
//    - Idiomatic Go error handling

// ==================== VALUE SEMANTICS ====================
//
// Because New() returns Note (value) not *Note (pointer):
//
// note1, _ := note.New("Title", "Content")
// note2 := note1  // note2 is a COPY of note1
//
// // note1 and note2 are independent
// // Modifying note2 won't affect note1 (if there were setters)
//
// This is VALUE SEMANTICS - each variable owns its own data
//
// Compare to POINTER SEMANTICS:
// note1, _ := note.New("Title", "Content")  // Returns *Note
// note2 := note1  // note2 points to SAME Note as note1
// // Modifying through note2 would affect note1
