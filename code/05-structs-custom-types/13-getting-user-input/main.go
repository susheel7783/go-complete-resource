package main

import (
	"errors" // Package for creating error values
	"fmt"    // Package for formatted I/O
)

func main() {
	// ==================== MAIN ENTRY POINT ====================
	
	// Call getNoteData() which returns three values:
	// - title (string): the note's title
	// - content (string): the note's content  
	// - err (error): any error that occurred, or nil if successful
	//
	// This is Go's standard error handling pattern:
	// Functions that can fail return (result, error)
	title, content, err := getNoteData()
	
	// ==================== ERROR CHECKING ====================
	// Check if an error occurred during note data collection
	if err != nil {
		// If error exists, print it and exit early
		fmt.Println(err) // Display the error message to user
		return           // Exit main function (ends program)
	}
	
	// If we reach here, note data was successfully collected
	// title and content are valid and can be used
	// (Currently nothing happens with them - placeholder for future logic)
	
	// Future code might do:
	// saveNote(title, content)
	// fmt.Printf("Note saved: %s\n", title)
	
	// Note: Go compiler will warn about unused variables (title, content)
	// In production, you'd use these values
	_ = title   // Suppress unused variable warning
	_ = content // Suppress unused variable warning
}

// ==================== ORCHESTRATION FUNCTION ====================
// getNoteData collects both title and content from the user
// 
// This function demonstrates:
// - Multiple return values (string, string, error)
// - Error propagation (passing errors up the call stack)
// - Early return on error pattern
//
// Returns:
// - string: note title (empty string if error)
// - string: note content (empty string if error)
// - error: nil if successful, error object if failed
func getNoteData() (string, string, error) {
	// ==================== GET TITLE ====================
	// Call getUserInput to get the note title
	// Returns (string, error) - we capture both values
	title, err := getUserInput("Note title:")
	
	// Check if getting title failed
	if err != nil {
		// ERROR PROPAGATION:
		// Return empty strings for title and content
		// Pass the error up to the caller (main function)
		// This is the "fail fast" pattern - don't continue if first step fails
		return "", "", err
	}
	
	// ==================== GET CONTENT ====================
	// If title was successful, get the note content
	// Reuse the 'err' variable (it's nil from previous success)
	content, err := getUserInput("Note content:")
	
	// Check if getting content failed
	if err != nil {
		// ERROR PROPAGATION:
		// Return empty strings and propagate the error
		// We got the title successfully, but content failed
		return "", "", err
	}
	
	// ==================== SUCCESS CASE ====================
	// Both title and content were collected successfully
	// Return the actual values and nil for error (no error occurred)
	return title, content, nil
}

// ==================== INPUT HELPER FUNCTION ====================
// getUserInput prompts the user and validates their input
// 
// This is a reusable function for getting validated user input
// Demonstrates:
// - Single responsibility (does one thing well)
// - Input validation
// - Error creation
//
// Parameters:
// - prompt (string): the message to display to the user
//
// Returns:
// - string: the user's input (empty string if error)
// - error: nil if successful, error object if validation failed
func getUserInput(prompt string) (string, error) {
	// ==================== DISPLAY PROMPT ====================
	// Print the prompt without a newline (so user types on same line)
	fmt.Print(prompt)
	
	// ==================== READ INPUT ====================
	// Declare variable to store user input
	var value string
	
	// Read a line of input from the user
	// Scanln reads until newline and stores in value
	// Note: Scanln stops at whitespace, so "Hello World" becomes just "Hello"
	// For multi-word input, use bufio.Scanner or fmt.Scan with different approach
	fmt.Scanln(&value)
	
	// ==================== VALIDATION ====================
	// Check if input is empty (user just pressed Enter)
	if value == "" {
		// Return empty string and an error
		// errors.New() creates a new error with the given message
		return "", errors.New("Invalid input.")
	}
	
	// ==================== SUCCESS CASE ====================
	// Input is valid (not empty)
	// Return the value and nil error (no error)
	return value, nil
}

