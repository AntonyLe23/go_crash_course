package main

import "fmt"

func main() {
	//define a map
	// emails := make(map[string]string)

	// //Assign key-value
	// emails["Phat"] = "phat@mail.com"
	// emails["Thanh"] = "thanh@mail.com"
	// emails["Thanh2"] = "thanh2@mail.com"

	//Declare map and add key-value
	emails := map[string]string{
		"Phat01":   "phat@mail01.com",
		"Thanh001": "Thanh001@mail.com",
	}

	emails["adsad"] = "dwdqd"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Phat"])

	//Delete from map
	delete(emails, "Thanh23")
	fmt.Println(emails)
}
