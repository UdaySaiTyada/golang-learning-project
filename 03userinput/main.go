package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	name, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for coming to the site %v", name)

}
