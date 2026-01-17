package main

import (
	"bufio"  // Buffered I/O for reading user input
	"fmt"    // Formatted I/O for printing
	"os"     // OS functionality for stdin access
	"strings" // String manipulation utilities
	
	// Import custom note package
	// Provides: Note type with New(), Display(), Save()
	"example.com/note/note"
	
	// ==================== NEW IMPORT: TODO PACKAGE ====================
	// Import custom todo package
	// Provides: Todo type with New(), Display(), Save()
	// Structure similar to note package but for todo items
	"example.com/note/todo"
)

// ==================== INTERFACE DEFINITION ====================
// saver is an interface that defines a contract for types that can be saved
//
// INTERFACE SYNTAX:
// type InterfaceName interface {
//     MethodName() ReturnType
// }
//
// WHAT IS AN INTERFACE?
// - A contract that types must fulfill
// - Defines behavior (methods) without implementation
// - Any type that has these methods "implements" the interface
// - Go uses implicit interface implementation (no "implements" keyword)
//
// WHY USE INTERFACES?
// ✅ Polymorphism - treat different types uniformly
// ✅ Abstraction - depend on behavior, not concrete types
// ✅ Flexibility - easy to add new types
// ✅ Testability - can mock implementations
//
// SAVER INTERFACE:
// - Requires one method: Save() error
// - Any type with Save() error automatically implements saver
// - Both Note and Todo implement this (they both have Save() error)
//
// CURRENT STATUS:
// - Interface defined but NOT USED YET in this code
// - saveData() function exists but is empty
// - This is a work in progress / teaching example
type saver interface {
	Save() error // Method signature: no parameters, returns error
}

// ==================== MAIN ENTRY POINT ====================
// Extended application that now handles both Notes and Todos
//
// NEW FEATURES:
// ✅ Creates both Note and Todo objects
// ✅ Displays both types
// ✅ Saves both types to separate files
// ⚠️  Has interface defined but not used yet (teaching in progress)
func main() {
	// ==================== COLLECT NOTE DATA ====================
	// Get title and content for a Note
	// Returns: (title string, content string)
	title, content := getNoteData()
	
	// ==================== COLLECT TODO DATA ====================
	// Get text for a Todo item
	// Returns: todo text string
	// Note: Todo only needs one field (text), unlike Note (title + content)
	todoText := getUserInput("Todo text: ")
	
	// ==================== CREATE TODO ====================
	// Create Todo using todo package's constructor
	//
	// todo.New(text string) (Todo, error)
	// - Similar pattern to note.New()
	// - Validates text is not empty
	// - Auto-sets creation timestamp
	// - Returns (Todo, error)
	todo, err := todo.New(todoText)
	
	// Check if Todo creation failed
	if err != nil {
		fmt.Println(err) // Display error (e.g., "Invalid input.")
		return           // Exit early
	}
	
	// ==================== CREATE NOTE ====================
	// Create Note using note package's constructor
	// Same pattern as before
	userNote, err := note.New(title, content)
	
	// Check if Note creation failed
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// ==================== DISPLAY TODO ====================
	// Show the created todo to user
	// todo.Display() prints todo information
	todo.Display()
	
	// ==================== SAVE TODO ====================
	// Persist todo to disk
	// todo.Save() creates a JSON file (similar to note.Save())
	err = todo.Save()
	
	// Check if saving todo failed
	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}
	
	// Confirm todo saved successfully
	fmt.Println("Saving the todo succeeded!")
	
	// ==================== DISPLAY NOTE ====================
	// Show the created note to user
	userNote.Display()
	
	// ==================== SAVE NOTE ====================
	// Persist note to disk
	err = userNote.Save()
	
	// Check if saving note failed
	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
	
	// Confirm note saved successfully
	fmt.Println("Saving the note succeeded!")
	
	// ==================== OBSERVATION: CODE DUPLICATION ====================
	// Notice the repetition:
	// 1. todo.Display() → todo.Save() → check error → print success
	// 2. userNote.Display() → userNote.Save() → check error → print success
	//
	// This is where the 'saver' interface would be useful!
	// We could create a function that works with ANY type that can Save()
}

