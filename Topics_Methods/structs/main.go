package main

import (
	"encoding/json"
	"fmt"
)
// https://www.youtube.com/watch?v=c8H0w4yBL10
// Structs - group related data -> collection of fields

type Employee struct {
	Name		string	`json:"name"`
	Age			int		`json:"age"`
	IsRemote 	bool	`json:"remote"`
	Address
}

type Address struct {
	Street		string	`json:"myStreet"`
	City		string	`json:"myCity"`
}


func (e *Employee) updateName(newName string) {
	e.Name = newName
}

func (a Address) printaddress() {
	fmt.Println(a)
}


func main() {
	address := Address {
		Street: "123 Main Street",
		City: "New York",
	}
	employee1 := Employee{
		Name: "Alice",
		Age: 30,
		IsRemote: true,
		Address: address,
	}
	employee1.printaddress()
	employee1.updateName("Bob")


	fmt.Println("Employee Name: ", employee1.Name)
	fmt.Println("Employee Age:", employee1.Age)
	fmt.Println("Address: ", employee1.Street + employee1.City)

// Anonimus structs - 1 time datetype, only inside a function, doesn't need a name
	job := struct {
		title string
		salary int
	} {
		title: "Software developer",
		salary: 600,
	}

	fmt.Println("Job title: ", job.title)
	fmt.Println("Job salary: ", job.salary)

	// IMPORTANT - POINTERS
	// We need pointers to mutate fields of a struct inside of a struct's function for eg.

	employeePtr := &employee1
	employeePtr.Age = 31


	jsonData, _ := json.Marshal(employee1)
	fmt.Println(string(jsonData))



}