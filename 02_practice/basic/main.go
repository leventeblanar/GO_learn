package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Weather struct {
	City		string	`json:"city"`
	Temperature float64 `json:"temperature"`
	Windspeed	float64 `json:"windspeed"`
}

func main() {
	readings := []Weather {
		{City: "Győr", Temperature: 18.7, Windspeed: 12.3},
		{City: "Budapest", Temperature: 20.1, Windspeed: 9.8},
		{City: "Sopron", Temperature: 17.4, Windspeed: 14.2},
	}

	data, err := json.Marshal(readings)
	if err != nil {
		log.Fatal("JSON Marshal error", err)
	}

	fmt.Println(string(data))

	var decoded []Weather

	err = json.Unmarshal(data, &decoded)
	if err != nil {
		log.Fatal("JSON Unmarshal error", err)
	}

	for index, reading := range decoded {
		index += 1
		fmt.Printf("%d. City: %s, Temp.: %.1f, Windspeed: %.1f\n", index, reading.City, reading.Temperature, reading.Windspeed)
	}
}