// ==================== UNUSED FUNCTION (TEACHING PLACEHOLDER) ====================
// saveData is currently unused and incomplete
//
// CURRENT PROBLEM:
// - Takes only note.Note parameter
// - Cannot handle todo.Todo
// - Type-specific, not polymorphic
//
// FUTURE WITH INTERFACE:
// Could be rewritten as:
// func saveData(data saver) error {
//     err := data.Save()
//     if err != nil {
//         return fmt.Errorf("saving failed: %w", err)
//     }
//     fmt.Println("Saving succeeded!")
//     return nil
// }
//
// Then could call:
// saveData(todo)     // Works! Todo implements saver
// saveData(userNote) // Works! Note implements saver
//
// POLYMORPHISM:
// - Same function works with different types
// - Types just need to implement the interface (have Save() method)
func saveData(data note.Note) {
	// Currently empty - teaching placeholder
	// Will be implemented later to demonstrate interface usage
}

// ==================== INPUT ORCHESTRATION ====================
// getNoteData collects both fields needed for a Note
// No changes from previous version
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

// ==================== ROBUST INPUT FUNCTION ====================
// getUserInput reads complete line of user input
// No changes from previous version
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

// ==================== TODO PACKAGE (todo/todo.go) ====================
// Here's what the todo package likely contains:
//
// package todo
//
// import (
//     "encoding/json"
//     "errors"
//     "fmt"
//     "os"
//     "strings"
//     "time"
// )
//
// type Todo struct {
//     Text      string    `json:"text"`
//     CreatedAt time.Time `json:"created_at"`
// }
//
// func (t Todo) Display() {
//     fmt.Printf("Todo: %v\n", t.Text)
// }
//
// func (t Todo) Save() error {
//     fileName := strings.ReplaceAll(t.Text, " ", "_")
//     fileName = strings.ToLower(fileName) + ".json"
//     json, err := json.Marshal(t)
//     if err != nil {
//         return err
//     }
//     return os.WriteFile(fileName, json, 0644)
// }
//
// func New(text string) (Todo, error) {
//     if text == "" {
//         return Todo{}, errors.New("Invalid input.")
//     }
//     return Todo{
//         Text:      text,
//         CreatedAt: time.Now(),
//     }, nil
// }

// ==================== INTERFACE IMPLEMENTATION ====================
//
// IMPLICIT INTERFACE IMPLEMENTATION:
// In Go, types implement interfaces automatically if they have the required methods
//
// The saver interface requires:
// - Save() error method
//
// note.Note has Save() error ✅ → implements saver
// todo.Todo has Save() error ✅ → implements saver
//
// No explicit declaration needed!
// Unlike other languages: no "class Note implements Saver"
//
// USAGE EXAMPLE (not in current code):
// var s saver
// s = todo           // ✅ Works - todo.Todo implements saver
// s = userNote       // ✅ Works - note.Note implements saver
// err := s.Save()    // Calls the appropriate Save() method
//
// POLYMORPHISM EXAMPLE:
// func saveAll(items []saver) {
//     for _, item := range items {
//         err := item.Save()
//         if err != nil {
//             fmt.Println("Save failed:", err)
//         }
//     }
// }
//
// saveAll([]saver{todo, userNote}) // Both types work!

// ==================== CODE ORGANIZATION ====================
//
// PROJECT STRUCTURE:
// project/
// ├── go.mod
// ├── main.go (this file)
// ├── note/
// │   └── note.go (Note type)
// └── todo/
//     └── todo.go (Todo type)
//
// PACKAGE RESPONSIBILITIES:
// main: UI, user interaction, orchestration
// note: Note type, validation, persistence
// todo: Todo type, validation, persistence

