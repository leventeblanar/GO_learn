package main

import "fmt"

func main() {
	var a int = 10  // expliciten típusos
	var b = 20		// típuslevezetés
	c := 30			// rövid deklarálás (csak függvényen belül él)
	var s string	// "" (zero value)
	var z int		// 0 (zero value)

	fmt.Println(a, b, c, s, z)
}