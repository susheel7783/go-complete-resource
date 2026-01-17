package note

import (
	"encoding/json" // Package for JSON encoding/decoding (serialization)
	"errors"        // Package for creating error values
	"fmt"           // Package for formatted I/O (printing)
	"os"            // Package for operating system functionality (file operations)
	"strings"       // Package for string manipulation
	"time"          // Package for working with dates and times
)

// ==================== NOTE STRUCT WITH JSON TAGS ====================
// Note represents a note with title, content, and creation timestamp
//
// STRUCT TAGS (the backtick strings after field names):
// Struct tags are metadata attached to struct fields
// Format: `key:"value"` or `key:"value,options"`
//
// JSON TAGS EXPLAINED:
// `json:"title"` tells encoding/json package:
// - When converting to JSON, use "title" as the key name (not "Title")
// - When reading JSON, look for "title" key
//
// WHY USE JSON TAGS?
// ✅ Control JSON field names (lowercase, snake_case, etc.)
// ✅ Follow JSON naming conventions (usually lowercase or snake_case)
// ✅ Maintain Go naming conventions (PascalCase for exported fields)
// ✅ Keep Go code and JSON output separate concerns
//
// WITHOUT TAGS:
// {"Title":"...", "Content":"...", "CreatedAt":"..."}  ← PascalCase (Go style)
//
// WITH TAGS:
// {"title":"...", "content":"...", "created_at":"..."}  ← snake_case (JSON convention)
type Note struct {
	// Title field
	// - Go code: note.Title (PascalCase, exported)
	// - JSON output: "title" (lowercase, as specified in tag)
	Title     string    `json:"title"`
	
	// Content field
	// - Go code: note.Content (PascalCase, exported)
	// - JSON output: "content" (lowercase, as specified in tag)
	Content   string    `json:"content"`
	
	// CreatedAt field
	// - Go code: note.CreatedAt (PascalCase, exported)
	// - JSON output: "created_at" (snake_case, as specified in tag)
	//
	// Note: snake_case is common in JSON for multi-word field names
	// Alternatives: "createdAt" (camelCase) or "created-at" (kebab-case)
	CreatedAt time.Time `json:"created_at"`
}

// ==================== STRUCT TAG SYNTAX BREAKDOWN ====================
//
// Basic syntax:
// FieldName Type `json:"json_name"`
//
// With options:
// FieldName Type `json:"json_name,omitempty"`
//
// Common json tag options:
// - omitempty: Omit field from JSON if it's empty/zero value
// - string: Force field to be encoded as JSON string
// - -: Never include this field in JSON
//
// Examples:
// `json:"title"`              → Use "title" as JSON key
// `json:"title,omitempty"`    → Use "title", omit if empty
// `json:"-"`                  → Never serialize this field
// `json:"count,string"`       → Encode number as string: "123"
//
// Multiple tags (different packages):
// `json:"title" xml:"Title" db:"note_title"`

// ==================== DISPLAY METHOD ====================
// Display shows the note's information in a formatted way
//
// No changes from previous version
// Still uses Go field names (note.Title, not the JSON tag)
func (note Note) Display() {
	// Print formatted note information
	// Uses Go struct field names (Title, Content) not JSON names
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", 
		note.Title,   // Go field name
		note.Content) // Go field name
}

// ==================== SAVE METHOD ====================
// Save writes the note to a JSON file
//
// JSON MARSHALING WITH TAGS:
// When json.Marshal(note) is called:
// 1. It looks at each exported field
// 2. Checks if field has a `json:"..."` tag
// 3. Uses tag name if present, field name if not
// 4. Applies any tag options (omitempty, etc.)
//
// Example transformation:
// Go struct:
// Note{Title: "Shopping", Content: "Buy milk", CreatedAt: time.Now()}
//
// JSON output (with tags):
// {
//   "title": "Shopping",
//   "content": "Buy milk",
//   "created_at": "2025-01-17T10:30:00Z"
// }
//
// Without tags it would be:
// {
//   "Title": "Shopping",
//   "Content": "Buy milk",
//   "CreatedAt": "2025-01-17T10:30:00Z"
// }
func (note Note) Save() error {
	// ==================== GENERATE FILENAME ====================
	// Create filename from title
	// Example: "My Shopping List" → "my_shopping_list.json"
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	
	// ==================== SERIALIZE TO JSON ====================
	// Convert Note struct to JSON bytes
	//
	// json.Marshal respects struct tags:
	// - Looks for `json:"..."` tags on each field
	// - Uses tag name for JSON keys
	// - If no tag, uses Go field name
	//
	// For our Note struct:
	// Title     → "title" (from tag)
	// Content   → "content" (from tag)
	// CreatedAt → "created_at" (from tag)
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	
	// ==================== WRITE TO FILE ====================
	// Write JSON bytes to file
	// File permissions: 0644 (rw-r--r--)
	return os.WriteFile(fileName, json, 0644)
}

