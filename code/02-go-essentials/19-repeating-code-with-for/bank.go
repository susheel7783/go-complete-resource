package main // Declares this file is part of the main package (entry point for executable programs)

import "fmt" // Imports the fmt package for formatted I/O operations

func main() { // The main function - program execution starts here
	
	// Declares and initializes account balance variable
	var accountBalance = 1000.0
	// Starting balance of 1000.0 (float64 type inferred from decimal)
	// This variable persists across all loop iterations
	
	// For loop: repeats the banking menu 200 times
	for i := 0; i < 200; i++ {
		
		// Prints the welcome message to the console
		fmt.Println("Welcome to Go Bank!")
		// This prints at the START of each loop iteration
		
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
		// This variable is re-created each loop iteration (local to loop body)
		
		// Prompts user to enter their choice
		fmt.Print("Your choice: ")
		// Cursor stays on same line for user input
		
		// Reads the user's input and stores it in the choice variable
		fmt.Scan(&choice)
		// Waits for user to type a number and press Enter
		
		// COMMENTED OUT: Alternative way to check if user chose option 1
		// wantsCheckBalance := choice == 1
		
		// If statement: checks if user chose option 1 (Check balance)
		if choice == 1 {
			// Executes if user enters 1
			
			fmt.Println("Your balance is", accountBalance)
			// Displays the current account balance
			// Balance is preserved across loop iterations
			
		} else if choice == 2 {
			// Handles option 2: Deposit money
			
			fmt.Print("Your deposit: ")
			// Prompts for deposit amount
			
			var depositAmount float64
			// Stores the deposit amount
			
			fmt.Scan(&depositAmount)
			// Reads deposit amount from user
			
			// Validation: checks if deposit amount is valid
			if depositAmount <= 0 {
				// Deposit must be positive
				
				fmt.Println("Invalid amount. Must be greater than 0.")
				// Error message
				
				return
				// EXITS THE ENTIRE PROGRAM (exits main function)
				// This also BREAKS OUT of the loop
				// Program terminates completely here if validation fails
			}
			
			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			// Adds deposit to balance
			// This change persists for future loop iterations
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			// Confirmation message with new balance
			
		} else if choice == 3 {
			// Handles option 3: Withdraw money
			
			fmt.Print("Withdrawal amount: ")
			// Prompts for withdrawal amount
			
			var withdrawalAmount float64
			// Stores the withdrawal amount
			
			fmt.Scan(&withdrawalAmount)
			// Reads withdrawal amount from user
			
			// First validation: checks if withdrawal amount is positive
			if withdrawalAmount <= 0 {
				// Amount must be greater than 0
				
				fmt.Println("Invalid amount. Must be greater than 0.")
				// Error message
				
				return
				// EXITS THE ENTIRE PROGRAM
				// Terminates immediately on invalid amount
			}
			
			// Second validation: checks if user has enough money
			if withdrawalAmount > accountBalance {
				// Can't withdraw more than available balance
				
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				// Insufficient funds error
				
				return
				// EXITS THE ENTIRE PROGRAM
				// Terminates immediately on insufficient funds
			}
			
			accountBalance -= withdrawalAmount // accountBalance = accountBalance - withdrawalAmount
			// Subtracts withdrawal from balance
			// This change persists for future loop iterations
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			// Confirmation message with new balance
			
		} else {
			// Handles option 4 (Exit) or any invalid choice
			
			fmt.Println("Goodbye!")
			// Displays exit message
			// Loop continues to next iteration (doesn't exit program)
			// User will see the menu again unless they hit iteration 200
		}
		
		// After this closing brace, loop goes back to the top
		// i is incremented (i++), condition is checked (i < 200)
		// If i < 200, the loop body runs again
		// Balance changes are preserved across iterations
		
	} // End of for loop
	
	// This line is reached ONLY if:
	// 1. Loop completes all 200 iterations, OR
	// 2. Never reached if 'return' was called earlier
	
} // End of main function



// --------
// For Loop Structure:
// for initialization; condition; post {}
// Initialization (i := 0): Runs once before loop starts
// Condition (i < 200): Checked before each iteration; loop continues if true
// Post statement (i++): Runs after each iteration
// Loop body: Code inside {} that repeats
// Loop Counter:
// i starts at 0
// Increments by 1 each iteration (i++)
// Loop runs while i < 200 (iterations 0 through 199 = 200 total)
// Variable Scope in Loops:

// accountBalance: Declared OUTSIDE loop - persists across all iterations
// choice, depositAmount, withdrawalAmount: Declared INSIDE loop - recreated each iteration
// Balance changes accumulate (deposit 100 twice = balance increases by 200 total)
