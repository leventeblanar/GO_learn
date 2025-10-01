package main

import (
	"fmt"
)

// Go-ban funkciókat is “ráakaszthatsz” a structra → ez lesz a method.
// A method olyan, mint egy függvény, csak van egy receiver-je (fogadó).
// Ez a receiver mondja meg, melyik típushoz tartozik a method.

type Person struct {
	Name		string
	Age			int
}

// érték receiver: másolatot kap, nem módosítja az eredetit
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

//  p *Person
func (p *Person) HaveBirthday() {
	p.Age++
}


func main() {
	p := Person{Name: "Alice", Age: 30}
	p.Greet()

	p.HaveBirthday()
	fmt.Println(p.Age)
}