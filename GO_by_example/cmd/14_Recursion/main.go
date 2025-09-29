package main

import (
	"fmt"
)

// A rekurzió azt jelenti, hogy egy függvény önmagát hívja meg
// Kell egy alapeset (base case), ami megállítja a végtelen hívást
// A többi esetben a függvény önmagát hívja kisebb részekre bontva a feladatot


// Gyakori:
// 		- matematikai sorozatoknál (faktoriális, Fibonacci)
//		- fájlstruktúrák bejárásánál
//		- fa/JSON feldolgozásnál


func factorial(n int) int {
	if n == 0 {
		return 1	// base case
	}
	return n * factorial(n-1)	// recursive case
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(factorial(5)) // 5 * 4 * 3 * 2 * 1 == 120

	for i := 0; i < 10 ; i++ {
		fmt.Println(fibonacci(i))
	} 
}