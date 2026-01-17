package main

import (
	"bufio"  // Buffered I/O for reading user input
	"fmt"    // Formatted I/O for printing
	"os"     // OS functionality for stdin access
	"strings" // String manipulation utilities
	
	// Import both custom packages
	"example.com/note/note" // Note type with Save() method
	"example.com/note/todo" // Todo type with Save() method
)

// ==================== SAVER INTERFACE ====================
// saver defines a contract for types that can be saved
//
// INTERFACE POWER:
// - Enables polymorphism (one function, many types)
// - Reduces code duplication
// - Makes code extensible (easy to add new types)
//
// ANY TYPE that has a Save() error method implements this interface
// - note.Note implements saver (has Save() error)
// - todo.Todo implements saver (has Save() error)
// - No explicit "implements" declaration needed (implicit in Go)
//
// USAGE:
// var s saver
// s = todo      // ✅ Works - todo.Todo has Save() error
// s = userNote  // ✅ Works - note.Note has Save() error
// s.Save()      // Calls the appropriate Save() method
type saver interface {
	Save() error // Method signature: no params, returns error
}

func main() {
	// ==================== INPUT COLLECTION ====================
	// Collect data for both Note and Todo
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	
	// ==================== CREATE TODO ====================
	// Create and validate Todo
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// ==================== CREATE NOTE ====================
	// Create and validate Note
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// ==================== DISPLAY AND SAVE TODO ====================
	// Display todo to user
	todo.Display()
	
	// ✨ INTERFACE IN ACTION ✨
	// Save todo using the polymorphic saveData() function
	// - saveData() accepts ANY type that implements saver interface
	// - todo.Todo has Save() error → implements saver
	// - Go automatically converts todo to saver interface type
	// - This is POLYMORPHISM: same function, different types
	err = saveData(todo)
	if err != nil {
		return // saveData already printed error message
	}
	
	// ==================== DISPLAY AND SAVE NOTE ====================
	// Display note to user
	userNote.Display()
	
	// ✨ INTERFACE IN ACTION AGAIN ✨
	// Save note using the SAME saveData() function
	// - saveData() works with note.Note too!
	// - note.Note has Save() error → implements saver
	// - Same function handles different type seamlessly
	// - NO CODE DUPLICATION
	err = saveData(userNote)
	if err != nil {
		return
	}
	
	// ==================== COMPARISON: BEFORE vs AFTER ====================
	//
	// BEFORE (without interface - duplicated code):
	// ───────────────────────────────────────────────────────
	// err = todo.Save()
	// if err != nil {
	//     fmt.Println("Saving the todo failed.")
	//     return
	// }
	// fmt.Println("Saving the todo succeeded!")
	//
	// err = userNote.Save()
	// if err != nil {
	//     fmt.Println("Saving the note failed.")
	//     return
	// }
	// fmt.Println("Saving the note succeeded!")
	// ───────────────────────────────────────────────────────
	//
	// AFTER (with interface - DRY principle):
	// ───────────────────────────────────────────────────────
	// err = saveData(todo)       // Reusable function
	// if err != nil { return }
	//
	// err = saveData(userNote)   // Same function!
	// if err != nil { return }
	// ───────────────────────────────────────────────────────
}

