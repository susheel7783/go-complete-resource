package note

import (
	"encoding/json" // JSON encoding/decoding for serialization
	"errors"        // Creating error values
	"fmt"           // Formatted I/O for printing
	"os"            // Operating system file operations
	"strings"       // String manipulation utilities
	"time"          // Date and time functionality
)

// ==================== NOTE STRUCT - COMPLETE DATA MODEL ====================
// Note represents a single note with title, content, and metadata
//
// DESIGN DECISIONS:
// ✅ Exported fields (Title, Content, CreatedAt) - Required for JSON serialization
// ✅ JSON struct tags - Control JSON output format
// ✅ Immutable design - No setter methods (create once, don't modify)
// ✅ Automatic timestamp - CreatedAt set in constructor, not by user
//
// STRUCT TAG BREAKDOWN:
// `json:"title"` means:
// - When marshaling to JSON: use "title" as key (not "Title")
// - When unmarshaling from JSON: look for "title" key
// - Follows JSON conventions (lowercase/snake_case) while keeping Go conventions (PascalCase)
//
// JSON OUTPUT FORMAT:
// {
//   "title": "My Note",
//   "content": "Note content here",
//   "created_at": "2025-01-17T10:30:00Z"
// }
type Note struct {
	// Title: The note's title/heading
	// - Exported for JSON serialization
	// - JSON key: "title" (lowercase via tag)
	// - Required field (validated in New())
	Title     string    `json:"title"`
	
	// Content: The main body/text of the note
	// - Exported for JSON serialization
	// - JSON key: "content" (lowercase via tag)
	// - Required field (validated in New())
	Content   string    `json:"content"`
	
	// CreatedAt: Timestamp when note was created
	// - Exported for JSON serialization
	// - JSON key: "created_at" (snake_case via tag)
	// - Automatically set in New(), user cannot specify
	// - time.Time serializes to ISO 8601 format: "2025-01-17T10:30:00Z"
	CreatedAt time.Time `json:"created_at"`
}

// ==================== DISPLAY METHOD - USER OUTPUT ====================
// Display prints the note in a human-readable format to stdout
//
// METHOD DESIGN:
// - Receiver: (note Note) - VALUE RECEIVER
//   * Receives a copy of the Note
//   * Appropriate because we only read data, don't modify
//   * Note struct is small, copying is cheap
// - Exported: Display (capitalized) - callable from other packages
// - No parameters: Works with the note it's called on
// - No return value: Prints directly to stdout
//
// OUTPUT FORMAT:
// "Your note titled [Title] has the following content:
//
// [Content]"
//
// USAGE:
// note.Display()  // Called on a Note instance
func (note Note) Display() {
	// Printf with format specifiers:
	// %v = default format (prints value as-is)
	// \n = newline character
	// \n\n = blank line for better readability
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", 
		note.Title,   // First %v replaced with title
		note.Content) // Second %v replaced with content
	
	// Note: CreatedAt is not displayed in this method
	// Could add: fmt.Printf("Created: %v\n", note.CreatedAt)
}

