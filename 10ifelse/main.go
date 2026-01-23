package main

import "fmt"

func main() {
	var num int = 15

	if num < 10 {
		fmt.Println("The number is valid ")
	}else if num >10 {
		fmt.Println("Watch out ")
	}else {
		fmt.Println("The number is exactly 10 ")
	}
}
