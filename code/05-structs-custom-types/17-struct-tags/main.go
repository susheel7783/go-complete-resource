package main

import (
	"bufio"  // Buffered I/O for efficient multi-word input reading
	"fmt"    // Formatted I/O for printing prompts and messages
	"os"     // OS functionality for accessing stdin (keyboard input)
	"strings" // String manipulation for cleaning input (removing newlines)
	
	// Custom note package import
	// Provides Note type with New(), Display(), and Save() methods
	// Located at: project/note/note.go
	"example.com/note/note"
)

// ==================== MAIN ENTRY POINT ====================
// This is a complete note-taking application that:
// 1. Collects user input (title and content)
// 2. Creates a validated Note
// 3. Displays the note to the user
// 4. Saves the note as a JSON file
// 5. Provides appropriate error handling and user feedback
func main() {
	// ==================== STEP 1: INPUT COLLECTION ====================
	// Collect note title and content from user
	// Returns two strings - validation happens later in note.New()
	title, content := getNoteData()
	
	// ==================== STEP 2: NOTE CREATION & VALIDATION ====================
	// Create Note with validation
	// note.New() checks:
	// - title is not empty
	// - content is not empty
	// - Sets CreatedAt to current timestamp
	//
	// Returns: (Note, error)
	// - Note: Valid note object (or empty Note{} if validation fails)
	// - error: nil if success, error message if validation fails
	userNote, err := note.New(title, content)
	
	// Check if note creation failed
	if err != nil {
		fmt.Println(err) // Display error (e.g., "Invalid input.")
		return           // Exit early - cannot continue without valid note
	}
	
	// ==================== STEP 3: DISPLAY NOTE ====================
	// Show the created note to the user for confirmation
	// Output format:
	// "Your note titled [Title] has the following content:
	//
	// [Content]"
	//
	// This provides immediate feedback before saving
	userNote.Display()
	
	// ==================== STEP 4: SAVE TO FILE ====================
	// Persist note to disk as JSON file
	// Save() does:
	// 1. Generate filename: "shopping_list.json"
	// 2. Serialize to JSON with snake_case fields
	// 3. Write to file with 0644 permissions
	//
	// Returns error if save fails (disk full, permissions, etc.)
	err = userNote.Save()
	
	// Check if saving failed
	if err != nil {
		fmt.Println("Saving the note failed.")
		return // Exit - note was displayed but not saved
	}
	
	// ==================== STEP 5: SUCCESS CONFIRMATION ====================
	// All steps completed successfully
	// User knows their note is safely stored
	fmt.Println("Saving the note succeeded!")
}

// ==================== INPUT ORCHESTRATION FUNCTION ====================
// getNoteData collects both title and content from the user
//
// Design pattern: Separation of concerns
// - This function handles input collection
// - Validation happens in note.New()
// - No error handling here (errors caught later)
//
// Returns:
// - string: note title (may be empty)
// - string: note content (may be empty)
func getNoteData() (string, string) {
	// Collect title using reusable input function
	title := getUserInput("Note title:")
	
	// Collect content using same function
	content := getUserInput("Note content:")
	
	// Return raw strings - no validation here
	return title, content
}

