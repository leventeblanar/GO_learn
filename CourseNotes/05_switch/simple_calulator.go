package main

import (
	"fmt"
)

func main() {

	for {
		var operator string
		fmt.Println("Enter operator: +, -, *, /. You can enter 'exit' to quit.")
		fmt.Scan(&operator)

		if operator != "+" && operator !="-" && operator !="*" && operator != "/" {
			fmt.Println("Unkown operator. Pleaser use: +, -, *, /.")
			continue
		}

		if operator == "exit" {
			fmt.Println("Goodbye!")
			break
		}

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