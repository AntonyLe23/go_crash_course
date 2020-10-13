package main

import "fmt"

func main() {
	//Using var
	// var name = "Phat"
	var age int32 = 23
	const isCool = true

	//Shorthand
	name, email := "Phat", "phat@mail.com"
	size := 1.3
	var otherSize float32 = 2.3

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", otherSize)
}
