package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to Switch Case !!")

	rand.Seed(time.Now().UnixNano())
	dicenumber:= rand.Intn(6)+1
	fmt.Println("Dice number",dicenumber)

	switch dicenumber {
	case 1 : 
		fmt.Println("You get 1")
	case 2 :
		fmt.Println("You get 2")
	case 3 : 
		fmt.Println("You get 3")
	case 4 :
		fmt.Println("You get 4")
	case 5 :
		fmt.Println("You get 5")
	case 6 : 
		fmt.Println("You get 6")
	default :
		fmt.Println("Watch Out !! ")
	}
}
