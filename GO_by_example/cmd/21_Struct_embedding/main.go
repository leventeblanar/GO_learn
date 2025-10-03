package main

import (
	"fmt"
)

// Ahelyett, hogy klasszikus öröklődést használnánk (mint Java/C#), Go-ban egyszerűen beágyazol egy structot egy másikba.
// Így a beágyazott mezők és metódusok úgy érhetők el, mintha a külső struct sajátjai lennének.

type Person struct {
	Name		string
	Age			int
}

type Employee struct {
	Person		// emedding
	Company		string
	EmployeeId 	int
}

func main() {
	// employee példány
	e := Employee{
		Person: Person{Name: "Alice", Age:30},
		Company: "TechCorp",
		EmployeeId: 123,
	}

	// Person mezők közvetlenül elérhetők
	fmt.Println(e.Name)
	fmt.Println(e.Age)

	// Saját mezők is működnek
	fmt.Println(e.Company)
	fmt.Println(e.EmployeeId)
}
