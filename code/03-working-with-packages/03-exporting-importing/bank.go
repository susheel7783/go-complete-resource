package main

import ( // Imports packages needed for this file
	"fmt" // Package for formatted I/O operations (printing, scanning)
	"example.com/bank/fileops" // Custom package import - file operations module
	// "example.com/bank" is the module path (defined in go.mod)
	// "fileops" is the package/folder name containing file operation functions
)

// Global constant - filename for storing account balance
const accountBalanceFile = "balance.txt"
// Defined here in main, passed to fileops functions as parameter

func main() { // The main function - program execution starts here
	
	// Load balance from file using custom package function
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	// fileops.GetFloatFromFile - calls function from imported fileops package
	// Must use package name prefix (fileops.) to access exported functions
	// Exported functions start with CAPITAL letter (GetFloatFromFile)
	// Returns (float64, error) - balance value and potential error
	
	if err != nil {
		// Error handling - file not found or corrupted data
		
		fmt.Println("ERROR")
		fmt.Println(err) // Prints the error message from fileops package
		fmt.Println("---------")
		
		// panic("Can't continue, sorry.") (COMMENTED OUT)
		// Instead of crashing, continues with default balance from fileops
	}
	
	fmt.Println("Welcome to Go Bank!")
	
	for { // Infinite loop - runs until user exits
		
		presentOptions()
		// Calls local function (defined elsewhere in main package)
		// No package prefix needed - same package, same directory
		// Displays menu: 1. Check, 2. Deposit, 3. Withdraw, 4. Exit
		
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice) // Reads user's menu selection
		
		// COMMENTED: wantsCheckBalance := choice == 1
		
		switch choice { // Evaluates user's choice
			
		case 1: // Check balance option
			fmt.Println("Your balance is", accountBalance)
			// Displays current balance stored in memory
			
		case 2: // Deposit money option
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount) // Read deposit amount from user
			
			if depositAmount <= 0 {
				// Validation: reject zero or negative deposits
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // Skip rest of iteration, show menu again
			}
			
			accountBalance += depositAmount // Add deposit to balance
			// accountBalance = accountBalance + depositAmount (expanded form)
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			// Calls fileops package function to save balance to file
			// Parameters: (value to write, filename)
			// Note: package prefix required (fileops.)
			
		case 3: // Withdraw money option
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount) // Read withdrawal amount
			
			if withdrawalAmount <= 0 {
				// Validation: reject zero or negative withdrawals
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // Skip rest, show menu again
			}
			
			if withdrawalAmount > accountBalance {
				// Validation: check sufficient funds
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue // Skip rest, show menu again
			}
			
			accountBalance -= withdrawalAmount // Subtract withdrawal
			// accountBalance = accountBalance - withdrawalAmount (expanded)
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			// Saves updated balance to file using fileops package
			
		default: // Exit or invalid choice (handles option 4 and any other number)
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return // Exit main function (terminates program)
			
			// break (COMMENTED OUT)
			// break would only exit switch, not loop
			// return exits entire function/program
		}
		
	} // End of infinite loop
	
}
```

-----------------------------------------------------

## **Key Concepts - Custom Package Imports:**

### **1. Module Structure:**
```
your-project/
├── go.mod              # Module definition file
├── main.go             # This file (main package)
└── fileops/            # Custom package folder
    └── fileops.go      # File operations code



// 2. Import Path Explained:
// goimport "example.com/bank/fileops"

// example.com/bank - Module path (defined in go.mod)
// fileops - Package/folder name
// This is NOT a real URL, just a unique identifier for your module

// 3. go.mod File Example:
// gomodule example.com/bank

// go 1.21  // Go version
// 4. Calling Exported Functions:
// go// In fileops/fileops.go:
// package fileops

// func GetFloatFromFile(filename string) (float64, error) {
//     // Capital 'G' = exported (public)
//     // Can be called from other packages
// }

// // In main.go:
// import "example.com/bank/fileops"

// fileops.GetFloatFromFile("balance.txt")
// // Must use: packageName.FunctionName

// Exported vs Unexported:
// Exported (Public):

// Start with CAPITAL letter
// Can be accessed from other packages
// Examples: GetFloatFromFile, WriteFloatToFile

// Unexported (Private):

// Start with lowercase letter
// Only accessible within same package
// Examples: parseBalance, validateData


// fileops/fileops.go (Likely looks like this):
// gopackage fileops // Package name (folder name)

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// // GetFloatFromFile reads a float64 from a file
// // Exported - starts with capital G
// func GetFloatFromFile(filename string) (float64, error) {
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		return 1000, errors.New("Failed to find file: " + filename)
// 	}
	
// 	valueText := string(data)
// 	value, err := strconv.ParseFloat(valueText, 64)
// 	if err != nil {
// 		return 1000, errors.New("Failed to parse file content.")
// 	}
	
// 	return value, nil
// }

// // WriteFloatToFile writes a float64 to a file
// // Exported - starts with capital W
// func WriteFloatToFile(value float64, filename string) error {
// 	valueText := fmt.Sprint(value)
// 	err := os.WriteFile(filename, []byte(valueText), 0644)
// 	return err // Could return error for better error handling
// }

// Benefits of This Approach:
// ✅ Reusability - fileops package can be used in multiple projects
// ✅ Separation of Concerns - File operations isolated from business logic
// ✅ Testability - Can test fileops package independently
// ✅ Cleaner Main - Main focuses on program flow, not implementation details
// ✅ Maintainability - Changes to file operations don't affect main logic

// Key Differences from Same-Package Split:
// Same Package (Multiple Files)Different PackagesNo import neededMust import packageAll functions accessibleOnly exported functions accessibleAll in same directorySeparate directoriespresentOptions()fileops.GetFloatFromFile()Lowercase OKMust use Capital for exports

// Quick Reference:
// Call local function (same package):
// gopresentOptions()  // No prefix needed
// Call imported package function:
// gofileops.GetFloatFromFile(filename)  // Package prefix required
// Export rules:
// go// fileops package:
// func GetFloatFromFile()   // ✅ Exported (Capital G)
// func parseData()          // ❌ Private (lowercase p)
