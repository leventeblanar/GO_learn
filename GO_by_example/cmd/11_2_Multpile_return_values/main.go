package main

import (
	"fmt"
)

//  built in function for multiple return values
//	(int, int) in function signature 
//  use _ to only return a subset of the returned values (eg. if you return an error as well but you don't want to)

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}