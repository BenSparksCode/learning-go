package main

import (
	"fmt"
)

func main() {

	x := 10
	y := 10

	color := "green"

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is not less than %d\n", x, y)
	}

	switch color {
	case "RED":
		fmt.Println("Color is Red")
	case "BLUE":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color is not blue or red")
	}

}
