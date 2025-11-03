package main

import (
	"fmt"
)

// Counter struct(count)
type Counter struct {
	count		int
}

// Increment (+1)
func (c *Counter) Increment() {
	c.count++
}

// Decrement (-1)
func (c *Counter) Decrement() {
	c.count--
}

// Reset 
func (c *Counter) Reset() {
	c.count = 0
}

// GetValue() int
func (c Counter) GetValue() int {
	return c.count
}

// IncrementBy (amount int)
func (c *Counter) IncrementBy(num int) {
	c.count = c.count + num
}

func main() {
	counter1 := Counter{
		count: 0,
	}
	
	counter1.Increment()
	counter1.Increment()
	counter1.Increment()
	counter1.IncrementBy(5)

	fmt.Printf("Új érték: %d\n", counter1.GetValue())

	counter1.Reset()

	fmt.Printf("Új érték: %d\n", counter1.GetValue())
}