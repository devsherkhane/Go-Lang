package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	languages := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many languages do you want to enter? : ")
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Print("\nEnter language name: ")
		key, _ := reader.ReadString('\n')
		key = strings.TrimSpace(key)

		fmt.Print("Enter description: ")
		value, _ := reader.ReadString('\n')
		value = strings.TrimSpace(value)

		languages[key] = value
	}

	fmt.Println("\nStored Languages:")
	for key, value := range languages {
		fmt.Printf("Language: %s | Description: %s\n", key, value)
	}
}
