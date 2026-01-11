package main // Declares this file is part of the main package (entry point for executable programs)

import ( // Imports multiple packages needed for the program
	"errors"  // Package for creating custom error messages
	"fmt"     // Package for formatted I/O operations (printing, scanning)
	"os"      // Package for operating system functions (file operations)
	"strconv" // Package for string conversions (string ↔ number)
)

// Global constant - filename for storing account balance
const accountBalanceFile = "balance.txt"
// Accessible from any function, prevents typos, easy to change in one place

// Function to read account balance from file with error handling
func getBalanceFromFile() (float64, error) {
	// Returns: (balance, error) - Go's standard pattern for functions that can fail
	
	data, err := os.ReadFile(accountBalanceFile)
	// Reads entire file, returns ([]byte, error)
	// data: file contents as byte slice
	// err: error if file doesn't exist or can't be read
	
	if err != nil {
		// err != nil means an error occurred (file not found, no permission, etc.)
		
		return 1000, errors.New("Failed to find balance file.")
		// Returns default balance (1000) and custom error message
		// exits function here if file read failed
	}
	
	balanceText := string(data)
	// Converts byte slice to string (e.g., []byte{49,50,53,48} → "1250")
	
	balance, err := strconv.ParseFloat(balanceText, 64)
	// Converts string to float64 (e.g., "1250.5" → 1250.5)
	// 64: bit size for float64
	// err is reused - holds new error from parsing
	
	if err != nil {
		// Parsing failed (file contains invalid data like "abc")
		
		return 1000, errors.New("Failed to parse stored balance value.")
		// Returns default balance and error about corrupted data
	}
	
	return balance, nil
	// SUCCESS: returns parsed balance and nil (no error)
}

// Function to save account balance to file
func writeBalanceToFile(balance float64) {
	// Takes balance as parameter, no return value
	
	balanceText := fmt.Sprint(balance)
	// Converts float64 to string (e.g., 1250.5 → "1250.5")
	
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	// Writes string to file as bytes
	// 0644: file permissions (owner read/write, others read-only)
	// Note: ignores write errors (could be improved)
}

func main() { // The main function - program starts here
	
	// Load balance from file at startup
	var accountBalance, err = getBalanceFromFile()
	// accountBalance: loaded balance or default 1000
	// err: error message if loading failed, nil if successful
	
	if err != nil {
		// Handle error if balance couldn't be loaded
		
		fmt.Println("ERROR")
		fmt.Println(err) // Prints the error message
		fmt.Println("---------")
		
		// panic("Can't continue, sorry.") (COMMENTED OUT)
		// panic would crash the program
		// Instead, we continue with default balance (1000)
	}
	
	fmt.Println("Welcome to Go Bank!")
	
	for { // Infinite loop - runs until user exits
		
		presentOptions()
		// Calls function to display menu (defined elsewhere in code)
		// Shows: 1. Check balance, 2. Deposit, 3. Withdraw, 4. Exit
		
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice) // Reads user's menu selection
		
		// COMMENTED: wantsCheckBalance := choice == 1
		
		switch choice { // Evaluates user's choice
			
		case 1: // Check balance option
			fmt.Println("Your balance is", accountBalance)
			// Displays current balance from memory
			
		case 2: // Deposit money option
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount) // Read deposit amount
			
			if depositAmount <= 0 {
				// Validation: reject zero or negative amounts
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // Skip rest of loop, show menu again
			}
			
			accountBalance += depositAmount // Add deposit to balance
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance) // Save to file
			
		case 3: // Withdraw money option
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount) // Read withdrawal amount
			
			if withdrawalAmount <= 0 {
				// Validation: reject zero or negative amounts
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // Skip rest, show menu again
			}
			
			if withdrawalAmount > accountBalance {
				// Validation: check sufficient funds
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue // Skip rest, show menu again
			}
			
			accountBalance -= withdrawalAmount // Subtract withdrawal from balance
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance) // Save to file
			
		default: // Exit or invalid choice
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return // Exit program
		}
		
	} // End of infinite loop
	
}

