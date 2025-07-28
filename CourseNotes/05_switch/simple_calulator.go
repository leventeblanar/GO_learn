package main

import (
	"fmt"
)

func main() {

	for {
		var operator string
		fmt.Println("Enter operator: +, -, *, /")
		fmt.Scan(&operator)

		var num1, num2 float64
		fmt.Println("Enter two numbers: ")
		fmt.Scan(&num1, &num2)

		switch operator{
		case "+":
			fmt.Println(num1 + num2)
		case "-":
			fmt.Println(num1 - num2)
		case "*":
			fmt.Println(num1 * num2)
		case "/":
			if num2 == 0 {
				fmt.Println("The number you are deviding with cannot be 0.")
				continue
			} else {
				fmt.Println(num1 / num2)
			}
		}
	}
}