package main

import "fmt"

// Engineer - stores the name and age of an engineer
type Engineer struct {
	Name string
	Age  int
}

func (e *Engineer) UpdateName() {
	e.Name = "Primeagen"
	fmt.Println(e) // {Primeagen 26}
}

func (e *Engineer) UpdateAge() {
	e.Age += 1
}

func UpdateAge(e *Engineer) {
	e.Age += 1
}

func main() {
	fmt.Println("Go Pointers tutorial")

	nelo := &Engineer{
		Name: "Nelo",
		Age:  25,
	}

	fmt.Println(nelo) // &{Nelo 25}

	nelo.UpdateAge()
	fmt.Println(nelo) // &{Nelo 26}

	nelo.UpdateName()
	fmt.Println(nelo) // &{Primeagen 26}

	UpdateAge(nelo)
	fmt.Println(nelo) // &{Primeagen 27}

	// var name string
	// name = "nelo"
	// fmt.Println(name) // nelo

	// ptr := &name
	// fmt.Println(ptr)  // 0xc00011e210
	// fmt.Println(*ptr) // nelo

	// *ptr = "lucifer"
	// fmt.Println(ptr)  // 0xc00011e210
	// fmt.Println(*ptr) // lucifer
	// fmt.Println(name) // lucifer
}
