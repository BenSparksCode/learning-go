package main

import (
	"fmt"
)

func main() {

	emails := make(map[string]string)

	emails["Ben"] = "ben@gmail.com"
	emails["Josh"] = "josh@gmail.com"

	// declare and assign map
	ages := map[string]int8{"Ben": 24, "Josh": 23, "Toby": 3}

	fmt.Println(emails)
	fmt.Println(emails["Ben"])
	fmt.Println(len(emails))

	// delete josh
	delete(emails, "Josh")

	fmt.Println(emails)

	fmt.Println(ages)

}
