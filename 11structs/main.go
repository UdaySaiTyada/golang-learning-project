package main

import "fmt"

func main() {
	uday := Pilot{"1", "Uday Sai Tyada", "uday@example.com", "password"}
	fmt.Println(uday)

	fmt.Printf("Details: %v \n", uday.Name)
}

type Pilot struct {
	Id       string
	Name     string
	Email    string
	Password string
}
