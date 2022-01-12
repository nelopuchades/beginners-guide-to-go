package main

import (
	"runtime"
)

func main() {
	var customerHeight int = 135
	customerAge := 18

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
