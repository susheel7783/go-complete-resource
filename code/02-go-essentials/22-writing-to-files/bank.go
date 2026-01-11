package main // Declares this file is part of the main package (entry point for executable programs)

import ( // Imports multiple packages (parentheses for multiple imports)
	"fmt" // Package for formatted I/O operations
	"os"  // Package for operating system functions (file operations, environment variables, etc.)
)

// Function to write the account balance to a file
func writeBalanceToFile(balance float64) {
	// Parameter: balance is the current account balance to save
	// No return value (function doesn't return anything)
	
	balanceText := fmt.Sprint(balance)
	// fmt.Sprint: converts the balance (float64) to a string
	// Sprint = "String Print" - formats value as string without printing it
	// Example: if balance is 1250.5, balanceText becomes "1250.5"
	// Necessary because WriteFile expects bytes, not numbers
	
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
	// os.WriteFile: writes data to a file (creates file if it doesn't exist)
	// Parameters:
	//   1. "balance.txt" - filename (created in same directory as program)
	//   2. []byte(balanceText) - converts string to byte slice (required format)
	//   3. 0644 - file permissions (owner read/write, others read-only)
	//      - 0644 is octal notation: 6=rw-, 4=r--, 4=r--
	// If file exists, it's OVERWRITTEN (old content is replaced)
	// Note: Error handling is ignored (should check for write errors in production)
}

func main() { // The main function - program execution starts here
	
	// Declares and initializes account balance variable
	var accountBalance = 1000.0
	// Starting balance of 1000.0
	// This is only the initial value - if balance.txt exists, you might want to read from it
	
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
			// Displays current balance from memory (not from file)
			// Does NOT read from balance.txt
			
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
				// Balance is NOT updated, file is NOT written
			}
			
			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			// Adds deposit to balance in memory
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			// Confirmation message
			
			writeBalanceToFile(accountBalance)
			// Calls function to save new balance to balance.txt file
			// This persists the balance to disk (survives program restart)
			// File is created/overwritten with new balance
			
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
				// Balance is NOT updated, file is NOT written
			}
			
			// Second validation: sufficient balance check
			if withdrawalAmount > accountBalance {
				
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				
				continue
				// Skips rest of iteration
				// Balance is NOT updated, file is NOT written
			}
			
			accountBalance -= withdrawalAmount // accountBalance = accountBalance - withdrawalAmount
			// Subtracts withdrawal from balance in memory
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			// Confirmation message
			
			writeBalanceToFile(accountBalance)
			// Saves new balance to balance.txt file
			// Persists the withdrawal to disk
			// File is overwritten with updated balance
			
		default: // Exit or invalid choice
			
			fmt.Println("Goodbye!")
			
			fmt.Println("Thanks for choosing our bank")
			
			return
			// Exits program
			// Balance remains saved in balance.txt from last write
			
			// break (COMMENTED OUT)
		}
		
		// Loop continues to next iteration
		
	} // End of infinite for loop
	
}
```

**Key Concepts:**

- **File Operations**:
  - `os.WriteFile(filename, data, permissions)`: Writes data to a file
  - Creates file if it doesn't exist
  - Overwrites file if it exists (doesn't append)
  - Requires data as `[]byte` (byte slice)

- **Type Conversion**:
  - `fmt.Sprint(value)`: Converts any value to string
  - `[]byte(string)`: Converts string to byte slice
  - Chain: `float64 → string → []byte`

- **File Permissions (0644)**:
  - Octal notation (starts with 0)
  - Three groups: Owner, Group, Others
  - 6 (Owner): Read + Write (4+2)
  - 4 (Group): Read only
  - 4 (Others): Read only
  - Common for text files

- **Data Persistence**:
  - **Before**: Balance lost when program exits
  - **After**: Balance saved to `balance.txt` after each transaction
  - File persists between program runs
  - **Missing Feature**: Program doesn't READ from file on startup

**Current Limitation**:
- Balance always starts at 1000.0
- Even if balance.txt shows 5000, program ignores it
- Should read from file at startup to restore previous balance

**File Content Example**:
After depositing 250 (balance = 1250), `balance.txt` contains:
```
1250
