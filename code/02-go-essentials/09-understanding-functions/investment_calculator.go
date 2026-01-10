package main // Declares this file is part of the main package (entry point for executable programs)

import ( // Imports external packages needed for the program
	"fmt"  // Package for formatted I/O (input/output) operations like printing and scanning
	"math" // Package for mathematical functions like Pow (power/exponentiation)
)

func main() { // The main function - program execution starts here
	const inflationRate = 2.5 // Declares a constant inflation rate of 2.5% (cannot be changed later)
	
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
	
	// Calculates future value using compound interest formula: FV = PV * (1 + r)^n
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// investmentAmount: initial investment (PV)
	// expectedReturnRate/100: converts percentage to decimal (e.g., 5.5% becomes 0.055)
	// math.Pow: raises the base (1 + rate) to the power of years
	
	// Calculates inflation-adjusted future value by dividing by inflation growth
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	// Divides futureValue by the inflation factor to get the "real" purchasing power
	// This shows what the money will actually be worth in today's dollars
	
	// Creates formatted strings with 1 decimal place using Sprintf (string print formatted)
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// %.1f means: format as floating-point number with 1 decimal place
	// \n adds a new line at the end
	
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	// Same formatting for the inflation-adjusted value
	
	// Prints both formatted strings to the console
	fmt.Print(formattedFV, formattedRFV) // Outputs the results to the user
}

// Custom function to output text to the console
func outputText(text string) { // Takes a string parameter called "text"
	fmt.Print(text) // Prints the text without adding a new line
}


-------------
// Key Concepts Explained:

// Variables vs Constants: var creates variables that can change, const creates values that cannot change
// := syntax: Shorthand for declaring and initializing variables in one line
// & symbol: Gets the memory address of a variable (needed for fmt.Scan to modify the variable)
// Compound Interest Formula: FV = PV × (1 + r)^n where r is the rate and n is time
// Inflation Adjustment: Divides by inflation growth to show real purchasing power
