package main

import (
	"bufio"  // Package for buffered I/O (better than fmt.Scan for multi-word input)
	"fmt"    // Package for formatted I/O
	"os"     // Package for operating system functionality (access to stdin)
	"strings" // Package for string manipulation functions
	
	// ==================== CUSTOM PACKAGE IMPORT ====================
	// Import the custom "note" package from the local project
	// "example.com/note/note" is the module path
	// 
	// Project structure:
	// project/
	// ├── go.mod (contains: module example.com/note)
	// ├── main.go (this file)
	// └── note/
	//     └── note.go (contains Note struct and methods)
	"example.com/note/note"
)

func main() {
	// ==================== DATA COLLECTION ====================
	// Call getNoteData() to collect title and content from user
	// 
	// KEY DIFFERENCE from previous version:
	// - NO error returned from getNoteData()
	// - Error handling moved to note.New() instead
	// - getUserInput() now always returns a string (empty if error)
	title, content := getNoteData()
	
	// ==================== NOTE CREATION WITH VALIDATION ====================
	// Create a new Note using the note package's constructor
	// 
	// note.New() performs validation and returns:
	// - *note.Note: pointer to the created Note (or nil if validation fails)
	// - error: nil if successful, error object if validation fails
	//
	// This moves validation logic into the Note type itself
	// which is better encapsulation - Note is responsible for its own validity
	userNote, err := note.New(title, content)
	
	// ==================== ERROR CHECKING ====================
	// Check if Note creation failed (likely due to empty title/content)
	if err != nil {
		fmt.Println(err) // Display validation error message
		return           // Exit program early
	}
	
	// ==================== USING THE NOTE ====================
	// If we reach here, userNote is valid and can be used
	// Call the Display() method to show the note's information
	// This method is defined in the note package
	userNote.Display()
}

// ==================== DATA COLLECTION ORCHESTRATION ====================
// getNoteData collects both title and content from the user
// 
// SIMPLIFIED from previous version:
// - No error handling here (errors handled by note.New() instead)
// - Just collects raw input and returns it
// - Returns two strings (title, content)
//
// This is a simpler design - input collection is separate from validation
func getNoteData() (string, string) {
	// Get note title from user
	title := getUserInput("Note title:")
	
	// Get note content from user
	content := getUserInput("Note content:")
	
	// Return both values (no error to return)
	return title, content
}

// ==================== IMPROVED INPUT FUNCTION ====================
// getUserInput gets a line of input from the user
// 
// MAJOR IMPROVEMENTS over fmt.Scanln:
// - Reads ENTIRE line including spaces ("Hello World" works correctly)
// - Handles newline characters properly across OS (Windows \r\n, Unix \n)
// - Uses buffered I/O for efficiency
//
// Parameters:
// - prompt (string): message to display to the user
//
// Returns:
// - string: user's input (or empty string if error)
//
// Note: No error returned - errors are silently handled by returning ""
// Validation happens later in note.New()
func getUserInput(prompt string) string {
	// ==================== DISPLAY PROMPT ====================
	// Print prompt with a space after it for better formatting
	// %v is a generic verb that works with any type
	fmt.Printf("%v ", prompt)
	
	// ==================== CREATE BUFFERED READER ====================
	// Create a buffered reader that reads from standard input (keyboard)
	// 
	// bufio.Reader provides efficient, buffered reading
	// os.Stdin is the standard input stream (keyboard input)
	// 
	// This is better than fmt.Scanln because:
	// - Can read entire lines with spaces
	// - More efficient for larger inputs
	// - More control over reading behavior
	reader := bufio.NewReader(os.Stdin)
	
	// ==================== READ INPUT LINE ====================
	// Read everything until newline character ('\n')
	// 
	// ReadString('\n') reads input until it encounters newline
	// Returns:
	// - text (string): everything read, INCLUDING the '\n'
	// - err (error): any error that occurred while reading
	//
	// Example: User types "Hello World" and presses Enter
	// text = "Hello World\n" (on Unix) or "Hello World\r\n" (on Windows)
	text, err := reader.ReadString('\n')
	
	// ==================== ERROR HANDLING ====================
	// If reading failed for any reason, return empty string
	// This is a simple error handling approach - just give up and return ""
	// The validation in note.New() will catch the empty string
	if err != nil {
		return ""
	}
	
	// ==================== CLEAN UP INPUT ====================
	// Remove newline characters from the end of the input
	// This is necessary because ReadString includes the delimiter
	
	// Remove Unix/Linux newline character ("\n")
	// TrimSuffix removes the suffix only if it exists at the end
	// Example: "Hello World\n" becomes "Hello World"
	text = strings.TrimSuffix(text, "\n")
	
	// Remove Windows carriage return ("\r")
	// Windows uses "\r\n" for newlines, so we need to remove both
	// After removing "\n" above, we remove "\r" if it exists
	// Example: "Hello World\r" becomes "Hello World"
	text = strings.TrimSuffix(text, "\r")
	
	// ==================== RETURN CLEANED INPUT ====================
	// Return the cleaned text (no trailing newlines)
	// This will be a proper string that can be used in the Note
	return text
}

