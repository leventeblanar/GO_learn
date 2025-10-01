package main

import "fmt"

// Az interface Go-ban egy olyan típus, ami megmondja:
// “ha egy struct tudja ezeket a methodokat, akkor ő az interface-nek a példánya”.
// - Nincs explicit “implements” kulcsszó.
// - Ha a struct megvalósítja az interface által előírt methodokat, akkor automatikusan “implementálja” az interface-t.
// - Ez adja a Go egyszerű, de erős polimorfizmusát.


// interface definíció
type Shape interface {
	Area() 	float64
}

type Square struct {
	Side 	float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}


func main() {
	var s Shape
	s = Square{Side:4}

	fmt.Println("Area:", s.Area())
}