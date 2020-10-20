package main

import "fmt"

func main() {

	ids := []int{33, 4, 5, 623, 34, 676, 1}

	// loop through ids using range

	for i, id := range ids {
		fmt.Printf("%d. ID: %d\n", i, id)
	}

	// when not using the i counter
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// range with map
	ages := map[string]int8{"Ben": 24, "Josh": 23, "Toby": 3}

	for person, age := range ages {
		fmt.Printf("%s is %d years old\n", person, age)
	}

}
