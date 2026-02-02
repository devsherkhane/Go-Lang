package main

import "fmt"

func main() {
	type Employee struct {
		name     string
		id       int
		isRemote bool
	}

	employee := Employee{
		name:     "Alice",
		id:       01,
		isRemote: true,
	}

	fmt.Println("The name of the employee is : ", employee.name)
	fmt.Println("The id of the employee is : ", employee.id)
	fmt.Println("Is employee a remote employee :", employee.isRemote)

	job := struct { //Anynomous struct
		name   string
		id     int
		salary int
	}{
		name:   "Alice",
		id:     02,
		salary: 15000,
	}

	fmt.Println("The employee name is : ", job.name)
	fmt.Println("The employee id and salary is ", job.id, " ", job.salary)
}
