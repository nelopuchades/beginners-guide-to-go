package main

import (
	"fmt"
)

func main() {
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
