package main

import (
	"fmt"
)

//  enum = enumeration -> felsorolás
//  olyan típus, aminek az értékei csak egy előre megadott fix készletből jöhetnek

//  Példa a való életből:
// 				- hét napja (Monday, Tuesday, Wednesday,...)
// 				- irányok (North, East, South, West)
// 				- Rendelés állapota (Pending, Approved, Rejected)





// létrehozunk egy új típust, ami úgy viselkedik, mint az int, de külön néven: Direction
type Direction int
// ez azért kell mert ha létrehozunk egy változót 'var d Direction'-ként akkor az csak olyan típusú értéket kaphat, amit a Direction típus megenged

//  az iota egy számláló amit a const blokkon belől a Go automatikusan növel (elso helyen iota=0)
const (
	North Direction = iota 	// 0
	East					// 1
	South					// 2
	West					// 3
)


//  String() metódus -> ezt adja meg a "nevet"
func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println(North)
	fmt.Println(East)
	fmt.Println(South)
	fmt.Println(West)

	// sliceban
	directions := []Direction{North, East, South, West}
	for _, dir := range directions {
		fmt.Println("Directions: ", dir)
	}

}
