package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	a := 0
	b := 1
	cnt := 0
	return func() int {
		if cnt == 0 {
			cnt++
			return a
		} else if cnt == 1 {
			cnt++
			return b
		} else {
			c := a + b
			a = b
			b = c
			cnt++
			return c
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 6; i++ {
		fmt.Println(f())
	}
}
