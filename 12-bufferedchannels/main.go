// Channels are effectively pipes that link between your defined
// goroutines and allow you to send and receive from other
// goroutines to each other

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValues(values chan int) {
	for i := 0; i <= 10; i++ {
		value := rand.Intn(10)
		fmt.Printf("Value Calculated in F: %d", value)
		values <- value
	}
}

func main() {
	values := make(chan int, 2)
	go CalculateValues(values) //Spans a new thread that executes this function
	for i := 0; i <= 10; i++ {
		time.Sleep(2 * time.Second)
		value := <-values // Blocks until values channel returns something
		fmt.Printf("Value Calculated outside F: %d", value)
	}
}