// ==================== ROBUST INPUT HELPER FUNCTION ====================
// getUserInput reads a complete line of input from the user
//
// KEY FEATURES:
// ✅ Handles multi-word input ("My First Note" works correctly)
// ✅ Cross-platform (Windows, Unix, Mac)
// ✅ Removes newline characters properly
// ✅ Uses buffered I/O for efficiency
//
// ADVANTAGES over fmt.Scanln:
// ❌ fmt.Scanln: "Hello World" → "Hello" (loses "World")
// ✅ bufio.Reader: "Hello World" → "Hello World" (complete)
//
// Parameters:
// - prompt: Message to display (e.g., "Note title:")
//
// Returns:
// - string: Complete user input with newlines removed
//           Empty string if reading fails
func getUserInput(prompt string) string {
	// ==================== DISPLAY PROMPT ====================
	// Print prompt with trailing space for better UX
	// Example: "Note title: " (cursor waits after space)
	fmt.Printf("%v ", prompt)
	
	// ==================== CREATE BUFFERED READER ====================
	// bufio.NewReader creates an efficient buffered reader
	// os.Stdin is the standard input stream (keyboard)
	//
	// Benefits of bufio.Reader:
	// - Reads entire lines including spaces
	// - More efficient than reading byte-by-byte
	// - Better control over input parsing
	reader := bufio.NewReader(os.Stdin)
	
	// ==================== READ COMPLETE LINE ====================
	// ReadString('\n') reads until newline (Enter key)
	// Returns everything INCLUDING the newline character
	//
	// Examples:
	// Unix/Mac:  "Hello World\n"
	// Windows:   "Hello World\r\n"
	//
	// Returns: (text string, err error)
	text, err := reader.ReadString('\n')
	
	// ==================== ERROR HANDLING ====================
	// If reading fails, return empty string
	// Validation in note.New() will catch this
	if err != nil {
		return ""
	}
	
	// ==================== CROSS-PLATFORM NEWLINE CLEANUP ====================
	// Remove newline characters for all operating systems
	//
	// Different OS use different newline conventions:
	// - Unix/Linux/Mac: "\n"   (Line Feed)
	// - Windows:        "\r\n" (Carriage Return + Line Feed)
	//
	// Strategy: Remove both to work everywhere
	
	// Step 1: Remove "\n" (Unix) or second part of Windows "\r\n"
	// "Hello World\n" → "Hello World"
	// "Hello World\r\n" → "Hello World\r"
	text = strings.TrimSuffix(text, "\n")
	
	// Step 2: Remove "\r" (Windows carriage return)
	// "Hello World\r" → "Hello World"
	text = strings.TrimSuffix(text, "\r")
	
	// Result: Clean text on all platforms
	// Alternative: strings.TrimSpace(text) removes all whitespace
	
	return text
}

// ==================== COMPLETE SYSTEM OVERVIEW ====================
//
// PROJECT STRUCTURE:
// ┌─────────────────────────────────────────────────────────┐
// │ project/                                                │
// │ ├── go.mod (module example.com/note)                    │
// │ ├── main.go (this file - user interface)               │
// │ └── note/                                               │
// │     └── note.go (Note type, business logic)             │
// └─────────────────────────────────────────────────────────┘
//
// ARCHITECTURE:
// ┌─────────────────────────────────────────────────────────┐
// │ PRESENTATION LAYER (main.go)                            │
// │ - User interaction                                      │
// │ - Input collection                                      │
// │ - Error messages                                        │
// │ - Success feedback                                      │
// ├─────────────────────────────────────────────────────────┤
// │ BUSINESS LOGIC LAYER (note/note.go)                     │
// │ - Note struct definition                                │
// │ - Validation rules                                      │
// │ - Display formatting                                    │
// │ - JSON serialization                                    │
// │ - File persistence                                      │
// └─────────────────────────────────────────────────────────┘
//
// DATA FLOW:
// ┌─────────────────────────────────────────────────────────┐
// │ 1. User Types Input                                     │
// │    "Shopping List" → "Buy milk, eggs, bread"            │
// ├─────────────────────────────────────────────────────────┤
// │ 2. getUserInput() Cleans Input                          │
// │    Removes \n, \r characters                            │
// ├─────────────────────────────────────────────────────────┤
// │ 3. note.New() Validates & Creates                       │
// │    Checks not empty, adds timestamp                     │
// ├─────────────────────────────────────────────────────────┤
// │ 4. Display() Shows Note                                 │
// │    User confirms correctness                            │
// ├─────────────────────────────────────────────────────────┤
// │ 5. Save() Persists to Disk                              │
// │    shopping_list.json created                           │
// ├─────────────────────────────────────────────────────────┤
// │ 6. Success Message                                      │
// │    "Saving the note succeeded!"                         │
// └─────────────────────────────────────────────────────────┘

