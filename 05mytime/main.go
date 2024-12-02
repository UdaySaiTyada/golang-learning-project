package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	yesterDay := presentTime.AddDate(1, 0, 0)
	fmt.Println(yesterDay.Format("02/01/2006 Monday 15:04:05"))

}
