package main

import (
	"fmt"
)

func main() {
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
}
