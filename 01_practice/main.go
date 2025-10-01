package main

import (
	"fmt"
)

type CanSpeak interface {
	Speak()		string
}

type Dog struct {
	Name		string
}

type Cat struct {
	Name		string
}

func (d Dog) Speak() string {
	return d.Name + " says: Woof!"
}

func (c Cat) Speak() string {
	return c.Name + " says: Meow!"
}

func main() {
	animals := []CanSpeak{
		Dog{Name: "Buddy"},
		Cat{Name: "Mittens"},
	}

	for _, a := range animals {
		fmt.Println(a.Speak())
	}
}