// ==================== SAMPLE EXECUTION ====================
//
// Terminal Session:
// ═══════════════════════════════════════════════════════════
// $ go run main.go
// Note title: Project Ideas
// Note content: Build a task manager and note-taking app
// 
// Your note titled Project Ideas has the following content:
//
// Build a task manager and note-taking app
//
// Saving the note succeeded!
// ═══════════════════════════════════════════════════════════
//
// File Created: project_ideas.json
// ───────────────────────────────────────────────────────────
// {
//   "title": "Project Ideas",
//   "content": "Build a task manager and note-taking app",
//   "created_at": "2025-01-17T16:20:30.123456789Z"
// }
// ═══════════════════════════════════════════════════════════

// ==================== ERROR SCENARIOS ====================
//
// Scenario 1: Empty Title
// ───────────────────────────────────────────────────────────
// Note title: [Enter]
// Note content: Some content
// Invalid input.
// [Program exits]
// ───────────────────────────────────────────────────────────
//
// Scenario 2: Empty Content
// ───────────────────────────────────────────────────────────
// Note title: My Note
// Note content: [Enter]
// Invalid input.
// [Program exits]
// ───────────────────────────────────────────────────────────
//
// Scenario 3: Save Failure (disk full, permissions)
// ───────────────────────────────────────────────────────────
// Note title: Test
// Note content: Content
// 
// Your note titled Test has the following content:
//
// Content
//
// Saving the note failed.
// [Program exits - note displayed but not saved]
// ───────────────────────────────────────────────────────────

// ==================== KEY DESIGN PATTERNS ====================
//
// 1. SEPARATION OF CONCERNS
//    - main.go: UI and user interaction
//    - note.go: Business logic and data
//
// 2. SINGLE RESPONSIBILITY
//    - getUserInput(): Only handles input reading
//    - getNoteData(): Only orchestrates data collection
//    - note.New(): Only creates and validates
//    - note.Save(): Only handles persistence
//
// 3. ERROR HANDLING
//    - Explicit error checking at each step
//    - Early return on errors
//    - Clear error messages to user
//
// 4. ENCAPSULATION
//    - Note fields are exported (for JSON)
//    - But creation controlled via New()
//    - Validation centralized
//
// 5. FAIL FAST
//    - Validate immediately on creation
//    - Don't continue with invalid data
//    - Exit early with clear messages

// ==================== TECHNOLOGIES USED ====================
//
// Standard Library Packages:
// - bufio: Buffered I/O for efficient reading
// - fmt: Formatted printing
// - os: File operations and stdin access
// - strings: String manipulation
// - encoding/json: JSON serialization
// - time: Timestamps
// - errors: Error creation
//
// Go Features:
// - Structs and methods
// - Multiple return values
// - Error handling pattern
// - Packages and imports
// - Struct tags for JSON
// - Value vs pointer receivers

// ==================== PRODUCTION READINESS ====================
//
// ✅ Robust input handling (multi-word, cross-platform)
// ✅ Comprehensive error handling
// ✅ User feedback at every step
// ✅ Data validation
// ✅ File persistence with JSON
// ✅ Clean architecture
// ✅ Reusable functions
// ✅ Clear separation of concerns
//
// Potential Enhancements:
// - Load existing notes
// - List all saved notes
// - Update/delete notes
// - Search functionality
// - Categories/tags
// - Markdown support
// - Export to different formats

// ----------------
🎯 COMPLETE SYSTEM SUMMARY:
This is a fully functional, production-ready note-taking application demonstrating:
✅ Clean Architecture - Separation between UI (main) and logic (note package)
✅ Robust Input - Handles multi-word input on all platforms
✅ Validation - Ensures data integrity at creation
✅ Persistence - Saves to JSON files with proper formatting
✅ Error Handling - Explicit checks with user-friendly messages
✅ User Experience - Confirmation, feedback, clear prompts
This application teaches fundamental Go concepts:

Package organization
Struct methods
Error handling patterns
File I/O
JSON serialization
Cross-platform compatibility
