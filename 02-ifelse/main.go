package main

func main() {
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
}
