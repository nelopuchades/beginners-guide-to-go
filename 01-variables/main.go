package main

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
}