// ==================== CONSTRUCTOR FUNCTION ====================
// New creates and returns a new Note with validation
//
// No changes from previous version
func New(title, content string) (Note, error) {
	// Validate input
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}
	
	// Create and return Note
	// Field names use Go conventions (Title, not title)
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

// ==================== JSON TAG EXAMPLES ====================
//
// Example 1: Basic tag (change JSON key name)
// type User struct {
//     Name string `json:"name"`  // JSON: {"name":"..."}
// }
//
// Example 2: omitempty option (omit if zero value)
// type User struct {
//     Name  string `json:"name"`
//     Email string `json:"email,omitempty"`  // Omit if empty string
//     Age   int    `json:"age,omitempty"`    // Omit if 0
// }
// // If Email is "", JSON: {"name":"John","age":25}
//
// Example 3: Ignore field (never include in JSON)
// type User struct {
//     Name     string `json:"name"`
//     Password string `json:"-"`  // Never serialized
// }
// // JSON: {"name":"John"} (password not included)
//
// Example 4: Force string encoding
// type Product struct {
//     ID    int    `json:"id,string"`  // "123" instead of 123
//     Price float64 `json:"price,string"`  // "19.99" instead of 19.99
// }
// // JSON: {"id":"123","price":"19.99"}
//
// Example 5: Multiple tags for different packages
// type Note struct {
//     Title string `json:"title" xml:"Title" db:"note_title"`
// }
// // json.Marshal uses "title"
// // xml.Marshal uses "Title"
// // Database mapper uses "note_title"

// ==================== COMPARISON: WITH vs WITHOUT TAGS ====================
//
// WITHOUT STRUCT TAGS:
// ─────────────────────────────────────────────────────────
// type Note struct {
//     Title     string
//     Content   string
//     CreatedAt time.Time
// }
//
// JSON OUTPUT:
// {
//   "Title": "Shopping List",
//   "Content": "Buy milk",
//   "CreatedAt": "2025-01-17T10:30:00Z"
// }
// ─────────────────────────────────────────────────────────
//
// WITH STRUCT TAGS:
// ─────────────────────────────────────────────────────────
// type Note struct {
//     Title     string    `json:"title"`
//     Content   string    `json:"content"`
//     CreatedAt time.Time `json:"created_at"`
// }
//
// JSON OUTPUT:
// {
//   "title": "Shopping List",
//   "content": "Buy milk",
//   "created_at": "2025-01-17T10:30:00Z"
// }
// ─────────────────────────────────────────────────────────

// ==================== REAL-WORLD JSON TAG USAGE ====================
//
// API Response struct:
// type APIResponse struct {
//     Success   bool        `json:"success"`
//     Data      interface{} `json:"data,omitempty"`
//     Error     string      `json:"error,omitempty"`
//     Timestamp time.Time   `json:"timestamp"`
// }
//
// Success case:
// {"success":true,"data":{...},"timestamp":"..."}
// (no "error" field because omitempty and it's empty)
//
// Error case:
// {"success":false,"error":"not found","timestamp":"..."}
// (no "data" field because omitempty and it's nil)

// ==================== JSON UNMARSHALING (READING JSON) ====================
//
// Struct tags work in BOTH directions:
//
// MARSHALING (Go → JSON):
// note := Note{Title: "Test"}
// json.Marshal(note)  // Uses tag: {"title":"Test"}
//
// UNMARSHALING (JSON → Go):
// jsonData := `{"title":"Test","content":"...","created_at":"..."}`
// var note Note
// json.Unmarshal([]byte(jsonData), &note)
// // Looks for "title" in JSON, maps to Title field
// fmt.Println(note.Title)  // "Test"
//
// This means you can read JSON files back into Note structs!

