package main

import "fmt"

func main() {
	fmt.Println(greeting("Phat"))
	fmt.Println(getSum(17, 23))
}

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}