// ==================== SAVE METHOD - PERSISTENCE ====================
// Save writes the note to a JSON file on disk
//
// METHOD DESIGN:
// - Receiver: (note Note) - VALUE RECEIVER
//   * Could use pointer receiver for consistency
//   * Value is fine since we don't modify the note
// - Returns: error
//   * nil if save successful
//   * error object if any step fails
//
// PROCESS:
// 1. Generate filename from title
// 2. Serialize note to JSON
// 3. Write JSON to file
//
// FILE NAMING STRATEGY:
// "My Shopping List" → "my_shopping_list.json"
// - Spaces replaced with underscores (safe for filesystems)
// - Converted to lowercase (consistent naming)
// - .json extension added (indicates file format)
//
// USAGE:
// err := note.Save()
// if err != nil {
//     // Handle error (disk full, permissions, etc.)
// }
func (note Note) Save() error {
	// ==================== STEP 1: FILENAME GENERATION ====================
	// Create a filesystem-safe filename from the note title
	
	// Replace all spaces with underscores
	// strings.ReplaceAll(source, old, new)
	// Example: "My Shopping List" → "My_Shopping_List"
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	
	// Convert to lowercase for consistency
	// Example: "My_Shopping_List" → "my_shopping_list"
	fileName = strings.ToLower(fileName)
	
	// Add .json file extension
	// Example: "my_shopping_list" → "my_shopping_list.json"
	fileName = fileName + ".json"
	
	// One-liner alternative:
	// fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_")) + ".json"
	
	// ==================== STEP 2: JSON SERIALIZATION ====================
	// Convert the Note struct to JSON bytes
	//
	// json.Marshal process:
	// 1. Examines each exported field in Note
	// 2. Checks for `json:"..."` struct tags
	// 3. Uses tag name if present, field name if not
	// 4. Serializes time.Time to ISO 8601 format
	// 5. Returns []byte (JSON data) and error
	//
	// Example transformation:
	// Note{
	//   Title: "Shopping",
	//   Content: "Buy milk",
	//   CreatedAt: time.Now()
	// }
	// ↓
	// {"title":"Shopping","content":"Buy milk","created_at":"2025-01-17T10:30:00Z"}
	//
	// Note: Variable name 'json' shadows the package name
	// This is intentional and common practice in Go
	json, err := json.Marshal(note)
	
	// Check if marshaling failed
	// Unlikely to fail for simple Note struct
	// Could fail with: circular references, channels, funcs, etc.
	if err != nil {
		return err // Propagate error to caller
	}
	
	// ==================== STEP 3: FILE WRITE ====================
	// Write JSON bytes to file on disk
	//
	// os.WriteFile is a convenience function that:
	// 1. Creates file (or truncates if exists)
	// 2. Writes data
	// 3. Closes file
	// 4. All atomically
	//
	// Parameters:
	// - fileName: "my_shopping_list.json"
	// - json: []byte containing JSON data
	// - 0644: Unix file permissions (octal notation)
	//
	// PERMISSION BREAKDOWN (0644):
	// 0 = octal prefix
	// 6 = owner: read(4) + write(2) = rw-
	// 4 = group: read(4) = r--
	// 4 = other: read(4) = r--
	// Result: -rw-r--r-- (owner can read/write, others can read)
	//
	// Returns error if:
	// - Cannot create file (permission denied)
	// - Cannot write data (disk full, I/O error)
	// - Invalid filename
	return os.WriteFile(fileName, json, 0644)
	
	// Note: Directly returning the error from WriteFile
	// If successful, WriteFile returns nil
	// Caller should check: if err := note.Save(); err != nil { ... }
}

// ==================== CONSTRUCTOR FUNCTION - OBJECT CREATION ====================
// New creates and returns a new Note with validation
//
// CONSTRUCTOR PATTERN:
// - Go doesn't have built-in constructors
// - Convention: function named "New" or "NewTypeName"
// - In dedicated package: just "New" (called as note.New)
// - In shared package: "NewNote" (called as pkg.NewNote)
//
// FUNCTION DESIGN:
// - Parameters: Required fields only (title, content)
// - Auto-sets: createdAt (user doesn't provide)
// - Returns: (Note, error) - both result and error status
// - Validates: All inputs before creating Note
//
// RETURN TYPE: Value vs Pointer
// - Returns Note (value), not *Note (pointer)
// - Alternative design would return (*Note, error)
// - Both valid - this is simpler for small structs
//
// ERROR HANDLING:
// - Returns empty Note{} and error if validation fails
// - Returns populated Note and nil if successful
// - Go's standard pattern for fallible operations
//
// USAGE:
// note, err := note.New("Title", "Content")
// if err != nil {
//     // Handle validation error
// }
// // Use note
func New(title, content string) (Note, error) {
	// ==================== INPUT VALIDATION ====================
	// Validate that all required fields are non-empty
	//
	// Validation rules:
	// - Title must not be empty string
	// - Content must not be empty string
	// - Both required (using OR operator)
	//
	// Why validate?
	// ✅ Data integrity - no invalid notes can exist
	// ✅ Fail fast - catch errors at creation time
	// ✅ Clear API - users know requirements
	// ✅ Consistent state - all Notes are valid
	if title == "" || content == "" {
		// Validation failed - return zero values
		//
		// Note{} creates a zero-value Note:
		// - Title: "" (empty string)
		// - Content: "" (empty string)
		// - CreatedAt: zero time (January 1, year 1, 00:00:00 UTC)
		//
		// errors.New creates error with given message
		// Caller receives this error and can display/handle it
		return Note{}, errors.New("Invalid input.")
		
		// More descriptive alternative:
		// return Note{}, errors.New("Title and content cannot be empty.")
		//
		// More specific alternatives:
		// if title == "" {
		//     return Note{}, errors.New("Title is required.")
		// }
		// if content == "" {
		//     return Note{}, errors.New("Content is required.")
		// }
	}
	
	// ==================== NOTE CREATION ====================
	// Validation passed - create and return Note
	//
	// Struct literal initialization:
	// - Explicitly name each field
	// - User-provided values: title, content
	// - Auto-generated value: createdAt
	//
	// Key design point:
	// time.Now() is called HERE, not by user
	// - Ensures accurate timestamp
	// - User cannot fake creation time
	// - Consistent across all Notes
	return Note{
		Title:     title,      // From parameter
		Content:   content,    // From parameter
		CreatedAt: time.Now(), // Auto-set to current timestamp
	}, nil // nil error means success
	
	// The note is returned BY VALUE (copied)
	// Caller gets their own copy of the Note
	// Changes to copy won't affect original (if there were setters)
}

