package main 

import "fmt" // Imports the fmt package for formatted I/O operations

func main() { // The main function - program execution starts here
	
	// Declares and initializes account balance variable
	var accountBalance = 1000.0
	// Starting balance of 1000.0 (float64 type)
	// Declared OUTSIDE the loop so it persists across all iterations
	
	// Prints welcome message once (before loop starts)
	fmt.Println("Welcome to Go Bank!")
	// Only prints ONCE when program starts, not on every loop iteration
	
	for { // Infinite loop (no initialization, condition, or post statement)
		// for { } creates a loop that runs forever
		// No counter variable (no i)
		// No exit condition (like i < 200)
		// Will run continuously until broken by 'break' or 'return'
		// This is the standard pattern for menu-driven programs
		
		// Displays the menu options (printed on EVERY loop iteration)
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		
		// Declares a variable to store the user's menu choice
		var choice int
		// Recreated each loop iteration (local to loop body)
		
		// Prompts user to enter their choice
		fmt.Print("Your choice: ")
		
		// Reads the user's input
		fmt.Scan(&choice)
		
		// COMMENTED OUT: Alternative boolean approach
		// wantsCheckBalance := choice == 1
		
		// Option 1: Check balance
		if choice == 1 {
			
			fmt.Println("Your balance is", accountBalance)
			// Displays current balance
			// After this, loop continues to next iteration (shows menu again)
			
		} else if choice == 2 {
			// Option 2: Deposit money
			
			fmt.Print("Your deposit: ")
			
			var depositAmount float64
			// Variable to store deposit amount
			
			fmt.Scan(&depositAmount)
			// Reads deposit amount from user
			
			// Validation: checks if deposit is positive
			if depositAmount <= 0 {
				
				fmt.Println("Invalid amount. Must be greater than 0.")
				// Error message for invalid deposit
				
				// return (COMMENTED OUT)
				// return would EXIT the entire program
				
				continue
				// continue: skips the rest of THIS iteration and starts next iteration
				// Jumps back to the top of the loop (shows menu again)
				// Does NOT exit the program or loop
				// Does NOT update the balance
			}
			
			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			// Adds deposit to balance (only if validation passed)
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			// Confirmation message
			// After this, loop continues to next iteration
			
		} else if choice == 3 {
			// Option 3: Withdraw money
			
			fmt.Print("Withdrawal amount: ")
			
			var withdrawalAmount float64
			// Variable to store withdrawal amount
			
			fmt.Scan(&withdrawalAmount)
			// Reads withdrawal amount
			
			// First validation: amount must be positive
			if withdrawalAmount <= 0 {
				
				fmt.Println("Invalid amount. Must be greater than 0.")
				
				return
				// return: EXITS the entire main() function
				// Program terminates immediately
				// Loop stops, "Thanks for choosing our bank" is NOT printed
				// This is inconsistent with the deposit behavior (which uses continue)
			}
			
			// Second validation: sufficient balance check
			if withdrawalAmount > accountBalance {
				
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				
				return
				// return: EXITS the entire program
				// Terminates immediately on insufficient funds
				// This is also inconsistent (could use continue instead)
			}
			
			accountBalance -= withdrawalAmount // accountBalance = accountBalance - withdrawalAmount
			// Subtracts withdrawal from balance
			
			fmt.Println("Balance updated! New amount:", accountBalance)
			// Confirmation message
			// After this, loop continues to next iteration
			
		} else {
			// Option 4: Exit (or any invalid choice like 5, -1, etc.)
			
			fmt.Println("Goodbye!")
			// Exit message
			
			// return (COMMENTED OUT)
			// return would exit the program but skip "Thanks for choosing our bank"
			
			break
			// break: exits the loop (but NOT the function)
			// Jumps to the first line AFTER the loop
			// Continues to line 109: "Thanks for choosing our bank"
			// This is the proper way to exit a loop gracefully
		}
		
		// After each iteration (unless break or return was called):
		// Loop goes back to the top (line 14: for {)
		// Menu is displayed again, waiting for next user choice
		
	} // End of infinite for loop
	
	// This line executes ONLY if 'break' was called (choice 4 or invalid)
	fmt.Println("Thanks for choosing our bank")
	// Prints final goodbye message
	// NOT printed if 'return' was called earlier (invalid deposit/withdrawal)
	
} 


-------------------
// Infinite Loop:

// for { }: Loop with no condition - runs forever
// Must use break or return to exit
// Perfect for menu-driven programs (keep showing menu until user exits)


// Loop Control Statements:

// continue: Skips rest of current iteration, starts next iteration immediately

// Used in deposit validation
// Shows menu again without updating balance


// break: Exits the loop completely, continues to code after loop

// Used in else block (exit option)
// Allows "Thanks for choosing our bank" to print


// return: Exits the entire function (and program)

// Used in withdrawal validations
// Skips "Thanks for choosing our bank" message
