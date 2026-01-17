package note

import (
	"encoding/json" // Package for JSON encoding/decoding (serialization)
	"errors"        // Package for creating error values
	"fmt"           // Package for formatted I/O (printing)
	"os"            // Package for operating system functionality (file operations)
	"strings"       // Package for string manipulation
	"time"          // Package for working with dates and times
)

// ==================== NOTE STRUCT DEFINITION ====================
// Note represents a note with title, content, and creation timestamp
//
// CRITICAL CHANGE: Fields are now EXPORTED (capitalized)
// Previously: title, content, createdAt (unexported)
// Now: Title, Content, CreatedAt (exported)
//
// WHY THE CHANGE?
// ✅ JSON marshaling requires EXPORTED fields
// ✅ json.Marshal() can only serialize exported (public) fields
// ✅ Unexported fields would be IGNORED during JSON conversion
//
// TRADE-OFF:
// ❌ Lost encapsulation - external code can now modify fields directly
// ✅ Gained JSON serialization capability
// ✅ Can save/load notes from files
//
// JSON field names will match struct field names:
// {"Title":"...", "Content":"...", "CreatedAt":"..."}
type Note struct {
	Title     string    // Exported: the note's title (was private before)
	Content   string    // Exported: the note's content (was private before)
	CreatedAt time.Time // Exported: timestamp when created (was private before)
}

// ==================== DISPLAY METHOD ====================
// Display shows the note's information in a formatted way
//
// METHOD RECEIVER: (note Note) - VALUE RECEIVER
// - Receives a COPY of the Note
// - Cannot modify the original Note
// - OK because we're only reading data
//
// Note: Now accessing exported fields (note.Title vs note.title)
func (note Note) Display() {
	// Print formatted note information
	// %v is a generic format verb that works with any type
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", 
		note.Title,   // Exported field (capitalized)
		note.Content) // Exported field (capitalized)
}

// ==================== SAVE METHOD (NEW!) ====================
// Save writes the note to a JSON file
//
// This method demonstrates:
// - File I/O operations
// - JSON serialization
// - String manipulation for filename generation
// - Error handling and propagation
//
// File naming strategy:
// - Title "My Shopping List" → filename "my_shopping_list.json"
// - Spaces replaced with underscores
// - Converted to lowercase
// - .json extension added
//
// Returns:
// - error: nil if save successful, error object if failed
func (note Note) Save() error {
	// ==================== GENERATE FILENAME ====================
	// Create a safe filename from the note title
	
	// Step 1: Replace all spaces with underscores
	// strings.ReplaceAll(source, old, new) replaces ALL occurrences
	// Example: "My Shopping List" → "My_Shopping_List"
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	
	// Step 2: Convert to lowercase for consistency
	// Example: "My_Shopping_List" → "my_shopping_list"
	fileName = strings.ToLower(fileName)
	
	// Step 3: Add .json extension
	// Example: "my_shopping_list" → "my_shopping_list.json"
	fileName = fileName + ".json"
	
	// Alternative one-liner (same result):
	// fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_")) + ".json"
	
	// ==================== SERIALIZE TO JSON ====================
	// Convert the Note struct to JSON format
	//
	// json.Marshal() converts Go struct to JSON bytes
	// - Takes any Go value (interface{})
	// - Returns []byte (JSON data) and error
	// - Only serializes EXPORTED fields (Title, Content, CreatedAt)
	// - Unexported fields would be IGNORED
	//
	// Example output:
	// {
	//   "Title": "Shopping List",
	//   "Content": "Buy milk and eggs",
	//   "CreatedAt": "2025-01-17T10:30:00Z"
	// }
	//
	// Variable name is 'json' (lowercase) - local variable, not the package
	json, err := json.Marshal(note)
	
	// Check if marshaling (serialization) failed
	if err != nil {
		// Propagate the error to the caller
		// Common errors: circular references, unsupported types
		// For this simple struct, should never fail
		return err
	}
	
	// ==================== WRITE TO FILE ====================
	// Write the JSON data to a file on disk
	//
	// os.WriteFile() is a convenience function that:
	// 1. Creates the file (or overwrites if exists)
	// 2. Writes the data
	// 3. Closes the file
	// 4. All in one function call
	//
	// Parameters:
	// - fileName: name of file to create (e.g., "my_note.json")
	// - json: the data to write ([]byte from json.Marshal)
	// - 0644: file permissions (owner: read+write, others: read-only)
	//
	// File permissions breakdown (Unix/Linux):
	// 0644 in octal = 110 100 100 in binary
	// - 6 (110): Owner can read(4) + write(2)
	// - 4 (100): Group can read(4)
	// - 4 (100): Others can read(4)
	//
	// Returns error if:
	// - Cannot create file (permission denied, disk full)
	// - Cannot write data (disk error, no space)
	// - Path is invalid
	return os.WriteFile(fileName, json, 0644)
	
	// Note: WriteFile returns error directly, so we return it
	// If successful, returns nil
	// Caller should check: if err := note.Save(); err != nil { ... }
}

