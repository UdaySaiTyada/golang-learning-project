package main

import "fmt"

const LoginToken string = "wesrxdctfvygbunimop,hgytfrde"

func main() {
	var username string = "uday"
	fmt.Println(username)
	fmt.Printf("varaiable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("varaiable is of type: %T \n", isLoggedIn)

	var smallValue int = 255
	fmt.Println(smallValue)
	fmt.Printf("varaiable is of type: %T \n", smallValue)

	// default values

	var anotherint int
	fmt.Println(anotherint)
	fmt.Printf("varaiable is of type: %T \n", anotherint)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("varaiable is of type: %T \n", anotherString)

	// implecet type

	var password = "password"
	fmt.Println(password)

	// no var style

	// walrus operator is not allowed out side the function
	numberOfUsers := 30000
	fmt.Println(numberOfUsers)

	// Printing a public variable
	fmt.Println(LoginToken)
}
