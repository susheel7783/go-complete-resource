package main

import "fmt" // Imports the fmt package for formatted I/O operations

func main() { // The main function - program execution starts here
	
	// Declares and initializes account balance variable
	var accountBalance = 1000.0
	// var keyword: explicitly declares a variable
	// = 1000.0: initializes with a float64 value (decimal number)
	// Go infers the type as float64 because of the decimal point
	
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
	// int type: stores whole numbers (integers) like 1, 2, 3, 4
	
	// Prompts user to enter their choice
	fmt.Print("Your choice: ")
	// Print: keeps cursor on same line for user input
	
	// Reads the user's input and stores it in the choice variable
	fmt.Scan(&choice)
	// &choice: passes the memory address so Scan can modify the value
	
	// COMMENTED OUT: Alternative way to check if user chose option 1
	// wantsCheckBalance := choice == 1
	// This would create a boolean variable (true/false) based on the comparison
	
	// If statement: checks if user chose option 1 (Check balance)
	if choice == 1 {
		// == is the equality comparison operator (checks if values are equal)
		// This block executes ONLY if choice equals 1
		
		fmt.Println("Your balance is", accountBalance)
		// Displays the current account balance
		// Println automatically adds a space between "is" and the number
		
	} else if choice == 2 {
		// else if: checks this condition ONLY if the previous if was false
		// This block executes if choice equals 2 (Deposit money)
		
		fmt.Print("Your deposit: ")
		// Prompts user to enter deposit amount (cursor stays on same line)
		
		var depositAmount float64
		// Declares a variable to store the deposit amount
		// float64: can hold decimal numbers (like 50.75)
		
		fmt.Scan(&depositAmount)
		// Reads the deposit amount from user input
		
		accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
		// += is the compound assignment operator (add and assign)
		// Takes current accountBalance, adds depositAmount, stores result back in accountBalance
		// Example: if balance is 1000 and deposit is 500, new balance becomes 1500
		
		fmt.Println("Balance updated! New amount:", accountBalance)
		// Confirms the deposit and shows the new balance
	}
	// Note: Currently no handling for choice 3 (withdraw) or choice 4 (exit)
	// Also no handling for invalid choices (like 5, -1, etc.)
	
	// Prints the user's choice (this runs regardless of which option was chosen)
	fmt.Println("Your choice:", choice)
	// This line executes after the if-else block completes
}
