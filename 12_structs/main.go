package main

import "fmt"

//Define person struct
type Person struct {
	firstname string
	lastname  string
	city      string
	gender    string
	age       int
}

func main() {
	fmt.Println("Hello world")
	//Init person using struct
	person1 := Person{
		firstname: "Phat",
		lastname:  "Le",
		city:      "HCM",
		gender:    "m",
		age:       23,
	}

	person2 := Person{"Thanh", "Ha", "HCM", "f", 23}

	fmt.Println(person1)
	fmt.Println(person2)
}
