package main

import (
	"fmt"
	"math"

	"github.com/bensparks/learning-go/03_packages/utils"
)

func main() {

	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(utils.Reverse("Hello World"))

}
