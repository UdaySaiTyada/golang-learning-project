package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Maps")

	languagesMap := make(map[string]string)

	languagesMap["1"] = "one"
	languagesMap["2"] = "two"
	languagesMap["3"] = "three"
	languagesMap["4"] = "four"
	languagesMap["5"] = "five"

	fmt.Println(languagesMap)

	// get value from key
	fmt.Println(languagesMap["1"])

	// Delete a key value pair
	delete(languagesMap, "1")
	fmt.Println(languagesMap)

	// Iterating through a map
	for key, value := range languagesMap {
		fmt.Printf("%v : %v \n", key, value)
	}
}
