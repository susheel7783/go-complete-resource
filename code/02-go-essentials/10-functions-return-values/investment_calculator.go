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
	fmt.Print("Investment Amount: ") //instead of this we can use the call the function which we created below
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
	// futureValue: stores the first returned value (nominal future value)
	// futureRealValue: stores the second returned value (inflation-adjusted value)
	
	// Creates formatted string with 1 decimal place for future value
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// %.1f formats the number with 1 decimal place, \n adds a new line
	
	// Creates formatted string with 1 decimal place for inflation-adjusted value
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	
	// Prints both formatted strings to the console
	fmt.Print(formattedFV, formattedRFV) // Outputs the final results to the user
}

// Custom function to output text to the console
func outputText(text string) { // Takes a string parameter called "text" and you can mention multiple parameter here
	fmt.Print(text) // Prints the text without adding a new line after it
}

// Function to calculate both future values - takes 3 parameters, returns 2 values
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	// Parameter list: all three parameters are float64 (shorthand when types are the same)
	// Return types: (float64, float64) means this function returns TWO float64 values
	
	// Calculates nominal future value using compound interest formula
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// investmentAmount * (1 + rate)^years
	
	// Calculates real future value (adjusted for inflation)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	// Divides by inflation growth factor to get purchasing power in today's dollars
	// Uses the global constant inflationRate (accessible from anywhere in the file)
	
	return fv, rfv // Returns both values (separated by comma) back to the caller
	// First value goes to futureValue, second goes to futureRealValue in main()
}
