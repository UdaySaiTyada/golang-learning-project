package main

import "fmt"

func main() {
	fmt.Println("Welcome to the lesson on arrays")

	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[2] = "Banana"
	fruitList[3] = "Papaya"
	fruitList[4] = "Tomato"

	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Fruit list length is : ", len(fruitList))

	var vegetableList = [3]string{"carrot", "beetroot", "cabbage"}

	fmt.Println("Vegetable list is : ", vegetableList)
	fmt.Println("Vegetable list length is : ", len(vegetableList))

	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	var stringArray = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Println(stringArray)
}
