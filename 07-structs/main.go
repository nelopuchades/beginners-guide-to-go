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

func main() {
	fmt.Println("Structs")

	engineer := Engineer{
		Name: "Elliot",
		Age:  28,
		Project: Project{
			Name:         "Beginner's Guide to Go",
			Priority:     "Hight",
			Technologies: []string{"Go"},
		},
	}

	fmt.Println(engineer)              // {Elliot, 28}
	fmt.Printf("%+v\n", engineer)      // {Name:Elliot Age:28}
	fmt.Println(engineer.Name)         // Elliot
	fmt.Println(engineer.Project.Name) // Beginner's Guide to Go
}
