package main

import (
	"fmt"
)

func makeAccount(balance int) func(int) int {

	currentBalance := balance
	return func(amount int) int {
		currentBalance = currentBalance + amount
		return currentBalance
	}
}



func main() {
	account := makeAccount(100)

	fmt.Println(account(50))
	fmt.Println(account(-50))
	fmt.Println(account(250))
	fmt.Println(account(-450))
}