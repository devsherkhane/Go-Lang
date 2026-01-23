package main

import "fmt"

func main() {
	courses := []string{"Reactjs", "Nodejs", "Python", "Kotlin", "java"}
	fmt.Println(courses)
	var index int = 2 

	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println("Courses :",courses)
}
