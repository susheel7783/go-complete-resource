package main

import ( // Imports multiple packages
	"fmt"    // Package for formatted I/O operations
	"os"     // Package for operating system functions (file operations)
	"strconv" // Package for string conversions (string to number, number to string)
)

// Global constant for the filename
const accountBalanceFile = "balance.txt"
// const: declares a constant (value cannot be changed)
// Global scope: accessible from any function in this file
// Using a constant prevents typos and makes it easy to change filename in one place
// Convention: constants often use camelCase or UPPER_CASE naming

// Function to read the account balance from file
func getBalanceFromFile() float64 {
	// No parameters needed (uses global constant for filename)
	// Returns: float64 (the balance read from file)
	
	data, _ := os.ReadFile(accountBalanceFile)
	// os.ReadFile: reads entire file contents and returns ([]byte, error)
	// accountBalanceFile: uses the constant "balance.txt"
	// data: contains the file contents as a byte slice ([]byte)
	// _ (blank identifier): ignores the error return value
	//   - If file doesn't exist, data will be empty (nil)
	//   - This is risky - better to check for errors in production
	
	balanceText := string(data)
	// string(data): converts byte slice to string
	// Example: []byte{49, 50, 53, 48} becomes "1250"
	// If file doesn't exist, balanceText will be empty string ""
	
	balance, _ := strconv.ParseFloat(balanceText, 64)
	// strconv.ParseFloat: converts string to float64
	// Parameters:
	//   1. balanceText - the string to convert (e.g., "1250.5")
	//   2. 64 - bit size (use 64 for float64, 32 for float32)
	// Returns: (float64, error)
	//   - balance: the converted number
	//   - _: error is ignored (risky - parsing could fail)
	// If balanceText is empty or invalid, balance will be 0.0
	
	return balance
	// Returns the balance (or 0.0 if file doesn't exist or is invalid)
}

// Function to write the account balance to a file
func writeBalanceToFile(balance float64) {
	// Parameter: balance is the current account balance to save
	// No return value
	
	balanceText := fmt.Sprint(balance)
	// fmt.Sprint: converts balance (float64) to string
	// Example: 1250.5 becomes "1250.5"
	
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	// os.WriteFile: writes data to file
	// Parameters:
	//   1. accountBalanceFile - filename constant ("balance.txt")
	//   2. []byte(balanceText) - converts string to byte slice
	//   3. 0644 - file permissions (owner read/write, others read-only)
	// Creates file if it doesn't exist, overwrites if it does
	// Error is ignored (not good practice, but simplified here)
}

func main() { // The main function - program execution starts here
	
	// Reads balance from file (or gets 0.0 if file doesn't exist)
	var accountBalance = getBalanceFromFile()
	// Calls getBalanceFromFile() on program startup
	// Balance persists between program runs!
	// First run: file doesn't exist → balance = 0.0
	// Subsequent runs: reads saved balance from file
	
	// Prints welcome message once
	fmt.Println("Welcome to Go Bank!")
	
	for { // Infinite loop - runs until user exits
		
		// Displays menu options
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		
		// Declares variable to store user's choice
		var choice int
		
		// Prompts user for input
		fmt.Print("Your choice: ")
		
		// Reads user input
		fmt.Scan(&choice)
		
		// COMMENTED OUT: Alternative boolean approach
		// wantsCheckBalance := choice == 1
		
		// Switch statement to handle user's choice
		switch choice {
		
		case 1: // Check balance
			
			fmt.Println("Your balance is", accountBalance)
			// Displays current balance from memory
			// This reflects the persisted value from file
			
		case 2: // Deposit money
			
			fmt.Print("Your deposit: ")
			
			var depositAmount float64
			// Variable to store deposit amount
			
			fmt.Scan(&depositAmount)
			// Reads deposit amount from user
			
			// Validation: checks if deposit is positive
			if depositAmount <= 0 {
				
				fmt.Println("Invalid amount. Must be greater than 0.")
				
				// return (COMMENTED OUT)
				
				continue
				// Skips rest of iteration, shows menu again
				// Balance NOT updated, file NOT written
			}
			
			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			// Adds deposit to balance in memory
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			// Confirmation message
			
			writeBalanceToFile(accountBalance)
			// Saves new balance to file
			// Now when program restarts, this new balance will be loaded
			
		case 3: // Withdraw money
			
			fmt.Print("Withdrawal amount: ")
			
			var withdrawalAmount float64
			// Variable to store withdrawal amount
			
			fmt.Scan(&withdrawalAmount)
			// Reads withdrawal amount
			
			// First validation: amount must be positive
			if withdrawalAmount <= 0 {
				
				fmt.Println("Invalid amount. Must be greater than 0.")
				
				continue
				// Skips rest of iteration
			}
			
			// Second validation: sufficient balance check
			if withdrawalAmount > accountBalance {
				
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				
				continue
				// Skips rest of iteration
			}
			
			accountBalance -= withdrawalAmount // accountBalance = accountBalance - withdrawalAmount
			// Subtracts withdrawal from balance in memory
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			// Confirmation message
			
			writeBalanceToFile(accountBalance)
			// Saves new balance to file
			// Persists the withdrawal to disk
			
		default: // Exit or invalid choice
			
			fmt.Println("Goodbye!")
			
			fmt.Println("Thanks for choosing our bank")
			
			return
			// Exits program
			// Balance remains saved in balance.txt
			
			// break (COMMENTED OUT)
		}
		
		// Loop continues to next iteration
		
	} // End of infinite for loop
	
}

// Global Constants:

// const accountBalanceFile = "balance.txt": Defined outside functions
// Accessible from anywhere in the package
// Prevents typos, easier to maintain
// Cannot be changed at runtime


// strconv Package:

// strconv.ParseFloat(string, bitSize): Converts string to float
// First parameter: string to parse ("1250.5")
// Second parameter: bit size (64 for float64)
// Returns: (float64, error) - converted value and any parsing error


// Error Handling with Blank Identifier:

// data, _ or balance, _: Ignores error return values
// _ (underscore): blank identifier that discards values
// Risky practice: File might not exist, parsing might fail
// Production code should check errors: if err != nil { ... }


// Complete Persistence Cycle:

// Program Start: getBalanceFromFile() reads saved balance
// Transactions: Balance updates in memory
