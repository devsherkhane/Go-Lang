package main

import "fmt"

func main() {
	fmt.Println("Welcom to Pointers")

	var mynum = 23 

	var ptr = &mynum

	fmt.Println("The pointer is ",ptr)
	fmt.Println("The acttual value is :",*ptr)
}
