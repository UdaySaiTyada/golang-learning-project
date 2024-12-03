package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the slices lesson")

	var fruitList = []string{"apple", "banana", "orange", "melon", "valium"}
	fmt.Printf("The type of the list is %T\n", fruitList)
	fmt.Println(fruitList)

	fmt.Println("Slicing")
	fmt.Println(fruitList[1:3])

	var newFruitList = append(fruitList, "new", "old")
	fmt.Println(fruitList)
	fmt.Println(newFruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 867
	// highScores[4] = 777
	fmt.Println(highScores)

	// EVEN THROUGH WE INITIALLY INITIATED THE SIZE TO BE 4, WE CAN APPEND
	highScores = append(highScores, 777, 567, 908)
	fmt.Println(highScores)

	// Sorting a slice
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}
