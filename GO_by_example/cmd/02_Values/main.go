package main

import "fmt"

func main() {
	name := "Go"
	age := 10
	pi := 3.14159


	fmt.Println("string:", "Go is fun")
	fmt.Println("int:", 42)
	fmt.Println("float:", 3.14)
	fmt.Println("bool:", true)

	fmt.Printf("Pi approx: %.2f\n", 3.14159)

	// ÁLTALÁNOS
	fmt.Printf("Name: %v, Age: %v\n", name, age)
	fmt.Printf("Type of pi: %T\n", pi)

	// EGÉSZ SZÁM
	fmt.Printf("Decimal: %d, Binary: %b, Hex: %x\n", 42, 42, 42)
	// %d - 10-es számrendszer
	// %b - 2es számrendszer
	// %x - 16os számrendszer

	// FLOAT
	fmt.Printf("Pi raw: %v, Two decimals: %.2f\n", pi, pi)
	// %v - raw
	// %.2f - két tizedes jegyik tartó formázás

	// STRING
	fmt.Printf("Normal: %s, Quoted: %q\n", name, name)
	// %s - string
	// %q - quote

	// BOOL
	fmt.Printf("Is Go fun? %t\n", true)
}


//  FORMÁZOTT STRING BASICS
// Általános

// %v → "default" formátum (bármit kiír, ahogy tudja).
// %+v → structnál mezőneveket is kiír.
// %#v → Go-szintaxis szerinti reprezentáció (debugra jó).
// %T → a változó típusát írja ki.

// Számok

// %d → egész szám (decimal).
// %b → bináris.
// %o → oktális.
// %x → hexadecimális.
// %f → lebegőpontos.
// %.2f → lebegőpontos, 2 tizedesre kerekítve.
// %e → tudományos jelölés (pl. 1.23e+06).

// String / karakter

// %s → string.
// %q → idézőjelek közötti string (pl. "hello").
// %c → egyetlen karakter (ASCII kód alapján).
// Bool
// %t → true vagy false.