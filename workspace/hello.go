package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Enter the two integers: ")
	var x, y int
	fmt.Scanln(&x)
	fmt.Scanln(&y)

	if y == 0 {
		fmt.Println("Please give non-zero value to y: ")
		fmt.Scanln(&y)
	}
	a := float64(x)
	b := float64(y)
	res := math.Mod(a, b)

	if res == 0 {
		fmt.Println("y is a factor of x")
	} else {
		fmt.Println("y is not a factor of x")
	}
}
