package note

import (
	"encoding/json" // JSON serialization/deserialization
	"errors"        // Error creation
	"fmt"           // Formatted output
	"os"            // File system operations
	"strings"       // String manipulation
	"time"          // Date and time handling
)

// ==================== NOTE STRUCT ====================
// Note represents a single note with metadata
//
// DESIGN CHOICES:
// ✅ Exported fields (Title, Content, CreatedAt) - Required for JSON marshaling
// ✅ JSON struct tags - Controls JSON field naming (lowercase/snake_case)
// ✅ Immutable by convention - No setter methods provided
// ✅ Auto-timestamping - CreatedAt set automatically in New()
//
// JSON OUTPUT:
// {"title":"My Note","content":"Note text","created_at":"2025-01-17T10:30:00Z"}
type Note struct {
	Title     string    `json:"title"`      // Note heading
	Content   string    `json:"content"`    // Note body
	CreatedAt time.Time `json:"created_at"` // Auto-set creation timestamp
}

// ==================== DISPLAY METHOD ====================
// Display prints the note in human-readable format
//
// RECEIVER: Value receiver (note Note) - receives a copy
// - Appropriate for read-only operations
// - Note is small, copying is cheap
func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", 
		note.Title, note.Content)
}

// ==================== SAVE METHOD ====================
// Save persists the note to a JSON file
//
// PROCESS:
// 1. Generate filename from title: "My Note" → "my_note.json"
// 2. Serialize to JSON using struct tags
// 3. Write to file with 0644 permissions (rw-r--r--)
//
// RETURNS: error (nil if successful)
func (note Note) Save() error {
	// Generate safe filename
	fileName := strings.ReplaceAll(note.Title, " ", "_") // Spaces → underscores
	fileName = strings.ToLower(fileName) + ".json"        // Lowercase + extension
	
	// Serialize to JSON
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	
	// Write to file
	return os.WriteFile(fileName, json, 0644)
}

// ==================== CONSTRUCTOR ====================
// New creates a validated Note
//
// VALIDATION:
// - Ensures title is not empty
// - Ensures content is not empty
// - Auto-sets CreatedAt to current time
//
// RETURN TYPE: (Note, error)
// - Returns value, not pointer
// - Simple for small structs
// - Alternative: (*Note, error) also valid
func New(title, content string) (Note, error) {
	// Validate inputs
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}
	
	// Create and return Note
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(), // Auto-generated
	}, nil
}

// ==================== PUBLIC API ====================
//
// TYPE: Note - Represents a note
// FUNCTION: New(title, content string) (Note, error) - Creates notes
// METHOD: Display() - Shows note
// METHOD: Save() error - Persists note
//
// USAGE:
// note, err := note.New("Title", "Content")
// if err != nil { /* handle */ }
// note.Display()
// note.Save()


// ----------------
🎯 CORE FEATURES:
FeatureImplementationBenefitValidationNew() checks empty stringsNo invalid notesJSON Tagsjson:"title" etc.Clean JSON outputAuto-timestamptime.Now() in constructorConsistent metadataFile namingSanitize + lowercaseSafe filenamesError handlingReturn errors explicitlyCaller control
WHAT IT DOES:
go// Create note
note, _ := note.New("Shopping", "Buy milk")

// Display
note.Display()
// Output: Your note titled Shopping has the following content:
//         
//         Buy milk

// Save
note.Save()
// Creates: shopping.json
// Content: {"title":"Shopping","content":"Buy milk","created_at":"..."}
