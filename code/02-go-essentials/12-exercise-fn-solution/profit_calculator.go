package main // Declares this file is part of the main package (entry point for executable programs)

import "fmt" // Imports the fmt package for formatted I/O operations (single import, no parentheses needed)

func main() { // The main function - program execution starts here
	
	// Gets revenue input from user by calling getUserInput function
	revenue := getUserInput("Revenue: ")
	// Calls getUserInput with the prompt text, returns a float64 value stored in revenue
	
	// Gets expenses input from user
	expenses := getUserInput("Expenses: ")
	// Reuses the getUserInput function with different prompt text
	
	// Gets tax rate input from user
	taxRate := getUserInput("Tax Rate: ")
	// Tax rate should be entered as a percentage (e.g., 20 for 20%)
	
	// Calls calculateFinancials and receives THREE return values
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	// ebt: Earnings Before Tax (revenue minus expenses)
	// profit: Net profit after tax
	// ratio: Ratio of EBT to profit
	
	// Prints EBT with 1 decimal place
	fmt.Printf("%.1f\n", ebt)
	// Printf: formatted print (P = formatted)
	// %.1f: format as floating-point with 1 decimal place
	// \n: adds a new line after the output
	
	// Prints profit with 1 decimal place
	fmt.Printf("%.1f\n", profit)
	// Same formatting as above
	
	// Prints ratio with 3 decimal places
	fmt.Printf("%.3f\n", ratio)
	// %.3f: format as floating-point with 3 decimal places (more precision for ratios)
}

// Function to calculate financial metrics - takes 3 parameters, returns 3 values
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	// Parameters: revenue, expenses, taxRate are all float64 (shorthand notation)
	// Returns: three float64 values (ebt, profit, ratio)
	
	// Calculates Earnings Before Tax
	ebt := revenue - expenses
	// Simple subtraction: total revenue minus total expenses
	
	// Calculates net profit after tax
	profit := ebt * (1 - taxRate/100)
	// taxRate/100 converts percentage to decimal (e.g., 20% becomes 0.20)
	// (1 - taxRate/100) gives the remaining percentage after tax (e.g., 1 - 0.20 = 0.80)
	// Multiply EBT by this to get profit after tax
	
	// Calculates the ratio of EBT to profit
	ratio := ebt / profit
	// Shows how EBT compares to final profit
	// This ratio indicates the tax impact (higher ratio = more tax paid)
	
	return ebt, profit, ratio // Returns all three calculated values to the caller
}

// Reusable function to get user input with a custom prompt
func getUserInput(infoText string) float64 {
	// Parameter: infoText is the prompt message to display
	// Returns: a single float64 value (the user's input)
	
	var userInput float64 // Declares a variable to store the user's input
	
	fmt.Print(infoText) // Displays the prompt text (no newline after)
	// Example: displays "Revenue: " and waits for input on the same line
	
	fmt.Scan(&userInput) // Reads user input from console and stores it in userInput
	// & (address-of operator) allows Scan to modify the userInput variable
	
	return userInput // Returns the value entered by the user
	// This value is sent back to wherever the function was called
}

---------------
// Key Concepts in This Code:

// DRY Principle: "Don't Repeat Yourself" - getUserInput() is reused 3 times instead of repeating the same code
// Multiple Returns: calculateFinancials() returns 3 values at once, making the code efficient
// Function Composition: main() calls helper functions to organize the program logically
// Printf vs Print: Printf allows formatted output with placeholders like %.1f, while Print outputs as-is
