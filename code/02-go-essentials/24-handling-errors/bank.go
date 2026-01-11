package main // Declares this file is part of the main package (entry point for executable programs)

import ( // Imports multiple packages
	"errors" // Package for creating custom error messages
	"fmt"    // Package for formatted I/O operations
	"os"     // Package for operating system functions (file operations)
	"strconv" // Package for string conversions (string to number, number to string)
)

// Global constant for the filename
const accountBalanceFile = "balance.txt"
// const: declares a constant that cannot be changed
// Used throughout the program to reference the balance file

// Function to read the account balance from file with ERROR HANDLING
func getBalanceFromFile() (float64, error) {
	// Returns TWO values: (float64, error)
	// float64: the balance if successful
	// error: error message if something goes wrong, nil if successful
	// This is Go's standard pattern for error handling
	
	data, err := os.ReadFile(accountBalanceFile)
	// os.ReadFile returns ([]byte, error)
	// data: file contents as byte slice
	// err: error if file doesn't exist or can't be read
	
	if err != nil {
		// nil means "no value" or "no error"
		// err != nil means "an error occurred"
		// This checks if reading the file failed
		
		return 1000, errors.New("Failed to find balance file.")
		// Returns TWO values:
		//   1. 1000 - default balance (used when file doesn't exist)
		//   2. errors.New() - creates a new error with custom message
		// errors.New: creates an error value with the specified text
		// Function exits here if file read failed
	}
	
	balanceText := string(data)
	// Converts byte slice to string
	// Only reached if file was read successfully (no error above)
	
	balance, err := strconv.ParseFloat(balanceText, 64)
	// strconv.ParseFloat returns (float64, error)
	// balance: parsed number
	// err: error if string can't be converted to float
	// Note: err is REUSED (same variable name, different value)
	
	if err != nil {
		// Checks if parsing the string to float failed
		// Could fail if file contains "abc" instead of "1250.5"
		
		return 1000, errors.New("Failed to parse stored balance value.")
		// Returns default balance and custom error message
		// Function exits here if parsing failed
	}
	
	return balance, nil
	// SUCCESS case: returns the parsed balance and nil (no error)
	// nil error means "everything worked fine"
	// Only reached if both file read AND parsing succeeded
}

// Function to write the account balance to a file
func writeBalanceToFile(balance float64) {
	// Takes balance as parameter
	// No error handling here (could be improved)
	
	balanceText := fmt.Sprint(balance)
	// Converts float64 to string
	
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	// Writes to file
	// Note: Error from WriteFile is IGNORED (not good practice)
	// Should return error and handle it in main
}

func main() { // The main function - program execution starts here
	
	// Calls getBalanceFromFile and receives BOTH return values
	var accountBalance, err = getBalanceFromFile()
	// accountBalance: the balance value (either from file or default 1000)
	// err: error message if something went wrong, nil if successful
	
	if err != nil {
		// Checks if there was an error loading the balance
		// This block runs if file doesn't exist OR file has invalid data
		
		fmt.Println("ERROR")
		// Prints "ERROR" header
		
		fmt.Println(err)
		// Prints the actual error message
		// Will show either "Failed to find balance file." or "Failed to parse stored balance value."
		
		fmt.Println("---------")
		// Prints separator line
		
		// panic("Can't continue, sorry.") (COMMENTED OUT)
		// panic: crashes the program immediately with error message
		// Commented out because we're using default balance instead
		// Program continues with accountBalance = 1000
	}
	
	// Prints welcome message
	fmt.Println("Welcome to Go Bank!")
	// Program continues even if there was an error (using default balance)
	
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
			// Displays current balance
			// If there was an error earlier, this shows the default balance (1000)
			
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
			}
			
			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			// Adds deposit to balance in memory
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			// Confirmation message
			
			writeBalanceToFile(accountBalance)
			// Saves new balance to file
			// Creates balance.txt if it didn't exist before
			// Now future runs won't get the error
			
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
			
		default: // Exit or invalid choice
			
			fmt.Println("Goodbye!")
			
			fmt.Println("Thanks for choosing our bank")
			
			return
			// Exits program
			
			// break (COMMENTED OUT)
		}
		
		// Loop continues to next iteration
		
	} // End of infinite for loop
	
} 
--------------
// Key Concepts - Error Handling in Go:

// Multiple Return Values for Errors:

// Go functions return (value, error) instead of throwing exceptions
// Last return value is typically error type
// nil error means success, non-nil means failure


// Error Checking Pattern:

// go  result, err := someFunction()
//   if err != nil {
//       // Handle error
//   }
//   // Use result

// Creating Custom Errors:

// errors.New("message"): Creates new error with custom text
// More descriptive than just returning generic errors


// Error vs Panic:

// Error: Graceful handling, program continues with fallback
// Panic: Crashes program immediately (commented out here)
// This code uses errors + default value (1000) instead of crashing


// nil in Go:

// nil: Special value meaning "no value" or "zero value" for pointers, errors, slices, etc.
// err != nil: Checks if error exists
// err == nil: Checks if no error (success)



// Error Flow Examples:
// Scenario 1 - File doesn't exist (first run):

// getBalanceFromFile() called
// os.ReadFile fails → err != nil
// Returns (1000, errors.New("Failed to find balance file."))
// In main: err != nil → prints error message
// Program continues with accountBalance = 1000
// User makes first deposit → creates balance.txt

// Scenario 2 - File has invalid data ("abc"):

// getBalanceFromFile() called
// os.ReadFile succeeds (file exists)
// strconv.ParseFloat("abc", 64) fails → err != nil
// Returns (1000, errors.New("Failed to parse stored balance value."))
// In main: prints error, uses default 1000

// Scenario 3 - Everything works:

// getBalanceFromFile() called
// File exists and contains "1250.5"
// Both operations succeed
// Returns (1250.5, nil)
// In main: err == nil → no error message, uses 1250.5

// Improvement Opportunity:
// writeBalanceToFile should also return and handle errors:
// gofunc writeBalanceToFile(balance float64) error {
//     balanceText := fmt.Sprint(balance)
//     err := os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
//     return err // Return the error to caller
// }
