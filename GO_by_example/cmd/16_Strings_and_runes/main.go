package main

import (
	"fmt"
	"unicode/utf8"
)

// A GO-ban a string egy UTF-8 bájt-sorozat
// Indexelés során (s[i]) bytot kapunk vissza (pl. é = 2 byte)
// A rune a Unicode "code point"-ja -> egy karakter, függetlenül attól, hogy hány bíte
// Ha for range-el mész stringen, akkor rune-okat kapsz (ez a helyes mód)
// len(s) -> byte hossz
//  utf8.RuneCountInString(s) -> karakterek száma

func main() {
	s := "héllo"

	fmt.Println("String:", s)
	fmt.Println("Byte length:", len(s))						// 6 (é = 2 byte)
	fmt.Println("Rune count:", utf8.RuneCountInString(s))	// 5

	// Byte-onként
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: %x\n", i, s[i])
	}

	for i, r := range s {
		fmt.Printf("%d: %c\n", i, r)
	}
}
