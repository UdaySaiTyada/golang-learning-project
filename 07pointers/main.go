package main

import "fmt"

func main() {
	fmt.Println("Welcome to chapter on Go pointers")

	var actualValue int = 2

	var ptr *int = &actualValue

	// OTHER WAYS TO DEFINE POINTERS
	// var ptr = &actualValue
	// ptr := &actualValue

	fmt.Println("the address of the actual value is ", ptr)
	fmt.Println("the Value of the actual value through the pointer is ", *ptr)

	// Changing the actual value by manipulating the pointers
	*ptr = *ptr + 2
	fmt.Println("the new Value of the actual value is ", actualValue)

}
