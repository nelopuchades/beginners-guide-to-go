// Channels are effectively pipes that link between your defined
// goroutines and allow you to send and receive from other
// goroutines to each other

package main

import (
	"fmt"
	"math/rand"
)

func CalculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Printf("Value Calculated: %d", value)
	values <- value
}

func main() {
	values := make(chan int)
	go CalculateValue(values) //Spans a new thread that executes this function
	value := <-values         // Blocks until values channel returns something
	fmt.Println(value)
}
