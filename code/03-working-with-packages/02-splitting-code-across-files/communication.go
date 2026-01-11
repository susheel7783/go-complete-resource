
// by this we can call our bank.go file and read each functions
// behind the scene both the files connected

package main

import "fmt"

func presentOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

// --and this function will be part of main function
