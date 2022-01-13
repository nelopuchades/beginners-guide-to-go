package main

import (
	"fmt"
	"time"
)

func HelloWorld(name string) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Hello: %s", name)
}

func main() {
	go HelloWorld("Nelo") //Spans a new thread that executes this function
	fmt.Println("I should be printed first")
	time.Sleep(2 * time.Second)
}
