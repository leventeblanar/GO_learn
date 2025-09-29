package pointers

import (
	"fmt"
)

// Pointer - egy változó memória címére mutatunk, nem az értékére
// 			- &x -> adja a változó címét (pointer)
//			- *p -> adja vissza a pointer által mutatott értéket (dereferencing)
// Függvényeknél azért jó, mert ha pointert adunk át, a függvény ténylegesen tudja módosítani a változót, nem másolattal dolgozik


func increment(n *int) {
	*n = *n + 1
}



func main() {
	x := 10
	p := &x			// p egy pointer, ami x-re mutat

	fmt.Println("x:", x)
	fmt.Println("p:", p)		// memóriacím (pl. 0xc0000140a8)
	fmt.Println("*p:", *p)		// az érték: 10

	// pointeren keresztül módosítjuk
	*p = 20
	fmt.Println("x után:", x)	// 20



	val := 5
	increment(&val)
	fmt.Println(val)
}