// ==================== PACKAGE PUBLIC API SUMMARY ====================
//
// EXPORTED (Public - usable from other packages):
// ┌─────────────────────────────────────────────────────────┐
// │ TYPE: Note                                              │
// │ - Can declare variables: var n note.Note                │
// │ - Can use as parameter: func Process(n note.Note)       │
// │ - Cannot directly create: must use New()               │
// ├─────────────────────────────────────────────────────────┤
// │ FUNCTION: New(title, content string) (Note, error)      │
// │ - Creates validated Note                                │
// │ - Only way to create Note (enforces validation)        │
// ├─────────────────────────────────────────────────────────┤
// │ METHOD: Display()                                       │
// │ - Prints note to stdout                                 │
// │ - Called as: note.Display()                            │
// ├─────────────────────────────────────────────────────────┤
// │ METHOD: Save() error                                    │
// │ - Persists note to JSON file                           │
// │ - Called as: err := note.Save()                        │
// └─────────────────────────────────────────────────────────┘
//
// FIELDS (Exported but should use through API):
// - note.Title (accessible but prefer using Display())
// - note.Content (accessible but prefer using Display())
// - note.CreatedAt (accessible but read-only in practice)

// ==================== USAGE EXAMPLES ====================
//
// Example 1: Basic usage
// ───────────────────────────────────────────────────────────
// import "example.com/note/note"
//
// func main() {
//     // Create note
//     myNote, err := note.New("Shopping", "Buy milk")
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     // Display it
//     myNote.Display()
//     // Output:
//     // Your note titled Shopping has the following content:
//     //
//     // Buy milk
//
//     // Save it
//     err = myNote.Save()
//     if err != nil {
//         log.Fatal(err)
//     }
//     // Creates file: shopping.json
// }
//
// Example 2: Error handling
// ───────────────────────────────────────────────────────────
// note, err := note.New("", "Content")
// if err != nil {
//     fmt.Println(err) // "Invalid input."
//     return
// }
//
// Example 3: Accessing fields (discouraged but possible)
// ───────────────────────────────────────────────────────────
// note, _ := note.New("Title", "Content")
// fmt.Println(note.Title)     // "Title"
// fmt.Println(note.CreatedAt) // 2025-01-17 10:30:00...
//
// // Can modify (but shouldn't - breaks validation)
// note.Title = ""  // ⚠️ Now invalid, but no error

// ==================== DESIGN PATTERNS DEMONSTRATED ====================
//
// 1. CONSTRUCTOR PATTERN
//    - New() function controls object creation
//    - Ensures valid state at construction
//
// 2. VALIDATION AT CREATION
//    - Invalid objects cannot exist
//    - Fail fast - errors caught early
//
// 3. ENCAPSULATION (Partial)
//    - Fields exported (required for JSON)
//    - Creation controlled via New()
//    - Best practice: use methods, not direct field access
//
// 4. IMMUTABILITY (By Convention)
//    - No setter methods provided
//    - Notes don't change after creation
//    - Fields could be modified but shouldn't be
//
// 5. ERROR HANDLING
//    - Explicit error returns
//    - Caller must check errors
//    - Idiomatic Go pattern
//
// 6. SINGLE RESPONSIBILITY
//    - Note: Data representation
//    - New(): Creation and validation
//    - Display(): Output formatting
//    - Save(): Persistence

// ==================== PRODUCTION ENHANCEMENTS ====================
//
// Possible additions:
//
// 1. Load function (read from file):
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
// 2. Pretty-printed JSON:
// func (note Note) Save() error {
//     fileName := /* generate filename */
//     json, err := json.MarshalIndent(note, "", "  ")
//     /* ... */
// }
//
// 3. Getter methods (more encapsulation):
// func (n Note) GetTitle() string { return n.Title }
// func (n Note) GetContent() string { return n.Content }
// func (n Note) GetCreatedAt() time.Time { return n.CreatedAt }
//
// 4. Update methods (would need pointer receivers):
// func (n *Note) UpdateContent(content string) error {
//     if content == "" {
//         return errors.New("Content cannot be empty")
//     }
//     n.Content = content
//     return nil
// }
//
// 5. Additional validation:
// - Max title length
// - Max content length
// - Allowed characters in title
// - Content format validation