// ==================== ERROR HANDLING FLOW ====================
//
// Scenario 1: User provides valid input
// ┌─────────────────────────────────────────────────┐
// │ main()                                          │
// │  ↓ calls                                        │
// │ getNoteData()                                   │
// │  ↓ calls                                        │
// │ getUserInput("Note title:") → "My Title", nil   │
// │  ↓ returns                                      │
// │ title = "My Title", err = nil ✅                │
// │  ↓ calls                                        │
// │ getUserInput("Note content:") → "Content", nil  │
// │  ↓ returns                                      │
// │ content = "Content", err = nil ✅               │
// │  ↓ returns                                      │
// │ getNoteData() → "My Title", "Content", nil      │
// │  ↓ main continues normally                     │
// └─────────────────────────────────────────────────┘
//
// Scenario 2: User provides empty title
// ┌──────────────────────────────────────────────────┐
// │ main()                                           │
// │  ↓ calls                                         │
// │ getNoteData()                                    │
// │  ↓ calls                                         │
// │ getUserInput("Note title:") → "", error ❌       │
// │  ↓ error detected                               │
// │ return "", "", err (propagate error)            │
// │  ↓ returns to main                              │
// │ err != nil → print error and return (exit)      │
// └──────────────────────────────────────────────────┘
//
// Scenario 3: Valid title, empty content
// ┌──────────────────────────────────────────────────┐
// │ main()                                           │
// │  ↓ calls                                         │
// │ getNoteData()                                    │
// │  ↓ calls                                         │
// │ getUserInput("Note title:") → "Title", nil ✅    │
// │  ↓ continues                                     │
// │ getUserInput("Note content:") → "", error ❌     │
// │  ↓ error detected                               │
// │ return "", "", err (propagate error)            │
// │  ↓ returns to main                              │
// │ err != nil → print error and return (exit)      │
// └──────────────────────────────────────────────────┘



// ----------------------
🔑 KEY CONCEPTS:
1. Multiple Return Values
Go functions can return multiple values - commonly used for (result, error):
go// Function signature with multiple returns
func getUserInput(prompt string) (string, error) {
    // Returns TWO values
}

// Calling and capturing both values
value, err := getUserInput("Enter name:")
if err != nil {
    // Handle error
}
// Use value
2. Error Propagation Pattern
gofunc getNoteData() (string, string, error) {
    title, err := getUserInput("Note title:")
    if err != nil {
        return "", "", err  // Propagate error up
    }
    
    content, err := getUserInput("Note content:")
    if err != nil {
        return "", "", err  // Propagate error up
    }
    
    return title, content, nil  // Success
}
Key points:

Errors flow UP the call stack
Each function checks and decides: handle or propagate
Return "zero values" (empty strings) when returning error
Return nil for error when successful

3. Fail Fast Pattern
go// As soon as an error occurs, return immediately
title, err := getUserInput("Note title:")
if err != nil {
    return "", "", err  // Don't continue if first step fails
}

// Only proceed if previous step succeeded
content, err := getUserInput("Note content:")
// ...
```

---

**📊 Sample Program Execution:**

**Scenario 1: Valid Input**
```
Note title:My First Note
Note content:This is the content
(Program exits normally - no output because title/content unused)
```

**Scenario 2: Empty Title**
```
Note title:
Invalid input.
(Program exits early)
```

**Scenario 3: Valid Title, Empty Content**
```
Note title:My Note
Note content:
Invalid input.
(Program exits - title was collected but content failed)

🎯 Common Error Handling Patterns in Go:
Pattern 1: Check and Propagate
goresult, err := someFunction()
if err != nil {
    return nil, err  // Pass error to caller
}
Pattern 2: Check and Wrap
goresult, err := someFunction()
if err != nil {
    return nil, fmt.Errorf("failed to do X: %w", err)  // Add context
}
Pattern 3: Check and Handle
goresult, err := someFunction()
if err != nil {
    log.Println("Warning:", err)
    // Continue with default value
    result = defaultValue
}
Pattern 4: Check and Retry
gofor i := 0; i < 3; i++ {
    result, err := someFunction()
    if err == nil {
        return result, nil  // Success
    }
    // Retry
}
return nil, errors.New("failed after 3 attempts")

💡 Improvements for Production:
gopackage main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("\nNote created successfully!\n")
	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Content: %s\n", content)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title: ")
	if err != nil {
		return "", "", fmt.Errorf("title input failed: %w", err)
	}
	
	content, err := getUserInput("Note content: ")
	if err != nil {
		return "", "", fmt.Errorf("content input failed: %w", err)
	}
	
	return title, content, nil
}

// Improved getUserInput that handles multi-word input
func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	
	// Use bufio.Scanner for better input handling
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", errors.New("failed to read input")
	}
	
	value := strings.TrimSpace(scanner.Text())
	
	if value == "" {
		return "", errors.New("input cannot be empty")
	}
	
	return value, nil
}

