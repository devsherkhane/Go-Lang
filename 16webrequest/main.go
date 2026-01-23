package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://quotes.toscrape.com/"

func main() {
	fmt.Println("Welcome To Web request !!")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The type os response is : %T", response)

	defer response.Body.Close() //Always Close the Coneection

	databytes , err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("The Data on the web is : ",content)

}