// ==================== CONSTRUCTOR FUNCTION ====================
// New creates and returns a new Note with validation
//
// No changes from previous version - validation logic remains the same
//
// Parameters:
// - title: the note's title (cannot be empty)
// - content: the note's content (cannot be empty)
//
// Returns:
// - Note: the created Note (or empty Note{} if validation fails)
// - error: nil if successful, error object if validation fails
func New(title, content string) (Note, error) {
	// ==================== VALIDATION ====================
	// Validate that both required fields are non-empty
	if title == "" || content == "" {
		// Return empty Note and error
		return Note{}, errors.New("Invalid input.")
	}
	
	// ==================== NOTE CREATION ====================
	// Create and return the Note with current timestamp
	// Now using EXPORTED field names (Title vs title)
	return Note{
		Title:     title,      // Exported field
		Content:   content,    // Exported field
		CreatedAt: time.Now(), // Exported field - auto-set to current time
	}, nil
}

// ==================== JSON SERIALIZATION EXAMPLE ====================
//
// Given this Note:
// note := Note{
//     Title:     "Shopping List",
//     Content:   "Buy milk, eggs, bread",
//     CreatedAt: time.Date(2025, 1, 17, 10, 30, 0, 0, time.UTC),
// }
//
// json.Marshal(note) produces:
// {
//   "Title": "Shopping List",
//   "Content": "Buy milk, eggs, bread",
//   "CreatedAt": "2025-01-17T10:30:00Z"
// }
//
// This JSON is written to file: "shopping_list.json"

// ==================== FILE OPERATIONS FLOW ====================
//
// When note.Save() is called:
//
// 1. FILENAME GENERATION
//    Title: "My First Note"
//    → Replace spaces: "My_First_Note"
//    → Lowercase: "my_first_note"
//    → Add extension: "my_first_note.json"
//
// 2. JSON SERIALIZATION
//    Note struct → JSON bytes
//    {Title: "My First Note", ...} → {"Title":"My First Note",...}
//
// 3. FILE WRITE
//    Create/overwrite file "my_first_note.json"
//    Write JSON bytes
//    Set permissions to 0644
//
// 4. ERROR HANDLING
//    If any step fails, return error
//    If all succeed, return nil

// ==================== USAGE EXAMPLE ====================
//
// package main
//
// import (
//     "fmt"
//     "example.com/note/note"
// )
//
// func main() {
//     // Create a note
//     myNote, err := note.New("Shopping List", "Buy milk and eggs")
//     if err != nil {
//         fmt.Println("Error creating note:", err)
//         return
//     }
//
//     // Display it
//     myNote.Display()
//     // Output:
//     // Your note titled Shopping List has the following content:
//     //
//     // Buy milk and eggs
//
//     // Save to file
//     err = myNote.Save()
//     if err != nil {
//         fmt.Println("Error saving note:", err)
//         return
//     }
//
//     fmt.Println("Note saved to shopping_list.json")
//
//     // The file "shopping_list.json" now contains:
//     // {
//     //   "Title": "Shopping List",
//     //   "Content": "Buy milk and eggs",
//     //   "CreatedAt": "2025-01-17T14:23:45.123456789Z"
//     // }
// }

