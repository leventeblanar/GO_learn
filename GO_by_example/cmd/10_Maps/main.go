package main

import (
	"fmt"
)

// Maps - asszociatív tömb, kulcs-érték párok
// a kulcs típusa bármilyen összehasonlítaható típus lehet (pl. string, int) - az érték tetszőleges

func main() {
	// 1. Map létrehozása make()-kel
	m := make(map[string]int)

	// 2. Értékek hozzáadása
	m["Alice"] = 23
	m["Bob"] = 30
	m["Charlie"] = 27

	// 3. Map kiírása
	fmt.Println("Full map:", m)

	// 4. Egy kulcs lekérdezése
	fmt.Println("Alice age:", m["Alice"])

	// 5. Kulcs ellenőrzés (létezik vagy sem)
	if age, ok := m["David"]; ok {
		fmt.Println("David age:", age)
	} else {
		fmt.Println("David not found in map.")
	}

	// 6. Elem törlése
	delete(m, "Bob")
	fmt.Println("After deleting Bob:", m)

	// 7. Iterálás range segítségével
	fmt.Println("Iterating map:")
	for name, age := range m {
		fmt.Println(name, "is", age, "years old")
	}
}

func rangeIterations() {

	// Slice bejárása
	nums := []int{10, 20, 30}

	for i, v := range nums {
		fmt.Println("Index", i, "Value:", v)
	}

	// Map bejárása
	ages := map[string]int{"Alice": 25, "Bob": 30}
	for name, age := range ages {
		fmt.Println(name, "is", age, "years old.")
	}

	// String bejárása
	for i, c := range "hello" {
		fmt.Println("Index:", i, "Rune", string(c))
	}
}