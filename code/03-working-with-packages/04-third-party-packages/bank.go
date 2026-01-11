package main // Declares this is the main package (entry point for executable)

import ( // Imports packages needed for this file
	"fmt" // Package for formatted I/O operations (printing, scanning)
	"example.com/bank/fileops" // Custom/local package - your file operations module
	"github.com/Pallinder/go-randomdata" // Third-party package - generates random data
	// This is an EXTERNAL package from GitHub
	// Must be downloaded first: go get github.com/Pallinder/go-randomdata
)

// Global constant - filename for storing account balance
const accountBalanceFile = "balance.txt"

func main() { // The main function - program execution starts here
	
	// Load balance from file using custom package
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	// Calls function from your custom fileops package
	// Returns (float64, error) - balance and potential error
	
	if err != nil {
		// Error handling - file not found or corrupted
		
		fmt.Println("ERROR")
		fmt.Println(err) // Prints error message
		fmt.Println("---------")
		
		// panic("Can't continue, sorry.") (COMMENTED OUT)
		// Continues with default balance instead of crashing
	}
	
	fmt.Println("Welcome to Go Bank!")
	
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	// randomdata.PhoneNumber() - calls function from third-party package
	// Generates a random fake phone number each time program runs
	// Example output: "Reach us 24/7 (555) 123-4567"
	// This demonstrates using an EXTERNAL package from GitHub
	// The package must be installed before running: go get github.com/Pallinder/go-randomdata
	
	for { // Infinite loop - runs until user exits
		
		presentOptions()
		// Calls local function (same package, different file)
		// Displays banking menu options
		
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
				// Validation: reject zero or negative deposits
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // Skip rest of iteration, show menu again
			}
			
			accountBalance += depositAmount // Add deposit to balance
			// accountBalance = accountBalance + depositAmount
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			// Saves balance to file using custom package
			
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
			// accountBalance = accountBalance - withdrawalAmount
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			// Saves updated balance to file
			
		default: // Exit or invalid choice
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return // Exit program
			
			// break (COMMENTED OUT)
		}
		
	} // End of infinite loop
	
}
// -------
// 1. Types of Imports:
// goimport (
//     "fmt"                              // ① Standard library (built into Go)
//     "example.com/bank/fileops"         // ② Custom/local package (your code)
//     "github.com/Pallinder/go-randomdata" // ③ Third-party package (external)
// )
