package main

import "fmt"

func main() {

	var name string = "Ben"
	var age int32 = 37
	var isCool bool = true
	isCool = false //cant do this if isCool was const

	// infer data type
	surname := "Sparks"

	fmt.Println(name, surname, age, isCool)
	// prints the var type
	fmt.Printf("%T\n", age)
}
