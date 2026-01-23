package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user ip !!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating of Our Pizza :")

	in,_ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating !!",in)
}