// ==================== POLYMORPHIC SAVE FUNCTION ====================
// saveData saves ANY type that implements the saver interface
//
// THIS IS THE POWER OF INTERFACES:
// - Single function works with multiple types
// - No need to know concrete type (Note vs Todo)
// - Only cares about behavior (has Save() method)
// - Reduces code duplication dramatically
//
// FUNCTION DESIGN:
// - Parameter: data saver (interface type, not concrete type)
// - Works with: note.Note, todo.Todo, or ANY future type with Save()
// - Returns: error (nil if successful, error if failed)
//
// HOW IT WORKS:
// 1. Receives value of type that implements saver
// 2. Calls Save() on that value
// 3. The ACTUAL type's Save() method is called (dynamic dispatch)
// 4. Handles success/failure uniformly
//
// POLYMORPHISM IN ACTION:
// saveData(todo)      → Calls todo.Todo's Save() method
// saveData(userNote)  → Calls note.Note's Save() method
// saveData(anything)  → Calls anything's Save() method (if implements saver)
//
// Parameters:
// - data: Any type implementing saver interface
//
// Returns:
// - error: nil if save successful, error object if failed
func saveData(data saver) error {
	// ==================== CALL SAVE METHOD ====================
	// Call Save() on the interface value
	//
	// DYNAMIC DISPATCH:
	// - Go automatically calls the correct Save() method
	// - If data is todo.Todo → calls todo.Todo.Save()
	// - If data is note.Note → calls note.Note.Save()
	// - This is determined at runtime, not compile time
	//
	// INTERFACE MECHANICS:
	// - Interface value contains: (type, value) pair
	// - Example: (todo.Todo, <todo data>)
	// - When calling Save(), Go looks up the method on the concrete type
	err := data.Save()
	
	// ==================== ERROR HANDLING ====================
	// Check if saving failed
	if err != nil {
		// Print generic error message
		// Note: Message says "note" but could be todo - minor bug
		// Better: "Saving failed." (type-agnostic)
		fmt.Println("Saving the note failed.")
		return err // Return error to caller
	}
	
	// ==================== SUCCESS MESSAGE ====================
	// Saving succeeded
	fmt.Println("Saving the note succeeded!")
	return nil // Return nil = no error
	
	// ==================== POTENTIAL IMPROVEMENT ====================
	// Could make messages type-aware:
	//
	// switch v := data.(type) {
	// case note.Note:
	//     fmt.Println("Saving the note succeeded!")
	// case todo.Todo:
	//     fmt.Println("Saving the todo succeeded!")
	// default:
	//     fmt.Println("Saving succeeded!")
	// }
}

// ==================== INPUT ORCHESTRATION ====================
// getNoteData collects title and content for a Note
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

// ==================== INTERFACE CONCEPTS EXPLAINED ====================
//
// 1. INTERFACE TYPE vs CONCRETE TYPE
// ───────────────────────────────────────────────────────────
// Concrete type: note.Note, todo.Todo (actual structs)
// Interface type: saver (abstract contract)
//
// Concrete types define data + behavior
// Interface types define ONLY behavior (method signatures)
//
// 2. IMPLICIT IMPLEMENTATION
// ───────────────────────────────────────────────────────────
// Go uses "duck typing": if it walks like a duck and quacks like a duck...
//
// If a type has Save() error method → implements saver interface
// No need to declare "Note implements saver"
// Compiler checks automatically
//
// 3. INTERFACE VALUES
// ───────────────────────────────────────────────────────────
// An interface value holds: (concrete type, concrete value)
//
// var s saver = todo
// s contains: (todo.Todo, <the todo data>)
//
// When you call s.Save():
// - Go looks at concrete type (todo.Todo)
// - Calls that type's Save() method
// - This is "dynamic dispatch"
//
// 4. WHY INTERFACES?
// ───────────────────────────────────────────────────────────
// ✅ Polymorphism - treat different types uniformly
// ✅ Abstraction - depend on behavior, not implementation
// ✅ Decoupling - change implementations without changing callers
// ✅ Testing - easy to create mock implementations
// ✅ Extensibility - add new types without changing existing code

// ==================== INTERFACE IMPLEMENTATION CHECK ====================
//
// How to verify a type implements an interface:
//
// COMPILE-TIME CHECK (recommended):
// var _ saver = note.Note{}  // Compile error if doesn't implement
// var _ saver = todo.Todo{}  // Compile error if doesn't implement
//
// RUNTIME CHECK:
// var s saver = todo
// if _, ok := s.(todo.Todo); ok {
//     fmt.Println("It's a Todo!")
// }

