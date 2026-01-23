package main

import "fmt"

func greeter() {
	fmt.Println("Hello ffrom function 1 ")
}
func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	total := 0

	for _, i := range values {
		total += i
	}
	return total
}

func main() {
	fmt.Println("Welcome to Functions !!! ")
	total := proAdder(23, 45)

	fmt.Println("The total is : ", total)
}
