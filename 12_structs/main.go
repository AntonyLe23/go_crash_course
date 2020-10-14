package main

import (
	"fmt"
	"strconv"
)

//Define Person struct
type Person struct {
	// firstname string
	// lastname  string
	// city      string
	// gender    string
	// age       int

	firstname, lastname, city, gender string
	age                               int
}

//Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello my name is " + p.firstname + " from " + p.city + ", age " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastname string) {
	if p.gender == "m" {
		return
	} else {
		p.lastname = spouseLastname
	}
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
	person1.hasBirthday()
	person1.getMarried(person2.lastname)
	person2.getMarried(person1.lastname)
	fmt.Println(person1.greet())
	fmt.Println(person1)
	fmt.Println(person2)
}
