package main

import (
	"fmt"
)

func main() {
	fmt.Println("Class on If Else")

	// reader := bufio.NewReader(os.Stdin)
	// ageStr, _ := reader.ReadString('\n')

	// var age int = strconv.ParseInt(ageStr, 10, 16)

	age := 18

	if age >= 18 {
		fmt.Println("The person is a major")
	} else {
		fmt.Println("The person is a minor")
	}
}
