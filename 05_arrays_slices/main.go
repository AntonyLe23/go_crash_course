package main

import "fmt"

func main() {
	//Arrays
	// var fruitArr [2]string

	//Assigns arr values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Lemon"

	//Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Slices
	fruitSlice := []string{"Apple", "Lemon", "Banana", "Graph"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:1])
}
