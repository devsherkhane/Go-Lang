package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to user ip !!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating of Our Pizza :")

	in,_ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating !!",in)

	numRating,err := strconv.ParseFloat(strings.TrimSpace(in),64)

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 in rating :", numRating+1)
	}
}
