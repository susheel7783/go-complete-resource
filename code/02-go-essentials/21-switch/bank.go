package main // Declares this file is part of the main package (entry point for executable programs)

import "fmt" // Imports the fmt package for formatted I/O operations

func main() { // The main function - program execution starts here
	
	// Declares and initializes account balance variable
	var accountBalance = 1000.0
	// Starting balance of 1000.0 (float64 type)
	// Persists across all loop iterations
	
	// Prints welcome message once when program starts
	fmt.Println("Welcome to Go Bank!")
	
	for { // Infinite loop - runs until explicitly exited
		// No condition means loop runs forever
		// Must use 'return' or 'break' to exit
		
		// Displays menu options on every iteration
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
		
		// COMMENTED OUT: Alternative approach using boolean
		// wantsCheckBalance := choice == 1
		
		// Switch statement - cleaner alternative to multiple if-else if chains
		switch choice {
		// switch: evaluates 'choice' once and compares it to each case
		// More readable than if-else when checking one variable against multiple values
		// Only ONE case will execute (no fall-through in Go by default)
		
		case 1: // Executes if choice == 1
			// case 1: equivalent to "if choice == 1"
			// No need for 'break' - Go automatically exits switch after case executes
			
			fmt.Println("Your balance is", accountBalance)
			// Displays current balance
			// After this case, switch ends and loop continues to next iteration
			
		case 2: // Executes if choice == 2 (Deposit)
			// case 2: equivalent to "else if choice == 2"
			
			fmt.Print("Your deposit: ")
			
			var depositAmount float64
			// Variable to store deposit amount
			
			fmt.Scan(&depositAmount)
			// Reads deposit amount from user
			
			// Validation: checks if deposit is positive
			if depositAmount <= 0 {
				
				fmt.Println("Invalid amount. Must be greater than 0.")
				// Error message
				
				// return (COMMENTED OUT)
				// return would exit the entire program
				
				continue
				// continue: skips rest of loop iteration, goes back to top
				// Shows menu again without updating balance
				// EXITS the switch AND skips rest of loop body
			}
			
			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			// Adds deposit to balance (only if validation passed)
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			// Confirmation message
			// After this, switch ends and loop continues
			
		case 3: // Executes if choice == 3 (Withdraw)
			// case 3: equivalent to "else if choice == 3"
			
			fmt.Print("Withdrawal amount: ")
			
			var withdrawalAmount float64
			// Variable to store withdrawal amount
			
			fmt.Scan(&withdrawalAmount)
			// Reads withdrawal amount
			
			// First validation: amount must be positive
			if withdrawalAmount <= 0 {
				
				fmt.Println("Invalid amount. Must be greater than 0.")
				
				continue
				// Skips rest of iteration, shows menu again
				// Consistent with deposit validation (uses continue, not return)
			}
			
			// Second validation: sufficient balance check
			if withdrawalAmount > accountBalance {
				
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				
				continue
				// Skips rest of iteration, shows menu again
				// Now CONSISTENT - all validations use continue
			}
			
			accountBalance -= withdrawalAmount // accountBalance = accountBalance - withdrawalAmount
			// Subtracts withdrawal from balance
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			// Confirmation message
			// After this, switch ends and loop continues
			
		default: // Executes if choice doesn't match any case
			// default: equivalent to "else" - catches everything not in cases 1, 2, or 3
			// Handles choice 4 (Exit) AND any invalid choices (5, -1, 100, etc.)
			
			fmt.Println("Goodbye!")
			// Exit message
			
			fmt.Println("Thanks for choosing our bank")
			// Final thank you message
			// Prints BEFORE exiting (because return comes after)
			
			return
			// return: exits the entire main() function
			// Program terminates here
			// Loop stops, no more iterations
			// This is a clean exit from the program
			
			// break (COMMENTED OUT)
			// break would only exit the switch, NOT the loop
			// Loop would continue running (menu shows again)
			// That's why return is used instead - to exit both switch AND loop
		}
		
		// After switch block completes (if no return/continue was called):
		// Loop goes back to top, displays menu again
		
	} // End of infinite for loop
	
	// This line is NEVER reached because:
	// - Loop is infinite
	// - Only exit is via 'return' in default case
	// - return exits the function before reaching here
	
} 

// --------------------

// Switch Statement:

// switch variable { case value: ... }: Cleaner than if-else chains
// Evaluates choice once, compares to each case
// Only ONE case executes (no fall-through by default in Go)
// No break needed after each case (automatic in Go)
// More readable when checking one variable against multiple values


// Switch vs If-Else:

// Before: if choice == 1 { } else if choice == 2 { } else if choice == 3 { } else { }
// After: switch choice { case 1: case 2: case 3: default: }
// Same logic, cleaner syntax


// Case vs Default:

// case 1:, case 2:, case 3:: Specific values to match
// default:: Catches everything else (like "else")