// ==================== ADDING NEW TYPES ====================
//
// To add a new savable type (e.g., Event):
//
// 1. Create event package with Event type
// 2. Add Save() error method to Event
// 3. That's it! Event automatically implements saver
// 4. Can use saveData(event) immediately
// 5. No changes to saveData() function needed
//
// Example:
// type Event struct { ... }
// func (e Event) Save() error { ... }
//
// // Now works automatically:
// event, _ := event.New("Meeting")
// saveData(event)  // ✅ Works!

// ==================== EXECUTION FLOW ====================
//
// Call: saveData(todo)
// ┌─────────────────────────────────────────────────────┐
// │ 1. todo (concrete type: todo.Todo) passed          │
// ├─────────────────────────────────────────────────────┤
// │ 2. Go converts to interface: saver                 │
// │    Interface value: (todo.Todo, <todo data>)       │
// ├─────────────────────────────────────────────────────┤
// │ 3. saveData receives interface value               │
// ├─────────────────────────────────────────────────────┤
// │ 4. data.Save() called                              │
// │    Go looks up Save() on concrete type (todo.Todo) │
// ├─────────────────────────────────────────────────────┤
// │ 5. todo.Todo.Save() executes                       │
// │    Creates JSON file for todo                      │
// ├─────────────────────────────────────────────────────┤
// │ 6. Returns error (or nil)                          │
// ├─────────────────────────────────────────────────────┤
// │ 7. saveData checks error and prints message        │
// └─────────────────────────────────────────────────────┘
//
// Same flow for: saveData(userNote)
// But step 5 calls note.Note.Save() instead

// ==================== BENEFITS ACHIEVED ====================
//
// BEFORE REFACTORING:
// ❌ Duplicate save logic (8 lines repeated)
// ❌ Hard to add new types
// ❌ Changes need to be made in multiple places
//
// AFTER REFACTORING:
// ✅ Single saveData() function (5 lines, reusable)
// ✅ Works with any type that has Save()
// ✅ Easy to add new types (just implement Save())
// ✅ DRY principle (Don't Repeat Yourself)
// ✅ Testable (can mock saver interface)

// --------------------
🔑 KEY CONCEPTS:
1. Interface as Contract
go// Interface defines behavior
type saver interface {
    Save() error
}

// Types that have this method automatically implement it
// note.Note has Save() error → implements saver ✅
// todo.Todo has Save() error → implements saver ✅
2. Polymorphism in Action
gofunc saveData(data saver) error {
    return data.Save() // Works with ANY type that implements saver
}

// Same function, different types:
saveData(todo)      // Calls todo.Todo.Save()
saveData(userNote)  // Calls note.Note.Save()
3. Before vs After
Before (duplication):
go// Todo save
err = todo.Save()
if err != nil {
    fmt.Println("Saving the todo failed.")
    return
}
fmt.Println("Saving the todo succeeded!")

// Note save (exact same pattern!)
err = userNote.Save()
if err != nil {
    fmt.Println("Saving the note failed.")
    return
}
fmt.Println("Saving the note succeeded!")
After (DRY with interface):
goerr = saveData(todo)
if err != nil { return }

err = saveData(userNote)
if err != nil { return }

📊 Interface Value Internals:
govar s saver = todo

// s contains:
// ┌──────────────────────┐
// │ Type: todo.Todo      │
// │ Value: <todo data>   │
// └──────────────────────┘

// When calling s.Save():
// 1. Go looks at Type field → todo.Todo
// 2. Finds Save() method on todo.Todo
// 3. Calls it with Value field

🎯 This Demonstrates:
✅ Interface definition - type saver interface
✅ Implicit implementation - No "implements" keyword
✅ Polymorphism - One function, multiple types
✅ Code reuse - No duplication
✅ Extensibility - Easy to add new types
