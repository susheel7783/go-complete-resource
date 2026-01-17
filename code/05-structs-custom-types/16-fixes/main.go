package main

import (
	"bufio"  // Package for buffered I/O (efficient reading from stdin)
	"fmt"    // Package for formatted I/O (printing output)
	"os"     // Package for OS functionality (access to stdin for input)
	"strings" // Package for string manipulation (trimming newlines)
	
	// Import custom note package
	// Provides Note type with New(), Display(), and Save() methods
	"example.com/note/note"
)

func main() {
	// ==================== STEP 1: COLLECT USER INPUT ====================
	// Get note title and content from the user
	// getNoteData() returns two strings: (title, content)
	// Uses bufio.Reader to properly handle multi-word input
	title, content := getNoteData()
	
	// ==================== STEP 2: CREATE NOTE WITH VALIDATION ====================
	// Attempt to create a Note using the collected data
	//
	// note.New() performs validation:
	// - Checks that title is not empty
	// - Checks that content is not empty
	// - Sets CreatedAt to current timestamp
	//
	// Returns:
	// - Note: the created note (or empty Note{} if validation fails)
	// - error: nil if successful, error message if validation fails
	userNote, err := note.New(title, content)
	
	// ==================== STEP 3: ERROR HANDLING (CREATION) ====================
	// Check if Note creation failed
	// Common reasons:
	// - User entered empty title (just pressed Enter)
	// - User entered empty content (just pressed Enter)
	if err != nil {
		// Display the error message from note.New()
		// Typically: "Invalid input."
		fmt.Println(err)
		
		// Exit program early - no point continuing without a valid note
		// return exits the main function, terminating the program
		return
	}
	
	// ==================== STEP 4: DISPLAY THE NOTE ====================
	// If we reach here, userNote is valid
	// Display the note to the user for confirmation
	//
	// userNote.Display() shows:
	// "Your note titled [Title] has the following content:
	//
	// [Content]"
	//
	// This gives the user immediate feedback about what was created
	userNote.Display()
	
	// ==================== STEP 5: SAVE NOTE TO FILE ====================
	// Persist the note to disk as a JSON file
	//
	// userNote.Save() does:
	// 1. Generates filename from title (e.g., "shopping_list.json")
	// 2. Serializes Note struct to JSON format
	// 3. Writes JSON to file with 0644 permissions
	//
	// Returns:
	// - error: nil if save successful, error if failed
	//
	// Possible errors:
	// - Disk full
	// - Permission denied
	// - Invalid filename characters
	// - JSON marshaling error (unlikely with simple Note struct)
	err = userNote.Save()
	
	// ==================== STEP 6: ERROR HANDLING (SAVING) ====================
	// Check if saving to file failed
	if err != nil {
		// Display generic error message
		// Note: We could display the actual error with fmt.Println(err)
		// but using a user-friendly message instead
		fmt.Println("Saving the note failed.")
		
		// Exit program
		// The note was created successfully but couldn't be saved
		// User saw the note via Display() but it's not persisted
		return
	}
	
	// ==================== STEP 7: SUCCESS CONFIRMATION ====================
	// If we reach here, everything succeeded:
	// ✅ User provided valid input
	// ✅ Note was created successfully
	// ✅ Note was displayed to user
	// ✅ Note was saved to disk
	//
	// Inform the user that their note is safely saved
	fmt.Println("Saving the note succeeded!")
	
	// Program ends naturally here
	// User has:
	// 1. Seen their note displayed
	// 2. Received confirmation it was saved
	// 3. A JSON file with their note data
}

// ==================== INPUT COLLECTION ORCHESTRATOR ====================
// getNoteData collects both title and content from the user
//
// This function:
// - Separates input collection from validation
// - Handles input in a reusable way
// - Returns raw strings (validation happens in note.New())
//
// Returns:
// - string: note title (may be empty if user just pressed Enter)
// - string: note content (may be empty if user just pressed Enter)
func getNoteData() (string, string) {
	// Collect note title
	// getUserInput() displays prompt and reads entire line (including spaces)
	title := getUserInput("Note title:")
	
	// Collect note content
	// Same process for content
	content := getUserInput("Note content:")
	
	// Return both values
	// No validation here - that's the responsibility of note.New()
	// Separation of concerns: input collection vs validation
	return title, content
}

