package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]int{
		"red":   14,
		"green": 25,
	}
	colors["white"] = 55
	colors["Black"] = 66
	printMap(colors)
}

func printMap(m map[string]int) {

	for k, v := range m {
		fmt.Println(k, v)
	}
}
