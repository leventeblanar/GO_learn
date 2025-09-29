package main

import "fmt"

// Egyszerű függvény
func add(a int, b int) int {
	return a + b
}

// Több visszatérési érték
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// Függvény változóként
func sayHello(name string) {
	fmt.Println("Hello,", name)
}


func main() {
	// egyszerű hívás
	sum := add(5, 7)
	fmt.Println("sum:", sum)

	// több visszatérési érték
	q, r := divide(17, 5)
	fmt.Println("Quotient:", q, "Remainder:", r)

	// egyszerű üdvözlés
	sayHello("Levi")
}