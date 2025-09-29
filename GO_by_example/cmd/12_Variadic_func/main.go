package main

import "fmt"

// Egy olyan függvény, ami tetszőleges számú paramétert képes fogadniegy bizonyos típusból
// ilyenkor a nums egy slice-ként érkezik a függvénybe
// lehet, hogy 0, 1 vagy több értéket adunk át, és mind ugyan úgy kezelhető
// olyan esetekre jó, ha nem tudjuk előre, hány számot akarunk összegezni, vagy hány stringet akarunk összefűzni


// Variadic függvény számok összegzésre
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}



func main() {
	// 0 paraméter -> slice üres
	fmt.Println("Sum:", sum())


	// több paraméter
	fmt.Println("Sum:", sum(1, 2))
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))


	// slice kibontása
	numbers := []int{10, 20, 30}
	fmt.Println("Sum:", sum(numbers...))	// ... kell a kibontáshoz
}
