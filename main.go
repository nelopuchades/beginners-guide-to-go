package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Variables
	var welcome string = "hello world"
	// welcome := "hello world"
	println(welcome)

	var myAge int = 25
	// myAge := 25

	// Arithmetic operators
	myAge += 10 // 35
	myAge -= 10 // 25
	myAge *= 2  // 50
	myAge--     // 49
	myAge++     // 50

	println(myAge) // 50

	// Const variable
	const salary int = 28000
	println(salary) // 28000
	// error -> salary := 25000

	const pi = 3.14
	println(pi) // go compiler infers pi type - float32

	var customerHeight int = 135
	customerAge := 18

	// if conditional
	if customerHeight >= 150 || customerAge >= 18 {
		println("Can access ride! :D")
	} else if customerHeight >= 120 {
		println("Can access ride but with an adult!")
	} else {
		println("Cannot access ride... :(")
	}

	// Can access ride! :D

	// switch conditional
	switch {
	case customerHeight >= 150 || customerAge >= 18:
		println("Can access ride! :D")
	case customerHeight >= 120:
		println("Can access ride but with an adult!")
	default:
		println("Cannot access ride... :(")
	}

	os := runtime.GOOS

	switch os {
	case "darwin":
		println("OS X")
	case "linux":
		println("Linux machine")
	default:
		println("something else")
	}

	// arrays
	planets := [6]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturns"}
	fmt.Println(planets)

	var planetsArray [8]string
	planetsArray[0] = "mercury"
	planetsArray[4] = "venus"
	fmt.Println(planetsArray)

	planetSlice := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturns"}
	fmt.Println(planetSlice)

	var planetSliceVerbose []string
	planetSliceVerbose = append(planetSliceVerbose, "mercury")
	planetSliceVerbose = append(planetSliceVerbose, "mars")
	fmt.Println(planetSliceVerbose)

	// For loop
	ages := []int{42, 28, 30, 27, 18}
	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}

	for i := 0; i < len(ages); {
		fmt.Println(ages[i])
		i++
	}

	// There's no do-while loops in GO

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}

		fmt.Println(i)
	}

	for index, value := range ages {
		fmt.Println(index)
		fmt.Println(value)
	}
}