// ==================== ROBUST INPUT FUNCTION ====================
// getUserInput gets a full line of input from the user
//
// This is a PRODUCTION-QUALITY input function that:
// ✅ Handles multi-word input correctly ("My First Note" works!)
// ✅ Works across different operating systems (Windows, Linux, Mac)
// ✅ Properly removes newline characters (\n and \r)
// ✅ Uses buffered I/O for efficiency
//
// ADVANTAGES OVER fmt.Scanln:
// - fmt.Scanln stops at whitespace → "Hello World" becomes "Hello" ❌
// - bufio.Reader reads entire line → "Hello World" stays "Hello World" ✅
//
// Parameters:
// - prompt (string): message to display to user (e.g., "Note title:")
//
// Returns:
// - string: user's complete input with newlines removed
//           Returns empty string if reading fails
func getUserInput(prompt string) string {
	// ==================== DISPLAY PROMPT ====================
	// Print the prompt with a space after it for better formatting
	// %v is a generic format verb that works with any type
	// Example output: "Note title: " (cursor waits after the space)
	fmt.Printf("%v ", prompt)
	
	// ==================== CREATE BUFFERED READER ====================
	// Create a buffered reader that reads from standard input (keyboard)
	//
	// bufio.Reader provides:
	// - Buffered reading (reads in chunks for efficiency)
	// - Can read entire lines including spaces
	// - Better control over input parsing
	//
	// os.Stdin is the standard input stream (user's keyboard)
	reader := bufio.NewReader(os.Stdin)
	
	// ==================== READ ENTIRE LINE ====================
	// Read everything the user types until they press Enter
	//
	// ReadString('\n') reads until newline character:
	// - '\n' is the newline character (Enter key)
	// - Returns everything INCLUDING the newline
	// - Works with multi-word input
	//
	// Examples:
	// User types: "Hello World" [Enter]
	// text = "Hello World\n" (Unix/Linux/Mac)
	// text = "Hello World\r\n" (Windows)
	//
	// Returns:
	// - text (string): everything read, including newline(s)
	// - err (error): any error that occurred (EOF, read error, etc.)
	text, err := reader.ReadString('\n')
	
	// ==================== ERROR HANDLING ====================
	// Check if reading input failed
	// Possible errors:
	// - EOF (end of file) if input stream closed
	// - I/O errors
	//
	// If error occurs, return empty string
	// The validation in note.New() will catch this and return error
	if err != nil {
		return ""
	}
	
	// ==================== CROSS-PLATFORM NEWLINE REMOVAL ====================
	// Remove newline characters from the end of input
	// This is necessary because different OS use different newline conventions:
	//
	// Unix/Linux/Mac: "\n"     (Line Feed)
	// Windows:        "\r\n"   (Carriage Return + Line Feed)
	// Old Mac:        "\r"     (Carriage Return - rare now)
	//
	// Our strategy: Remove both "\n" and "\r" to work everywhere
	
	// Step 1: Remove "\n" (Unix/Linux/Mac newline, or second part of Windows newline)
	// strings.TrimSuffix removes the suffix ONLY if it exists at the end
	// Example: "Hello World\n" → "Hello World"
	// Example: "Hello World\r\n" → "Hello World\r"
	text = strings.TrimSuffix(text, "\n")
	
	// Step 2: Remove "\r" (Windows carriage return, or old Mac newline)
	// After removing "\n" above, Windows input has "\r" remaining
	// Example: "Hello World\r" → "Hello World"
	text = strings.TrimSuffix(text, "\r")
	
	// Now text is clean on all platforms!
	// Result: "Hello World" (no newline characters)
	
	// Alternative approach (same result):
	// text = strings.TrimSpace(text)
	// TrimSpace removes ALL whitespace from both ends (\n, \r, spaces, tabs)
	
	// ==================== RETURN CLEANED INPUT ====================
	// Return the user's input with newlines removed
	// This is a clean string ready to be used in the Note
	return text
}

