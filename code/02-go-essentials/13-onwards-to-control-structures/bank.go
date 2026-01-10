package main // Declares this file is part of the main package (entry point for executable programs)

import "fmt" // Imports the fmt package for formatted I/O operations

func main() { // The main function - program execution starts here
	
	// Prints the welcome message to the console
	fmt.Println("Welcome to Go Bank!")
	// Println: Print line - prints text and automatically adds a new line at the end
	
	// Prints a prompt asking what the user wants to do
	fmt.Println("What do you want to do?")
	// Each Println call prints on a new line
	
	// Displays menu option 1
	fmt.Println("1. Check balance")
	// Shows the first action available to the user
	
	// Displays menu option 2
	fmt.Println("2. Deposit money")
	// Shows the second action available to the user
	
	// Displays menu option 3
	fmt.Println("3. Withdraw money")
	// Shows the third action available to the user
	
	// Displays menu option 4
	fmt.Println("4. Exit")
	// Shows the option to exit the program
	
	// Declares a variable to store the user's menu choice
	var choice int
	// int type: stores whole numbers (integers) like 1, 2, 3, 4
	// Initialized to 0 by default until user provides input
	
	// Prompts user to enter their choice
	fmt.Print("Your choice: ")
	// Print (without "ln"): prints text WITHOUT adding a new line
	// This keeps the cursor on the same line for user input
	
	// Reads the user's input and stores it in the choice variable
	fmt.Scan(&choice)
	// &choice: passes the memory address of choice so Scan can modify its value
	// Waits for user to type a number and press Enter
	
	// Prints back the user's choice for confirmation
	fmt.Println("Your choice:", choice)
	// Displays "Your choice: " followed by the number the user entered
	// For example, if user enters 2, it prints: "Your choice: 2"
}
