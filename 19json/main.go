package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name       string `json:"Your Name"`
	Class      string
	Rollno     int
	CourseName string
	Password   string   `json:"-"`
	Tags       []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json lecture !!!")
	//Encodejson()
	DecodeJson()
}

func Encodejson() {

	studentsData := []course{
		{"Dev Sherkhane", "FinalYear", 48, "AIDS", "abc123", []string{"Android Dev", "Kotlin"}},
		{"Omkar Phaphale", "ThirdYear", 56, "ENTC", "abc123", []string{"Electronics", "Matlab"}},
		{"Atharava Sonawane", "FinalYear", 50, "AIDS", "abc123", []string{"Web Dev", "Full Stack"}},
	}

	finalJson, err := json.MarshalIndent(studentsData, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Json :%s", finalJson)

}

func DecodeJson() {
	jsonDatsFromWeb := []byte(`{
	"Your Name": "Atharava Sonawane",
                "Class": "FinalYear",
                "Rollno": 50,
                "CourseName": "AIDS",
                "tags": ["Web Dev","Full Stack"]
}	
`)

	var lcoCourse course
	checkValid := json.Valid(jsonDatsFromWeb)
	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDatsFromWeb, &lcoCourse)
		fmt.Printf("%v\n", lcoCourse)
	} else {
		fmt.Println("Json Was not valid")
	}

	//some cases where you need to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDatsFromWeb, &myOnlineData)
	fmt.Printf("%v#\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("The key is %v and the value is %v and the thype is %T\n", k, v, v)
	}
}