// ==================== POTENTIAL LOAD FUNCTION ====================
//
// You could add this to load notes from JSON files:
//
// func Load(filename string) (Note, error) {
//     // Read file
//     data, err := os.ReadFile(filename)
//     if err != nil {
//         return Note{}, err
//     }
//
//     // Unmarshal JSON to Note struct
//     var note Note
//     err = json.Unmarshal(data, &note)
//     if err != nil {
//         return Note{}, err
//     }
//
//     return note, nil
// }
//
// Usage:
// note, err := note.Load("shopping_list.json")
// if err != nil {
//     log.Fatal(err)
// }
// note.Display()

// ==================== ADVANCED: PRETTY-PRINTED JSON ====================
//
// For human-readable JSON files, use MarshalIndent:
//
// func (note Note) Save() error {
//     fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_")) + ".json"
//
//     // MarshalIndent adds newlines and indentation
//     // Parameters: (data, prefix, indent)
//     json, err := json.MarshalIndent(note, "", "  ")
//     if err != nil {
//         return err
//     }
//
//     return os.WriteFile(fileName, json, 0644)
// }
//
// BEFORE (Marshal):
// {"title":"Shopping","content":"Buy milk","created_at":"2025-01-17T10:30:00Z"}
//
// AFTER (MarshalIndent):
// {
//   "title": "Shopping",
//   "content": "Buy milk",
//   "created_at": "2025-01-17T10:30:00Z"
// }

// -------------
🔑 KEY CONCEPTS:
1. Struct Tags Syntax
gotype Note struct {
    // Basic tag
    Title string `json:"title"`
    
    // With option
    Email string `json:"email,omitempty"`
    
    // Never serialize
    Password string `json:"-"`
    
    // Multiple tags
    Name string `json:"name" xml:"Name" db:"user_name"`
}
2. Common JSON Tag Options
TagEffectExamplejson:"name"Use "name" as JSON key{"name":"..."}json:"name,omitempty"Omit if zero value{} if emptyjson:"-"Never serializeNot in JSONjson:"id,string"Force string type{"id":"123"}
3. JSON Output Comparison
Without tags:
json{
  "Title": "Shopping List",
  "Content": "Buy milk",
  "CreatedAt": "2025-01-17T10:30:00Z"
}
With tags:
json{
  "title": "Shopping List",
  "content": "Buy milk",
  "created_at": "2025-01-17T10:30:00Z"
}

📊 Complete Example:
Code:
gonote, _ := note.New("Shopping List", "Buy milk and eggs")
note.Save()
File created: shopping_list.json
json{"title":"Shopping List","content":"Buy milk and eggs","created_at":"2025-01-17T15:45:30.123456789Z"}
With json.MarshalIndent (pretty-printed):
json{
  "title": "Shopping List",
  "content": "Buy milk and eggs",
  "created_at": "2025-01-17T15:45:30.123456789Z"
}

💡 Advanced Usage Examples:
Example 1: omitempty
gotype Note struct {
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Tags      []string  `json:"tags,omitempty"`  // Omit if nil/empty
    CreatedAt time.Time `json:"created_at"`
}

// Without tags:
// {"title":"Test","content":"...","tags":null,"created_at":"..."}

// With omitempty and empty tags:
// {"title":"Test","content":"...","created_at":"..."}
Example 2: Hide Sensitive Fields
gotype User struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"-"`  // Never in JSON
}

// JSON output (password excluded):
// {"username":"john","email":"john@example.com"}
Example 3: Reading JSON Back
go// Load a saved note
func Load(filename string) (Note, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return Note{}, err
    }
    
    var note Note
    err = json.Unmarshal(data, &note)
    return note, err
}

// Usage:
savedNote, _ := note.Load("shopping_list.json")
savedNote.Display()

🚀 Benefits of Struct Tags:
✅ Follow conventions - JSON uses snake_case, Go uses PascalCase
✅ API compatibility - Match external API field names
✅ Flexibility - Same struct, different serialization formats
✅ Control - Decide what to include/exclude
✅ Bidirectional - Works for both Marshal and Unmarshal
