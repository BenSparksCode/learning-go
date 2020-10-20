package main

import (
	"fmt"
	"strconv"
)

// Person Define person struct
type Person struct {
	firstName string
	lastName  string
	age       int
}

// Greeting method - value reciever
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + ". I am " + strconv.Itoa(p.age)
}

// hasBirthday method - pointer reciever
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// init person struct
	person1 := Person{firstName: "Ben", lastName: "Sparks", age: 24}
	person2 := Person{"Josh", "Wally", 9001}

	fmt.Println(person1)
	fmt.Println(person2)

	person1.age++

	fmt.Println(person1)

	person1.hasBirthday()

	fmt.Println(person1.greet())

}
