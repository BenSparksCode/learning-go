package main

import (
	"fmt"
)

func main() {

	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// declare and assign
	newArr := [2]string{"Bob", "Jamie"}
	newArr2 := []string{"Bob", "Jamie", "Samantha", "Toby"}

	fmt.Println(fruitArr[0])
	fmt.Println(newArr)
	fmt.Println(newArr2)
	fmt.Println(len(newArr2))
	fmt.Println(newArr2[2:])

}
