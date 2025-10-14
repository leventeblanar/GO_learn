package main

import (
	"fmt"
	"log"
	"unit_converter/engine"
)

func main() {
	var value float64
	var from, to ,category string

	fmt.Println("Welcome to Go Unit Converter!")

	fmt.Print("Enter category (length / weight / temp): ")
	fmt.Scanln(&category)
	
	switch category {
	case "length": 
		fmt.Println("Length unit options: 'millimeter', 'centimeter', 'meter', 'kilometer', 'inch', 'foot', 'yard', 'mile'")
	case "weight":
		fmt.Println("Weight unit options: 'milligram', 'gram', 'kilogram', 'ounc', 'pound'")
	case "temp":
		fmt.Println("Temperature unit options: 'celsius', 'fahrenheit', 'kevin'")
	}

	fmt.Println("Enter From unit: ")
	fmt.Scanln(&from)
	fmt.Println("Enter To unit: ")
	fmt.Scanln(&to)
	fmt.Println("Enter value: ")
	fmt.Scanln(&value)

	var result float64
	var err error

	switch category {
	case "length": 
		result, err = engine.ConvertLength(value, from, to)
	case "weight":
		result, err = engine.ConvertWeight(value, from, to)
	case "temp":
		result, err = engine.ConvertTemperature(value, from, to)
	default:
		log.Fatalf("Unknown category: %s", category)
	}

	if err != nil {
		log.Fatalf("Error during the calculation: %s", err)
	}
	fmt.Printf("result: %.2f %s\n", result, to)
}
