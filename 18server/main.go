package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Get lecture !!")
	//PerformGetRequest()
	//PerformPostRequest()
	SendDataToForm()
}
func PerformGetRequest() {
	const myurl = "http://localhost:6000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Code Lenght is : ", response.ContentLength)

	var newString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	byteCount, err := newString.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Byte Count : ", byteCount)
	fmt.Println(newString.String())
}

func PerformPostRequest() {
	const myurl = "http://localhost:6000/post"

	requestBody := strings.NewReader(`
	   {
		"Name"   :"Dev Sherkhane",
		"Class"  :"Final year Btech",
		"Roll no":"48"
	   }
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content : ", string(content))

}

func SendDataToForm() {
	const myurl = "http://localhost:6000/postform"

	data := url.Values{}

	data.Add("First Name ", "Dev")
	data.Add("Last  Name ", "Sherkhane")
	data.Add("Roll no    ", "48")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

}
