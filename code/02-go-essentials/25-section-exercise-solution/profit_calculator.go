package main // Declares this file is part of the main package (entry point for executable programs)

import ( // Imports multiple packages
	"errors" // Package for creating custom error messages
	"fmt"    // Package for formatted I/O operations
	"os"     // Package for operating system functions (file operations)
)

// Goals (comments describing the program's objectives)
// 1) Validate user input
//    => Show error message & exit if invalid input is provided
//    - No negative numbers
//    - Not 0
// 2) Store calculated results into file

func main() { // The main function - program execution starts here
	
	// Gets revenue input with validation
	revenue, err := getUserInput("Revenue: ")
	// getUserInput returns (float64, error)
	// revenue: the input value if valid
	// err: error if validation fails, nil if successful
	
	if err != nil {
		// Checks if there was an error (invalid input)
		
		fmt.Println(err)
		// Prints the error message
		// Will show "Value must be a positive number."
		
		return
		// Exits the main function (and program) immediately
		// No further code executes
		
		// panic(err) (COMMENTED OUT)
		// panic: would crash the program with stack trace
		// return is cleaner - just exits gracefully
	}
	
	// Gets expenses input with validation
	expenses, err := getUserInput("Expenses: ")
	// Reuses the err variable (same name, new value)
	
	if err != nil {
		// Checks if expenses input was invalid
		
		fmt.Println(err)
		// Prints error message
		
		return
		// Exits program if expenses are invalid
	}
	
	// Gets tax rate input with validation
	taxRate, err := getUserInput("Tax Rate: ")
	// Again reuses err variable
	
	if err != nil {
		// Checks if tax rate input was invalid
		
		fmt.Println(err)
		// Prints error message
		
		return
		// Exits program if tax rate is invalid
	}
	
	// COMMENTED OUT: Alternative approach to check all errors at once
	// if err1 != nil || err2 != nil || err3 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// This would require different variable names (err1, err2, err3)
	// Current approach is better - fails fast on first error
	
	// Calculates financial metrics using validated inputs
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	// Returns three calculated values
	// Only reached if all inputs were valid
	
	// Prints EBT (Earnings Before Tax) with 1 decimal place
	fmt.Printf("%.1f\n", ebt)
	// %.1f: format as float with 1 decimal place
	// \n: adds new line
	
	// Prints profit with 1 decimal place
	fmt.Printf("%.1f\n", profit)
	
	// Prints ratio with 3 decimal places
	fmt.Printf("%.3f\n", ratio)
	// %.3f: format as float with 3 decimal places (more precision)
	
	// Saves the results to a file
	storeResults(ebt, profit, ratio)
	// Writes all three values to results.txt
	
} // End of main function

// Function to store financial results to a file
func storeResults(ebt, profit, ratio float64) {
	// Parameters: three float64 values to save
	// No return value (could be improved to return error)
	
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	// fmt.Sprintf: formats string WITHOUT printing it
	// Returns a formatted string instead of printing to console
	// Creates multi-line string with all three values
	// Example output:
	//   "EBT: 500.0
	//    Profit: 400.0
	//    Ratio: 1.250
	//   "
	// \n: creates new lines in the file for readability
	
	os.WriteFile("results.txt", []byte(results), 0644)
	// os.WriteFile: writes data to file
	// Parameters:
	//   1. "results.txt" - filename (created in same directory)
	//   2. []byte(results) - converts string to byte slice
	//   3. 0644 - file permissions (owner read/write, others read-only)
	// Note: Error is IGNORED (not best practice)
	// Should check: err := os.WriteFile(...); if err != nil { ... }
}

// Function to calculate financial metrics
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	// Parameters: all three are float64 (shorthand when types are same)
	// Returns: three float64 values (ebt, profit, ratio)
	
	// Calculates Earnings Before Tax
	ebt := revenue - expenses
	// Simple subtraction: total revenue minus total expenses
	
	// Calculates net profit after tax
	profit := ebt * (1 - taxRate/100)
	// taxRate/100: converts percentage to decimal (20% becomes 0.20)
	// (1 - taxRate/100): remaining percentage after tax (1 - 0.20 = 0.80)
	// Multiply EBT by remaining percentage to get profit
	
	// Calculates ratio of EBT to profit
	ratio := ebt / profit
	// Shows relationship between pre-tax and post-tax earnings
	// Higher ratio = more tax paid
	
	return ebt, profit, ratio
	// Returns all three calculated values
}

// Function to get and validate user input
func getUserInput(infoText string) (float64, error) {
	// Parameter: infoText is the prompt to display to user
	// Returns: (float64, error) - the input value and any validation error
	// This is Go's standard pattern for functions that can fail
	
	var userInput float64
	// Declares variable to store user input
	// Initialized to 0.0 by default
	
	fmt.Print(infoText)
	// Displays the prompt (e.g., "Revenue: ")
	// Print (not Println): keeps cursor on same line
	
	fmt.Scan(&userInput)
	// Reads user input from console
	// &userInput: passes memory address so Scan can modify the variable
	// Waits for user to type a number and press Enter
	
	// Validation: checks if input is valid
	if userInput <= 0 {
		// <= means "less than or equal to"
		// Catches both negative numbers AND zero
		// Both are invalid for revenue, expenses, and tax rate
		
		return 0, errors.New("Value must be a positive number.")
		// Returns TWO values:
		//   1. 0 - dummy value (not used when there's an error)
		//   2. errors.New() - creates error with custom message
		// Function exits here if validation fails
	}
	
	return userInput, nil
	// SUCCESS case: returns the input and nil (no error)
	// nil means "no error occurred"
	// Only reached if userInput > 0
}