// ==================== PROGRAM FLOW VISUALIZATION ====================
//
// Successful Flow:
// ┌─────────────────────────────────────────────────────────────────┐
// │ 1. User Input                                                   │
// │    Note title: My Shopping List                                 │
// │    Note content: Buy milk, eggs, and bread                      │
// ├─────────────────────────────────────────────────────────────────┤
// │ 2. Create Note                                                  │
// │    note.New("My Shopping List", "Buy milk, eggs, and bread")    │
// │    ✅ Validation passes                                         │
// │    ✅ Note created with current timestamp                       │
// ├─────────────────────────────────────────────────────────────────┤
// │ 3. Display Note                                                 │
// │    Your note titled My Shopping List has the following content: │
// │                                                                 │
// │    Buy milk, eggs, and bread                                    │
// ├─────────────────────────────────────────────────────────────────┤
// │ 4. Save Note                                                    │
// │    Filename: my_shopping_list.json                              │
// │    JSON: {"Title":"My Shopping List","Content":"..."}           │
// │    ✅ File written successfully                                 │
// ├─────────────────────────────────────────────────────────────────┤
// │ 5. Success Message                                              │
// │    Saving the note succeeded!                                   │
// └─────────────────────────────────────────────────────────────────┘
//
// Error Flow 1: Empty Input
// ┌─────────────────────────────────────────────────────────────────┐
// │ 1. User Input                                                   │
// │    Note title: [Enter]                                          │
// │    Note content: Some content                                   │
// ├─────────────────────────────────────────────────────────────────┤
// │ 2. Create Note                                                  │
// │    note.New("", "Some content")                                 │
// │    ❌ Validation fails (empty title)                            │
// │    ❌ Returns error: "Invalid input."                           │
// ├─────────────────────────────────────────────────────────────────┤
// │ 3. Error Display                                                │
// │    Invalid input.                                               │
// │    [Program exits]                                              │
// └─────────────────────────────────────────────────────────────────┘
//
// Error Flow 2: Save Failure
// ┌─────────────────────────────────────────────────────────────────┐
// │ 1-3. [Same as successful flow]                                  │
// ├─────────────────────────────────────────────────────────────────┤
// │ 4. Save Note                                                    │
// │    ❌ Disk full / Permission denied / Other I/O error           │
// │    ❌ Returns error                                             │
// ├─────────────────────────────────────────────────────────────────┤
// │ 5. Error Display                                                │
// │    Saving the note failed.                                      │
// │    [Program exits]                                              │
// └─────────────────────────────────────────────────────────────────┘

// ==================== EXAMPLE OUTPUT ====================
//
// Terminal Session:
// ─────────────────────────────────────────────────────────
// $ go run main.go
// Note title: Shopping List
// Note content: Buy milk, eggs, and bread
// 
// Your note titled Shopping List has the following content:
//
// Buy milk, eggs, and bread
//
// Saving the note succeeded!
// ─────────────────────────────────────────────────────────
//
// File Created: shopping_list.json
// ─────────────────────────────────────────────────────────
// {
//   "Title": "Shopping List",
//   "Content": "Buy milk, eggs, and bread",
//   "CreatedAt": "2025-01-17T15:45:30.123456789Z"
// }
// ─────────────────────────────────────────────────────────

// ==================== KEY IMPROVEMENTS FROM EARLIER VERSIONS ====================
//
// 1. ✅ Multi-word input support
//    - OLD: fmt.Scanln → "My Note" becomes "My"
//    - NEW: bufio.Reader → "My Note" stays "My Note"
//
// 2. ✅ Cross-platform newline handling
//    - Works on Windows (\r\n), Unix (\n), Mac (\r)
//
// 3. ✅ Note is actually used
//    - Display() shows the note to user
//    - Save() persists to disk
//
// 4. ✅ User feedback
//    - Shows what was created
//    - Confirms successful save
//    - Clear error messages
//
// 5. ✅ Error handling at each step
//    - Creation validation
//    - Save operation
//    - Graceful exits with messages
//
// 6. ✅ Separation of concerns
//    - Input collection (main.go)
//    - Validation (note.New)
//    - Display (note.Display)
//    - Persistence (note.Save)


// --------------
🔑 KEY FEATURES:
1. Complete Error Handling
go// Error handling at EVERY step:

// Step 1: Note creation
userNote, err := note.New(title, content)
if err != nil {
    fmt.Println(err)  // "Invalid input."
    return
}

// Step 2: File save
err = userNote.Save()
if err != nil {
    fmt.Println("Saving the note failed.")
    return
}
2. Cross-Platform Input
go// Handles different OS newline conventions:

// Windows: "Hello\r\n"
text = strings.TrimSuffix(text, "\n")  // → "Hello\r"
text = strings.TrimSuffix(text, "\r")  // → "Hello"

// Unix/Linux/Mac: "Hello\n"
text = strings.TrimSuffix(text, "\n")  // → "Hello"
text = strings.TrimSuffix(text, "\r")  // → "Hello" (no change)
3. User Feedback Loop
go1. User enters data
2. Display() shows what was created → User confirms it's correct
3. Save() persists to disk
4. Success message → User knows it's saved
```

---

**📊 Sample Sessions:**

**Session 1: Success**
```
Note title: My Ideas
Note content: Learn Go programming and build projects

Your note titled My Ideas has the following content:

Learn Go programming and build projects

Saving the note succeeded!
```
**Result:** File `my_ideas.json` created

**Session 2: Empty Title**
```
Note title: 
Note content: Some content
Invalid input.
```
**Result:** No file created, program exits

**Session 3: Multi-word Title**
```
Note title: Project Planning Notes
Note content: Design database schema and API endpoints

Your note titled Project Planning Notes has the following content:

Design database schema and API endpoints

Saving the note succeeded!
Result: File project_planning_notes.json created
