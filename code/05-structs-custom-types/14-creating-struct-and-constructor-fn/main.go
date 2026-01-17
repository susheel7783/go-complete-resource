package main

import (
	"fmt" // Standard library for formatted I/O
	
	// Import the custom note package
	// This package provides the Note type and New() constructor
	"example.com/note/note"
)

func main() {
	// ==================== COLLECT USER INPUT ====================
	// Get note title and content from the user
	// getNoteData() returns two strings: (title, content)
	// No error handling at this stage - errors handled later in note.New()
	title, content := getNoteData()
	
	// ==================== CREATE NOTE WITH VALIDATION ====================
	// Attempt to create a Note using the collected data
	// 
	// note.New() performs validation and returns:
	// - Note: the created note (or empty Note{} if validation fails)
	// - error: nil if successful, error message if validation fails
	//
	// Validation happens HERE, not during input collection
	// This separates input collection from validation logic
	userNote, err := note.New(title, content)
	
	// ==================== ERROR HANDLING ====================
	// Check if Note creation failed
	// Common reasons for failure:
	// - Empty title (user just pressed Enter)
	// - Empty content (user just pressed Enter)
	if err != nil {
		// Display the error message to the user
		fmt.Println(err) // Output: "Invalid input."
		
		// Exit the program early - no point continuing with invalid note
		return // Ends the main function, program terminates
	}
	
	// ==================== MISSING: WHAT TO DO WITH THE NOTE? ====================
	// If we reach here, userNote is valid and ready to use
	// BUT there's no code to do anything with it!
	//
	// PROBLEM: userNote is created but never used
	// - Not displayed to the user
	// - Not saved to a file
	// - Not sent anywhere
	//
	// The Go compiler will warn: "userNote declared and not used"
	//
	// What SHOULD be here:
	// userNote.Display()                    // If Note has Display() method
	// saveToFile(userNote)                  // Save to file
	// fmt.Println("Note created successfully!") // At least confirm success
	
	// As it stands, this program:
	// 1. Collects input ✅
	// 2. Creates and validates note ✅
	// 3. Does nothing with it ❌
	// 4. Exits silently ❌
	
	_ = userNote // Suppress "unused variable" compiler warning (temporary)
}

// ==================== INPUT COLLECTION ORCHESTRATOR ====================
// getNoteData collects both title and content from the user
//
// DESIGN PATTERN: Input collection without validation
// - Just gathers raw input
// - No error checking
// - Validation deferred to note.New()
//
// Returns:
// - string: note title (may be empty)
// - string: note content (may be empty)
//
// Note: No error return value - assumes getUserInput() always succeeds
func getNoteData() (string, string) {
	// Collect note title
	// getUserInput() returns whatever the user types (including empty string)
	title := getUserInput("Note title:")
	
	// Collect note content
	content := getUserInput("Note content:")
	
	// Return both values
	// These might be empty strings if user just pressed Enter
	// Validation happens later in note.New(), not here
	return title, content
}

// ==================== SIMPLE INPUT HELPER ====================
// getUserInput displays a prompt and reads user input
//
// LIMITATIONS OF THIS APPROACH:
// ❌ Only reads SINGLE WORD (stops at whitespace)
// ❌ "My First Note" becomes just "My"
// ❌ No error handling
// ❌ No validation
//
// This is a MINIMAL implementation suitable for:
// ✅ Single-word inputs
// ✅ Learning/prototyping
// ✅ Simple use cases
//
// For production, use bufio.Reader (as shown in previous examples)
//
// Parameters:
// - prompt (string): message to display to user
//
// Returns:
// - string: user's input (or empty string if user pressed Enter)
//           Only captures first word if user types multiple words
func getUserInput(prompt string) string {
	// ==================== DISPLAY PROMPT ====================
	// Print the prompt without a newline
	// User will type on the same line
	fmt.Print(prompt)
	
	// ==================== DECLARE INPUT VARIABLE ====================
	// Variable to store the user's input
	var value string
	
	// ==================== READ INPUT ====================
	// fmt.Scanln() reads from standard input (keyboard)
	// - Reads until newline (Enter key)
	// - BUT stops at FIRST WHITESPACE
	// - Stores result in value via pointer (&value)
	//
	// BEHAVIOR:
	// User types: "Hello" → value = "Hello" ✅
	// User types: "Hello World" → value = "Hello" ❌ (loses "World")
	// User types: "" (just Enter) → value = "" ✅
	//
	// No error checking - assumes Scanln always succeeds
	// In reality, Scanln returns (count, error) but we ignore them
	fmt.Scanln(&value)
	
	// ==================== RETURN INPUT ====================
	// Return whatever was read (might be empty string)
	// No validation or trimming performed
	return value
}