// ==================== TRADE-OFFS OF EXPORTED FIELDS ====================
//
// BEFORE (Unexported):
// type Note struct {
//     title     string  // ❌ json.Marshal ignores
//     content   string  // ❌ json.Marshal ignores
//     createdAt time.Time
// }
// Pros: Encapsulation, data protection
// Cons: Cannot serialize to JSON
//
// AFTER (Exported):
// type Note struct {
//     Title     string  // ✅ json.Marshal includes
//     Content   string  // ✅ json.Marshal includes
//     CreatedAt time.Time
// }
// Pros: Can serialize to JSON, can save to files
// Cons: Lost encapsulation, fields can be modified directly
//
// IMPACT ON EXTERNAL CODE:
// // Now possible (but not recommended):
// note.Title = ""  // ❌ Breaks validation rules!
// note.CreatedAt = pastTime  // ❌ Corrupts timestamp!
//
// BEST PRACTICE:
// Even though fields are exported, external code should:
// - Create notes only via New() (enforces validation)
// - Not modify fields directly (respect the intent)
// - Use provided methods (Display, Save)

// ==================== POTENTIAL IMPROVEMENTS ====================
//
// 1. Custom JSON field names (lowercase):
// type Note struct {
//     Title     string    `json:"title"`
//     Content   string    `json:"content"`
//     CreatedAt time.Time `json:"created_at"`
// }
//
// 2. Omit empty fields:
// type Note struct {
//     Title     string    `json:"title"`
//     Content   string    `json:"content,omitempty"`
//     CreatedAt time.Time `json:"created_at"`
// }
//
// 3. Pretty-print JSON:
// json, err := json.MarshalIndent(note, "", "  ")
//
// 4. Add Load() function to read from file:
// func Load(filename string) (Note, error) {
//     data, err := os.ReadFile(filename)
//     if err != nil {
//         return Note{}, err
//     }
//     var note Note
//     err = json.Unmarshal(data, &note)
//     return note, err
// }
//
// 5. Sanitize filename more thoroughly:
// func sanitizeFilename(title string) string {
//     // Remove invalid characters: / \ : * ? " < > |
//     // Limit length
//     // Handle edge cases
// }


// -------------------
🔑 KEY CONCEPTS:
1. JSON Marshaling Requirements
go// ❌ DOESN'T WORK - Unexported fields
type Note struct {
    title   string  // Ignored by json.Marshal
    content string  // Ignored by json.Marshal
}
// Result: {} (empty JSON)

// ✅ WORKS - Exported fields
type Note struct {
    Title   string  // Included in JSON
    Content string  // Included in JSON
}
// Result: {"Title":"...","Content":"..."}
```

### **2. File Permissions (0644)**
```
0644 = rw-r--r--

Owner (6 = 4+2):   read + write
Group (4):         read only
Others (4):        read only

Common permissions:
0644 - Regular files (readable by all, writable by owner)
0755 - Executable files (executable by all)
0600 - Private files (only owner can read/write)
3. String Manipulation Chain
go// Title: "My Shopping List"

// Step 1: Replace spaces
strings.ReplaceAll("My Shopping List", " ", "_")
// → "My_Shopping_List"

// Step 2: Lowercase
strings.ToLower("My_Shopping_List")
// → "my_shopping_list"

// Step 3: Add extension
"my_shopping_list" + ".json"
// → "my_shopping_list.json"

📊 Complete Example Output:
Code:
gonote, _ := note.New("Shopping List", "Buy milk and eggs")
note.Save()
File created: shopping_list.json
json{
  "Title": "Shopping List",
  "Content": "Buy milk and eggs",
  "CreatedAt": "2025-01-17T14:30:45.123456789Z"
}

💡 Advanced: Custom JSON Tags
gotype Note struct {
    Title     string    `json:"title"`           // Custom name
    Content   string    `json:"content"`         // Custom name
    CreatedAt time.Time `json:"created_at"`      // Snake case
    
    // Advanced tags:
    // `json:"name,omitempty"` - Omit if empty
    // `json:"-"` - Never serialize this field
    // `json:"name,string"` - Force to JSON string
}

// Result:
// {"title":"...","content":"...","created_at":"..."}

🚀 Summary of Changes:
AspectBeforeAfterWhyFieldsUnexportedExportedRequired for JSONEncapsulation✅ Protected❌ ExposedTrade-off for serializationMethodsDisplay onlyDisplay + SaveAdded persistenceFile formatN/AJSONStandard, readable formatFilenameN/AAuto-generatedBased on title
