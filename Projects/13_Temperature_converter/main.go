package main

import (
	"fmt"
)

// Temperature struct (value, unit)
type Temperature struct {
	value 	float64
	unit	string
}

// ToFahrenheit()
func (t Temperature) ToFahrenheit() Temperature {
	if t.unit == "C" {
		return Temperature {
			value: t.value*9/5 + 32,
			unit: "F",
		}
	}
	return t
}

// ToCelsius()
func (t Temperature) ToCelsius() Temperature {
	if t.unit == "F" {
		return Temperature {
			value: (t.value-32) * 5/9,
			unit: "C",
		}
	}
	return t
}

// Display - a struct mezőit használja!
func (t Temperature) Display() {
	unitName := "Celsius"
	if t.unit == "F" {
		unitName = "Fahrenheit"
	}
	fmt.Printf("%.2f°%s (%s)\n", t.value, t.unit, unitName)
}



func main() {
	temp1 := Temperature{
		value: 31.5,
		unit:  "C",
	}

	temp1.Display()

	temp2 := temp1.ToFahrenheit()
	temp2.Display()
}
