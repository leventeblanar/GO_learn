package main

import (
	"fmt"
)

//  Típusa [N]T - a méret is a típus része ([3]int és [4]int két külön típus)
//  Fix hossz - nem nő/csökken
//  Érték-szemantika - másodolódik assignmentnél és függvényhívásnál
//  Összehasonlítható (==, !=), ha az elemei is azok (pl. [3]int{1,2,3} == [3]int{1,2,3} -> true)
//  Ritkán használod közvetlenül, inkább sliceokat.


//  Mikor jó? - ha nagyon szigorú, fix méret kell (pl. 16 bájtos azonosító), vagy alacsony szintű optimalizálásnál

func main() {
	var a [3]int  			// [0, 0, 0]
	b := [3]int{1,2,3}
	c := b					// MÁSOLAT
	c[0] = 99				// Elem Módosítás index alapján 
	fmt.Println(a, b, c)

	//  array átadása függvényeknek: másolat
	fmt.Println(sumArray(b))
}

func sumArray(x [3]int) int {
	return x[0] + x[1] + x[2]
}