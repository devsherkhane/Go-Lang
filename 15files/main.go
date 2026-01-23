package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to File Handling")
	content := "Name : Dev Sherkhane , Class : Final Year Btech , Roll no : 48 "

	file, err := os.Create("./DevFile")
	CalculateError(err)

	length, err := io.WriteString(file, content)
	fmt.Println("Write Successfully")
	CalculateError(err)
	fmt.Println("Length is : ", length)
	defer file.Close()
	ReadFile("./DevFile")

}

func CalculateError(err error) {
	if err != nil {
		panic(err)
	}
}
func ReadFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	CalculateError(err)

	fmt.Println("The Data in the file is \n ", string(databyte))
}