// ==================== SAMPLE EXECUTION ====================
//
// Terminal Session:
// ═══════════════════════════════════════════════════════════
// $ go run main.go
// Note title: Shopping List
// Note content: Buy milk and eggs
// Todo text: Call dentist
// 
// Todo: Call dentist
// Saving the todo succeeded!
// 
// Your note titled Shopping List has the following content:
//
// Buy milk and eggs
//
// Saving the note succeeded!
// ═══════════════════════════════════════════════════════════
//
// Files Created:
// 1. call_dentist.json
//    {"text":"Call dentist","created_at":"2025-01-17T10:30:00Z"}
//
// 2. shopping_list.json
//    {"title":"Shopping List","content":"Buy milk and eggs","created_at":"2025-01-17T10:30:00Z"}

// ==================== PROBLEMS IN CURRENT CODE ====================
//
// 1. CODE DUPLICATION
//    - Display → Save → Check error → Print success (repeated twice)
//    - Could use a helper function with interface
//
// 2. INTERFACE NOT USED
//    - saver interface defined but not utilized
//    - saveData() function exists but empty
//
// 3. NO ABSTRACTION
//    - Handling each type separately
//    - Not leveraging polymorphism
//
// SOLUTION: Use the saver interface!
// func processAndSave(s saver, displayFunc func()) error {
//     displayFunc()
//     err := s.Save()
//     if err != nil {
//         return err
//     }
//     fmt.Println("Saving succeeded!")
//     return nil
// }

// ==================== INTERFACE BENEFITS (TO BE DEMONSTRATED) ====================
//
// WITHOUT INTERFACE (current code):
// - Separate handling for each type
// - Duplicate save logic
// - Hard to add new types
//
// WITH INTERFACE (future improvement):
// - Single function handles all savable types
// - No duplication
// - Easy to add new types (just implement Save())
//
// EXAMPLE:
// items := []saver{todo, userNote}
// for _, item := range items {
//     if err := item.Save(); err != nil {
//         fmt.Println("Failed:", err)
//     } else {
//         fmt.Println("Succeeded!")
//     }
// }


// -------------
🔑 KEY CONCEPTS:
1. Interface Definition
go// Define interface
type saver interface {
    Save() error
}

// Types automatically implement if they have the method
// note.Note has Save() error → implements saver ✅
// todo.Todo has Save() error → implements saver ✅
2. Implicit Implementation
go// No explicit declaration needed!

// Java/C# style (NOT Go):
// class Note implements Saver { ... }

// Go style:
// Just have the method, that's it!
type Note struct { ... }
func (n Note) Save() error { ... }
// Note now implements saver automatically
3. Current State vs Future State
Now (code duplication):
go// Todo handling
todo.Display()
err = todo.Save()
if err != nil { /* ... */ }
fmt.Println("Saving the todo succeeded!")

// Note handling (same pattern, duplicated)
userNote.Display()
err = userNote.Save()
if err != nil { /* ... */ }
fmt.Println("Saving the note succeeded!")
Future (with interface):
gofunc saveItem(item saver, name string) {
    err := item.Save()
    if err != nil {
        fmt.Printf("Saving the %s failed.\n", name)
        return
    }
    fmt.Printf("Saving the %s succeeded!\n", name)
}

// Use it
saveItem(todo, "todo")
saveItem(userNote, "note")
```

---

**📊 Sample Output:**
```
Note title: Shopping List
Note content: Buy milk and eggs
Todo text: Call dentist

Todo: Call dentist
Saving the todo succeeded!

Your note titled Shopping List has the following content:

Buy milk and eggs

Saving the note succeeded!
Files created:

call_dentist.json
shopping_list.json


🎯 This Code Demonstrates:
✅ Multiple packages - note and todo
✅ Interface definition - saver interface
✅ Implicit implementation - both types implement saver
⚠️ Interface not yet used - teaching in progress
⚠️ Code duplication - opportunity for refactoring
