package main

import "fmt"

func main() {

	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T %T\n", a, b)
	fmt.Println(*b)

	// change val with pointer
	fmt.Println("---------------")
	*b = 10
	fmt.Println(a)

}
