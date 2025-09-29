package main

import (
	"fmt"
)

// closure - olyan függvény, ami hozzáfér a környezete változóihoz, még akkor is, ha azok a scopeon kívülre kerülnének normális esetben
// funkciót adhatunk vissza másik funkcióból
// így pl. csinálhatunk számlálókat, state-eket, kezelő kis mini függvényeket


// MIÉRT JÓ EGY CLOSURE?
// 		- ha csinálunk egy külön függvényt az mindig újraindul, csak a paraméterből dolgozik
// 		- nincs memória, nincs állapot

// EZT ÚGY OLDJA MEG EGY CLOSURE:
// 		- amikor egy függvény visszaad egy másik függvényt, akkor a belső függvény hozzáfér a külső függvény változóihoz
// 		- ezek a változók nem tűnnek el, hanem "életrben maradnak" annyi ideig, amennyi ideig a closure él

// ELŐNYÖK
// 		- olyan, mintha egy mini objektumot hoznánk létre állapottal együtt ,anélkül, hogy osztályokat kéne definiálni
// 		- függvény, ami emlékszik a multjára
// 		- két külön closure példány -> külön állapot



// adder létrehoz egy függvényt, ami mindig hozzáad egy számot a belső statehez
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


func main() {
	// két külön "példány", saját state-el
	posSum := adder()
	negSum := adder()

	fmt.Println(posSum(1))
	fmt.Println(posSum(2))
	fmt.Println(posSum(3))

	fmt.Println(negSum(-1))
	fmt.Println(negSum(-2))
	fmt.Println(negSum(-3))
}