// ==================== COMPARISON WITH PREVIOUS VERSION ====================
//
// OLD VERSION (with error handling everywhere):
// func getUserInput(prompt string) (string, error) {
//     fmt.Print(prompt)
//     var value string
//     fmt.Scanln(&value)  // ❌ Stops at whitespace!
//     if value == "" {
//         return "", errors.New("Invalid input.")
//     }
//     return value, nil
// }
//
// NEW VERSION (cleaner, no error handling here):
// func getUserInput(prompt string) string {
//     // ... bufio reading ...
//     return text  // Just returns the input
// }
//
// BENEFITS:
// - Separation of concerns: input collection vs validation
// - Handles multi-word input correctly
// - Works across different operating systems
// - Validation moved to note.New() where it belongs


// --------------------
📁 THE NOTE PACKAGE (note/note.go)
Here's what the note/note.go file would contain:
gopackage note

import (
	"errors"
	"fmt"
)

// ==================== NOTE STRUCT ====================
// Note represents a note with title and content
type Note struct {
	title   string  // Private field - note title
	content string  // Private field - note content
}

// ==================== CONSTRUCTOR WITH VALIDATION ====================
// New creates a new Note with validation
// 
// This is where validation happens (not in main.go)
// Better encapsulation - Note is responsible for its own validity
func New(title, content string) (*Note, error) {
	// Validate that both fields are provided
	if title == "" || content == "" {
		return nil, errors.New("Title and content cannot be empty.")
	}
	
	// Create and return the Note
	return &Note{
		title:   title,
		content: content,
	}, nil
}

// ==================== DISPLAY METHOD ====================
// Display shows the note's information
func (n *Note) Display() {
	fmt.Printf("\nYour Note:\n")
	fmt.Printf("Title: %s\n", n.title)
	fmt.Printf("Content: %s\n", n.content)
}

🔑 KEY IMPROVEMENTS:
1. bufio.Reader vs fmt.Scanln
Featurefmt.Scanlnbufio.ReaderMulti-word input❌ Stops at space✅ Reads entire line"Hello World"❌ Gets "Hello"✅ Gets "Hello World"Cross-platform⚠️ Issues with line endings✅ Handles \n and \r\nEfficiency⚠️ OK for small input✅ Buffered, efficient
Example:
go// User types: "My First Note" and presses Enter

// OLD (fmt.Scanln):
fmt.Scanln(&value)  // value = "My" ❌ Only gets first word!

// NEW (bufio.Reader):
reader := bufio.NewReader(os.Stdin)
text, _ := reader.ReadString('\n')  // text = "My First Note\n"
text = strings.TrimSuffix(text, "\n")  // text = "My First Note" ✅
2. Cross-Platform Newline Handling
go// User presses Enter

// Unix/Linux/Mac:
// Input: "Hello\n"

// Windows:
// Input: "Hello\r\n"

// Our code handles both:
text = strings.TrimSuffix(text, "\n")   // Remove \n
text = strings.TrimSuffix(text, "\r")   // Remove \r (if exists)
// Result: "Hello" on all platforms ✅
3. Separation of Concerns
go// INPUT COLLECTION (main.go)
// - Just gets raw input
// - No validation logic
// - Simple and focused
func getUserInput(prompt string) string {
    // Read input
    return text
}

// VALIDATION (note/note.go)
// - Note type validates itself
// - Encapsulation - Note knows its own rules
// - Single source of truth for validation
func New(title, content string) (*Note, error) {
    if title == "" || content == "" {
        return nil, errors.New("...")
    }
    return &Note{...}, nil
}
```

---

**📊 Sample Program Execution:**

**Scenario 1: Valid Multi-Word Input**
```
Note title: My First Note
Note content: This is a great note with multiple words

Your Note:
Title: My First Note
Content: This is a great note with multiple words
```

**Scenario 2: Empty Title**
```
Note title: 
Note content: Some content
Title and content cannot be empty.
```

**Scenario 3: Empty Content**
```
Note title: My Note
Note content: 
Title and content cannot be empty.
```

---

**🎯 Program Flow:**
```
┌─────────────────────────────────────────────────────────────┐
│ main()                                                      │
│  ↓                                                          │
│ getNoteData()                                               │
│  ├─ getUserInput("Note title:")                            │
│  │   └─ bufio.Reader → "My Note" ✅                        │
│  └─ getUserInput("Note content:")                          │
│      └─ bufio.Reader → "Content here" ✅                   │
│  ↓                                                          │
│ Returns: ("My Note", "Content here")                       │
│  ↓                                                          │
│ note.New("My Note", "Content here")                        │
│  ├─ Validates: both non-empty ✅                           │
│  └─ Returns: (*Note{...}, nil)                             │
│  ↓                                                          │
│ userNote.Display()                                          │
│  └─ Prints formatted note                                  │
└─────────────────────────────────────────────────────────────┘

💡 Why This Design is Better:

Robust Input: Handles spaces, special characters, different OS
Clean Separation: Input collection ≠ validation
Encapsulation: Note validates itself, not main.go
User-Friendly: Can enter natural text with spaces
Maintainable: Easy to add more Note fields/validation


🚀 Additional Improvements:
go// Alternative with TrimSpace (removes all whitespace from both ends)
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	
	// TrimSpace removes \n, \r, spaces, tabs from both ends
	return strings.TrimSpace(text)  // Simpler!
}
This is production
