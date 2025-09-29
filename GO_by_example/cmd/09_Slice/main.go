package main

import (
	"fmt"
)

// Típusa: []T - nincs méret a típusban
// Három mezőből álló fejlécként gondolj rá: pointer a backing array-re + len + cap
// Referenciaként viselkedik: ha függvénynek adod, a tartalom módosulhat
// Nőhet append-del, a Go kezeli a kapacitást (néha új tömböt is allokálhat)
// Nem összehasonlítható (kivéve nil-lel)

func main() {
	s := []int{1, 2, 3}			// len=3, cap=3
	s = append(s, 4)			// ha kell kapacitást növel
	fmt.Println(s)				// [1 2 3 4]

	// slicing
	a := []int{10, 20, 30, 40, 50}
	part := a[1:4]				// [20 30 40], len=3, cap (gyakran 4)
	part[0] = 999

	fmt.Println(a)				// [10 999 30 40 50] <- közös backing array

	// függvény módosítja a slice tartalmát
	makeAllPositive(a)
	fmt.Println(a)
}

func makeAllPositive(nums []int) {
	for i := range nums {
		if nums[i] < 0 {
			nums[i] = -nums[i]
		}
	}
}



//  MAKE() 

//  A make() új sliceot hoz létre a memóriában és megadhatod neki
//  	- a hosszát (mennyi elem legyen elérhető azonnal)
// 		- a kapacitását (mennyi hely van lefoglalva a háttérben)

func mapExample() {
	// csak length
	s1 := make([]int, 5)
	fmt.Println(s1)
	fmt.Println("len:", len(s1), "cap:", cap(s1))
	// -> 5 elem, 5 kapacitás

	// length és cap külön
	s2 := make([]int, 3 ,5)
	fmt.Println(s2)
	fmt.Println("len:", len(s2), "cap:", cap(s2))
	// -> 3 elem, de háttérben 5 hely

	// append
	s2 = append(s2, 10, 20)
	fmt.Println(s2)
	fmt.Println("len:", len(s2), "cap:", cap(s2))
	// -> 5 elem, 5 kapactiás -> pont kitelt

	// még egy append
	s2 = append(s2, 99)
	fmt.Println(s2)
	fmt.Println("len:", len(s2), "cap:", cap(s2))
	// -> 6 elem, 10 kapacitás -> Go új tömböt allkotált dupla mérettel
}