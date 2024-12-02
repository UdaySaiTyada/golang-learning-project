package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza store")
	fmt.Println("submit rating: ")

	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')

	fmt.Printf("Thank you for rating it %v", rating)

	numberRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Thank you for rating it %v", numberRating)

}
