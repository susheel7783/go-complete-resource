package main // Declares this file is part of the main package (entry point for executable)

import "fmt" // Imports fmt package for printing output

func main() { // The main function - program execution starts here
	
	age := 32 // Regular variable declaration and initialization
	// age: variable name
	// := shorthand declaration (type inferred as int from value 32)
	// Stores the VALUE 32 in memory
	
	var agePointer *int
	// Declares a POINTER variable
	// *int: type is "pointer to an int" (stores memory addresses, not values)
	// Initially nil (points to nothing)
	
	agePointer = &age
	// & is the "address-of" operator
	// &age: gets the MEMORY ADDRESS where age is stored
	// agePointer now holds the address of age (not the value 32)
	// Example: agePointer might contain 0x1040a120
	
	fmt.Println("Age:", *agePointer)
	// * is the "dereference" operator
	// *agePointer: "follow the pointer and get the VALUE at that address"
	// Since agePointer points to age, this accesses age's value (32)
	// Output: "Age: 32"
	
	adultYears := getAdultYears(agePointer)
	// Calls getAdultYears function
	// Passes agePointer (the MEMORY ADDRESS, not the value)
	// This is "pass by reference" - function gets the address
	// adultYears will store the returned result (14)
	
	fmt.Println(adultYears)
	// Prints the result
	// Output: "14"
}

// Function that calculates adult years using a POINTER parameter
func getAdultYears(age *int) int {
	// Parameter: age *int - receives a POINTER to an int
	// age parameter holds a memory ADDRESS (not a value)
	// This is the SAME address as agePointer from main
	// Returns: int - the calculated result
	
	return *age - 18
	// *age: dereferences the pointer to get the actual VALUE
	// "Follow the pointer to get the number, then subtract 18"
	// Example: if pointer points to 32, returns 32 - 18 = 14
	// Must use * to access the value since age is a pointer
}



----------------explanation of code------
// In main():
age := 32                           // Create value 32
agePointer = &age                   // Get its address (e.g., 0x1040a120)
adultYears := getAdultYears(agePointer)  // Pass the ADDRESS

// In getAdultYears():
func getAdultYears(age *int) int {  // Receives the ADDRESS
    return *age - 18                // Dereference to get value (32), subtract 18
}
```

---

### **2. Memory Visualization:**
```
┌─────────────────────────────────────────────────────┐
│ main() function                                     │
├─────────────────────────────────────────────────────┤
│ Memory Address    Variable        Value             │
│ 0x1040a120       age             32                 │
│ 0x1040a124       agePointer      0x1040a120 ──┐     │
└──────────────────────────────────────────────┼─────┘
                                               │
        ┌──────────────────────────────────────┘
        │
        ↓
┌─────────────────────────────────────────────────────┐
│ getAdultYears() function                            │
├─────────────────────────────────────────────────────┤
│ Parameter: age (pointer)                            │
│ Value: 0x1040a120 (same address!)                  │
│                                                     │
│ *age → follows pointer → gets 32                    │
│ Returns: 32 - 18 = 14                               │
└─────────────────────────────────────────────────────┘
