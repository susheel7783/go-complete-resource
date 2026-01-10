package main // Declares this file is part of the main package (entry point for executable programs)

import ( // Imports external packages needed for the program
	"fmt"  // Package for formatted I/O (input/output) operations like printing and scanning
	"math" // Package for mathematical functions like Pow (power/exponentiation)
)

const inflationRate = 2.5 // Global constant - can be accessed by any function in this file

func main() { // The main function - program execution starts here
	var investmentAmount float64 // Declares a variable to store the investment amount (decimal number)
	var years float64            // Declares a variable to store the number of years (decimal number)
	
	expectedReturnRate := 5.5 // Declares and initializes expected return rate to 5.5% (shorthand syntax)
	
	// Prompts user to enter the investment amount
	outputText("Investment Amount: ") // Calls our custom function to display text
	fmt.Scan(&investmentAmount)       // Reads user input and stores it in investmentAmount variable (& means "address of")
	
	// Prompts user to enter the expected return rate
	outputText("Expected Return Rate: ") // Displays prompt for return rate
	fmt.Scan(&expectedReturnRate)        // Reads user input and updates expectedReturnRate variable
	
	// Prompts user to enter the number of years
	outputText("Years: ")     // Displays prompt for years
	fmt.Scan(&years)          // Reads user input and stores it in years variable
	
	// Calls calculateFutureValues function and receives TWO return values
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue: stores the first returned value (fv from the function)
	// futureRealValue: stores the second returned value (rfv from the function)
	
	// Creates formatted string with 1 decimal place for future value
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// %.1f formats the number with 1 decimal place, \n adds a new line
	
	// Creates formatted string with 1 decimal place for inflation-adjusted value
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	
	// Prints both formatted strings to the console
	fmt.Print(formattedFV, formattedRFV) // Outputs the final results to the user
}

// Custom function to output text to the console
func outputText(text string) { // Takes a string parameter called "text"
	fmt.Print(text) // Prints the text without adding a new line after it
}

// Function to calculate both future values using NAMED RETURN VALUES
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	// Parameters: all three parameters are float64 (shorthand when types are the same)
	// Named return values: (fv float64, rfv float64) - these variables are automatically declared
	// fv and rfv are already created and initialized to zero when the function starts
	
	// Calculates nominal future value using compound interest formula
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// Assigns the calculated value to the named return variable fv
	// No need to use := because fv is already declared in the function signature
	
	// Calculates real future value (adjusted for inflation)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// Assigns the calculated value to the named return variable rfv
	// Uses the global constant inflationRate (accessible from anywhere in the file)
	
	return fv, rfv // Explicitly returns both named variables
	// With named return values, you could also just write "return" (naked return)
	// The commented "return" below would work the same way
	// return
}



-------
// Named Returns: (fv float64, rfv float64) declares the return variables in the function signature
// Auto-declaration: fv and rfv are automatically created when the function starts (initialized to zero)
// Use = not :=  Since they're already declared, you use = to assign values, not :=
// Explicit vs Naked Return:

// return fv, rfv - explicitly returns the variables (clearer, recommended)
// return - "naked return" that automatically returns all named variables (works but less clear)


// Benefit: Makes the code slightly more readable by showing what the function returns right in the signature

// Best Practice: Named returns are useful for documentation, but explicit returns (return fv, rfv) are generally preferred over naked returns for clarity, especially in longer functions.
