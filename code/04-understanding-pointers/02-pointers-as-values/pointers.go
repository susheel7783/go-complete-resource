package main // Declares this file is part of the main package (entry point for executable)

import "fmt" // Imports fmt package for printing output

func main() { // The main function - program execution starts here
	
	age := 32 // Regular variable declaration and initialization
	// age: variable name
	// := shorthand declaration (type inferred as int from value 32)
	// 32: the actual value stored in memory
	// This creates a variable that holds the VALUE 32
	
	var agePointer *int
	// Declares a POINTER variable
	// var: explicit variable declaration
	// agePointer: variable name (conventionally ends with "Pointer")
	// *int: type is "pointer to an int"
	//   - * means "pointer to"
	//   - int is the type it points to
	// Initially nil (points to nothing) until assigned
	
	agePointer = &age
	// & is the "address-of" operator
	// &age: gets the MEMORY ADDRESS of the age variable
	// Example: if age is stored at memory location 0x1040a124
	//          then &age returns 0x1040a124
	// agePointer now stores this memory address (not the value 32)
	// agePointer "points to" the age variable
	
	fmt.Println("Age:", *agePointer)
	// * is the "dereference" operator (when used with a pointer variable)
	// *agePointer: "go to the address stored in agePointer and get the VALUE"
	// This accesses the value at the memory location agePointer points to
	// Since agePointer points to age, *agePointer gives us 32
	// Output: "Age: 32"
	
	// adultYears := getAdultYears(age) (COMMENTED OUT)
	// Would call the function, passing age BY VALUE (copy of 32)
	// Returns age - 18 = 32 - 18 = 14
	
	// fmt.Println(adultYears) (COMMENTED OUT)
	// Would print: 14
}

// Function that calculates adult years (years over 18)
func getAdultYears(age int) int {
	// Parameter: age int - receives a COPY of the value passed in
	// This is "pass by value" - changes to age here don't affect original
	// Returns: int - the calculated result
	
	return age - 18
	// Subtracts 18 from age and returns the result
	// Example: if age is 32, returns 14
	// Example: if age is 20, returns 2
}

// ----------------
// age := 32              // Regular variable (holds VALUE)
// var ptr *int           // Pointer declaration (holds ADDRESS)
// ptr = &age             // & = get address of age
// value := *ptr          // * = get value at address (dereference)
// *ptr = 40              // Change value through pointer
