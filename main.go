package main

import "runtime"

func main() {
	var welcome string = "hello world"
	// welcome := "hello world"
	println(welcome)

	var myAge int = 25
	// myAge := 25

	myAge += 10 // 35
	myAge -= 10 // 25
	myAge *= 2  // 50
	myAge--     // 49
	myAge++     // 50

	println(myAge) // 50

	const salary int = 28000
	println(salary) // 28000
	// error -> salary := 25000

	const pi = 3.14
	println(pi) // go compiler infers pi type - float32

	var customerHeight int = 135
	customerAge := 18

	if customerHeight >= 150 || customerAge >= 18 {
		println("Can access ride! :D")
	} else if customerHeight >= 120 {
		println("Can access ride but with an adult!")
	} else {
		println("Cannot access ride... :(")
	}

	// Can access ride! :D

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
}
