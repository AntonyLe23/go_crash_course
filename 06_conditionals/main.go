package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	x, y := 10, 5
	if x <= y {
		fmt.Printf("%d is less or equal than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	color := "green"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue or red")
	}

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not red or blue")
	}
}
