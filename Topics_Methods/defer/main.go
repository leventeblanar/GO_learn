package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// defer - there is a deferred processing mechanism that executes a function before the current function returns - always runs in LIFO order
// pretty common in golang code - closing resources, closing resources to databases

// 1. Read a file
func readfile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err  := io.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}


// 2. Measure Performance
func processData(data []int) {
	start := time.Now()

	defer func() {
		fmt.Println(
			"Data processing completed in",
			time.Since(start),
		)
	}()

	for _, d := range data {
		fmt.Println(d)
		time.Sleep(time.Millisecond * 100)
	}
}


// 3. Recovering

func safeOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic.")
		}
	}()

	panic("Something went wrong")
}


func main() {	
	err := readfile("output.txt")
	if err != nil {
		fmt.Println(err)
	}

	data := []int{1, 2, 3, 4, 5}
	processData(data)

	safeOperation()
}