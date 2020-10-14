package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	ids := []int{1, 212, 12, 321, 33}

	//Loops through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	fmt.Println("=======")

	//Not using index
	for id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	fmt.Println("=======")
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//Range with map
	emails := map[string]string{
		"Phat01":   "phat@mail01.com",
		"Thanh001": "Thanh001@mail.com",
	}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}

	for _, v := range emails {
		fmt.Printf("Email: %s\n", v)
	}
}
