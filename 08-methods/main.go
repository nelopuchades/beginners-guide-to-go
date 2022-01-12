package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func (e Engineer) Print() { // Just a copy of the object
	fmt.Println("Engineer Information")
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Current Project: %s\n", e.Project.Name)
}

func (e *Engineer) UpdateAge(newAge int) { // Pointer to memory
	e.Age = newAge
}

func (e *Engineer) GetProjectPriority() string {
	return e.Project.Priority
}

func main() {
	fmt.Println("Structs")

	engineer := Engineer{
		Name: "Elliot",
		Age:  28,
		Project: Project{
			Name:         "Beginner's Guide to Go",
			Priority:     "High",
			Technologies: []string{"Go"},
		},
	}

	engineer.Print() // 28
	engineer.UpdateAge(23)
	engineer.Print() // 23

	fmt.Println(engineer.GetProjectPriority())
}