// ==================== PROGRAM FLOW ====================
//
// Scenario 1: Valid single-word input
// ┌─────────────────────────────────────────────┐
// │ User Input:                                 │
// │   Note title: Shopping                      │
// │   Note content: Groceries                   │
// ├─────────────────────────────────────────────┤
// │ Program Flow:                               │
// │ getNoteData() → ("Shopping", "Groceries")   │
// │ note.New() → (Note{...}, nil) ✅            │
// │ err == nil → continue                       │
// │ ...nothing happens (userNote unused) ❌     │
// │ Program exits silently                      │
// └─────────────────────────────────────────────┘
//
// Scenario 2: Multi-word input (PROBLEM!)
// ┌─────────────────────────────────────────────┐
// │ User Input:                                 │
// │   Note title: My First Note                 │
// │   Note content: This is content             │
// ├─────────────────────────────────────────────┤
// │ Program Flow:                               │
// │ getNoteData() → ("My", "This") ❌           │
// │   (Loses "First Note" and "is content")     │
// │ note.New("My", "This") → (Note{...}, nil)   │
// │ Creates note with incomplete data ❌        │
// └─────────────────────────────────────────────┘
//
// Scenario 3: Empty input
// ┌─────────────────────────────────────────────┐
// │ User Input:                                 │
// │   Note title: [Enter]                       │
// │   Note content: Content                     │
// ├─────────────────────────────────────────────┤
// │ Program Flow:                               │
// │ getNoteData() → ("", "Content")             │
// │ note.New() → (Note{}, error) ❌             │
// │ err != nil → print error and exit           │
// │ Output: Invalid input.                      │
// └─────────────────────────────────────────────┘

// ==================== ISSUES WITH THIS CODE ====================
//
// 1. UNUSED VARIABLE
//    - userNote is created but never used
//    - Compiler warning: "userNote declared and not used"
//    - User gets no feedback that note was created
//
// 2. MULTI-WORD INPUT BROKEN
//    - fmt.Scanln only reads first word
//    - "My Shopping List" becomes "My"
//    - Data loss without user knowing
//
// 3. NO SUCCESS FEEDBACK
//    - If note creation succeeds, program exits silently
//    - User doesn't know if it worked
//    - No confirmation message
//
// 4. NO DISPLAY FUNCTIONALITY
//    - Note is created but not shown
//    - Can't verify the note was created correctly
//    - Missing the whole point of creating a note
//
// 5. NO PERSISTENCE
//    - Note exists only in memory
//    - Lost when program exits
//    - No way to save or retrieve notes

// ==================== HOW TO FIX ====================
//
// Fix 1: Add display functionality
// if err != nil {
//     fmt.Println(err)
//     return
// }
// userNote.Display()  // Requires Display() method in note package
//
// Fix 2: Use bufio.Reader for multi-word input
// func getUserInput(prompt string) string {
//     fmt.Print(prompt)
//     reader := bufio.NewReader(os.Stdin)
//     text, _ := reader.ReadString('\n')
//     return strings.TrimSpace(text)
// }
//
// Fix 3: Add success confirmation
// if err != nil {
//     fmt.Println(err)
//     return
// }
// fmt.Println("✅ Note created successfully!")
// userNote.Display()
//
// Fix 4: Add file persistence
// err = userNote.Save("notes.txt")
// if err != nil {
//     fmt.Println("Failed to save:", err)
// }


// --------------------------
🔑 CRITICAL ISSUES:
1. Multi-Word Input Problem
go// Current code:
fmt.Scanln(&value)  // Stops at whitespace!

// User types: "My Shopping List"
// Result: value = "My"  ❌ Lost "Shopping List"

// Fix: Use bufio.Reader
reader := bufio.NewReader(os.Stdin)
text, _ := reader.ReadString('\n')
value = strings.TrimSpace(text)  // Gets full line ✅
2. Unused Variable
gouserNote, err := note.New(title, content)
if err != nil {
    fmt.Println(err)
    return
}
// userNote created but never used! ❌

// Fix: Add display or save
userNote.Display()  // Show the note
// OR
saveNote(userNote)  // Save to file
3. No User Feedback
go// Current: Silent success
if err != nil {
    fmt.Println(err)  // Only shows errors
    return
}
// ...nothing... ❌ User confused

// Fix: Show success message
if err != nil {
    fmt.Println("❌ Error:", err)
    return
}
fmt.Println("✅ Note created successfully!")
userNote.Display()
```

---

**📊 What Happens (Step-by-Step):**

**Test 1: Single-word input (Works)**
```
Note title:Shopping
Note content:Milk
(Program exits silently - note created but unused)
```

**Test 2: Multi-word input (Broken)**
```
Note title:My Shopping List
Note content:Buy milk and eggs
(Only captures "My" and "Buy" - rest is lost!)
```

**Test 3: Empty input (Error handled)**
```
Note title:
Note content:Something
Invalid input.
(Program exits with error message)

💡 Complete Fixed Version:
gopackage main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/note/note"
)

func main() {
	title, content := getNoteData()
	
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	
	// FIXED: Actually use the note!
	fmt.Println("\n✅ Note created successfully!")
	fmt.Println("\nYour Note:")
	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Content: %s\n", content)
	
	// Or if note package has Display():
	// userNote.Display()
	
	_ = userNote  // Still need this if note package has no methods
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

// FIXED: Handle multi-word input
func getUserInput(prompt string) string {
	fmt.Printf("%s ", prompt)
	
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	
	return strings.TrimSpace(text)
}

🚀 Summary of Issues:
IssueCurrent CodeImpactFixMulti-word inputfmt.ScanlnData lossUse bufio.ReaderUnused noteNo display/saveWasted effortAdd Display() or saveNo feedbackSilent successUser confusionPrint success messageNo persistenceMemory onlyLost on exitAdd file saving
