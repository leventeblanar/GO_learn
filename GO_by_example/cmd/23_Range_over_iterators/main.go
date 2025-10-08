package main

import (
	"fmt"
	"time"
)

//	Go-ban a range kulcsoszóval általában kollekciókat járunk be

//	for i, v := range []int{10, 20, 30} {
//	fmt.Println(i, v)}

//	ez működik szeletekkel, mapekkel, stringekkel, csatornákkal - de eddig nem működött olyan saját típusokkal, amiket te hozol létre (pl. egy saját iterátor)

// 	Miért kell range over iterators?
// 	korábban csak két lehetőség (Channel -a go routine küldi, a range olvassa || Next() metódus - explicit, kézzel kell ciklizni)

//			for v := range MyGenerator() {
// 				fmt.Println(v)
// 			}
// ez pl egy range function egy függvényen iterál és addig hívogatja a "yield" hívást amíg az vissza nem tér false-al


//////////////////////////////////////////////
/// 	1. Channel alapú iterátor
//////////////////////////////////////////////


// 	CountwithChannel visszaad egy <- chan int-et
//	a goroutine tölti fel a számokat, a for-range olvassa
func CountWithChannel(n int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)  				// amikor vége bezárjuka. csatornát
		for i := 1; i <= n; i++ {
			ch <- 1						// küldjük a számokat
			time.Sleep(100 * time.Millisecond)  //csak hogy látni lehessen a sorrendet
		}
	}()
	return ch							// visszaadjuk a csatornát
}





//////////////////////////////////////////////
/// 	2. NEXT() Metódusos iterátor
//////////////////////////////////////////////

// Struct tárolja az állapotot
type Counter struct {
	current			int
	limit			int
}

// konstruktor
func NewCounter(limit int) *Counter {
	return &Counter{current:0, limit: limit}
}

// Next() egyenként adja a számokat
// Ha elérte a határt, false-szal jelzi a végét
func (c *Counter) Next() (int, bool) {
	if c.current >= c.limit {
		return 0, false
	}
	c.current++
	return c.current, true
}





//////////////////////////////////////////////
/// 	3. Range-over-func (Go 1.22+)
//////////////////////////////////////////////

//CountWithRangeFunc visszaad egy olyan függvényt
// amit a range be tud járni
func CountWithRangeFunc(n int) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) {				// ha a for-range leállítja
				return
			}
		}
	}
}




func main() {

	fmt.Println("\n--- Chanel alapú iterátor ---")
	for v := range CountWithChannel(5) {
		fmt.Println(v)
	}


	fmt.Println("\n--- Next() metódusos iterátor ---")
	c := NewCounter(5)
	for val, ok := c.Next(); ok; val, ok = c.Next() {
		fmt.Println(val)
	}

	
	fmt.Println("\n--- Range-over-func iterator ---")
	for v:= range CountWithRangeFunc(5) {
		fmt.Println(v)
	}
}