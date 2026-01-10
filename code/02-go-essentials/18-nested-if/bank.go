package main

import "fmt" // Imports the fmt package for formatted I/O operations

func main() { // The main function - program execution starts here
	
	// Declares and initializes account balance variable
	var accountBalance = 1000.0
	// Starting balance of 1000.0 (float64 type inferred from decimal)
	
	// Prints the welcome message to the console
	fmt.Println("Welcome to Go Bank!")
	
	// Prints a prompt asking what the user wants to do
	fmt.Println("What do you want to do?")
	
	// Displays menu option 1
	fmt.Println("1. Check balance")
	
	// Displays menu option 2
	fmt.Println("2. Deposit money")
	
	// Displays menu option 3
	fmt.Println("3. Withdraw money")
	
	// Displays menu option 4
	fmt.Println("4. Exit")
	
	// Declares a variable to store the user's menu choice
	var choice int
	// int type: stores whole numbers (integers)
	
	// Prompts user to enter their choice
	fmt.Print("Your choice: ")
	// Print: keeps cursor on same line for user input
	
	// Reads the user's input and stores it in the choice variable
	fmt.Scan(&choice)
	// &choice: passes the memory address so Scan can modify the value
	
	// COMMENTED OUT: Alternative way to check if user chose option 1
	// wantsCheckBalance := choice == 1
	// Would create a boolean variable based on comparison
	
	// If statement: checks if user chose option 1 (Check balance)
	if choice == 1 {
		// == is the equality comparison operator
		
		fmt.Println("Your balance is", accountBalance)
		// Displays the current account balance
		
	} else if choice == 2 {
		// Handles option 2: Deposit money
		
		fmt.Print("Your deposit: ")
		// Prompts user to enter deposit amount
		
		var depositAmount float64
		// Declares variable to store deposit amount (decimal number)
		
		fmt.Scan(&depositAmount)
		// Reads the deposit amount from user
		
		// Validation: checks if deposit amount is valid
		if depositAmount <= 0 {
			// <= means "less than or equal to"
			// Deposit must be a positive number (greater than 0)
			
			fmt.Println("Invalid amount. Must be greater than 0.")
			// Displays error message
			
			return
			// return: exits the main function immediately
			// Program stops here if deposit is invalid
			// No code after this point in main() will execute
		}
		
		accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
		// += adds depositAmount to current balance and stores the result
		// Example: 1000 + 250 = 1250
		
		fmt.Println("Balance updated! New amount:", accountBalance)
		// Confirms deposit and shows new balance
		
	} else if choice == 3 {
		// Handles option 3: Withdraw money
		
		fmt.Print("Withdrawal amount: ")
		// Prompts user to enter withdrawal amount
		
		var withdrawalAmount float64
		// Declares variable to store withdrawal amount
		
		fmt.Scan(&withdrawalAmount)
		// Reads the withdrawal amount from user
		
		// First validation: checks if withdrawal amount is positive
		if withdrawalAmount <= 0 {
			// Withdrawal must be greater than 0
			
			fmt.Println("Invalid amount. Must be greater than 0.")
			// Displays error message
			
			return
			// Exits the program immediately if amount is invalid
		}
		
		// Second validation: checks if user has enough money
		if withdrawalAmount > accountBalance {
			// > means "greater than"
			// Can't withdraw more money than what's in the account
			
			fmt.Println("Invalid amount. You can't withdraw more than you have.")
			// Displays insufficient funds error
			
			return
			// Exits the program immediately if insufficient funds
		}
		
		accountBalance -= withdrawalAmount // accountBalance = accountBalance - withdrawalAmount
		// -= subtracts withdrawalAmount from current balance
		// Example: 1000 - 250 = 750
		
		fmt.Println("Balance updated! New amount:", accountBalance)
		// Confirms withdrawal and shows new balance
		
	} else {
		// Handles any other choice (including option 4: Exit, or invalid choices)
		// This is the "catch-all" for anything that isn't 1, 2, or 3
		
		fmt.Println("Goodbye!")
		// Displays exit message
		// Could be for choice 4, or any invalid number like 5, 10, -1, etc.
	}
	
	// Program ends here (after the if-else block completes)
	// Unless return was called earlier due to invalid input
}

// ------------------
// Nested If Statements:

// You can have if statements inside other if or else if blocks
// Used here for validation checks within deposit and withdrawal operations


// Comparison Operators:

// <=: Less than or equal to
// >: Greater than
// ==: Equal to
