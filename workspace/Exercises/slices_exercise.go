package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	b := make([]([]uint8), dy)
	for i := range b {
		b[i] = make([]uint8, dx)

		for j := range b[i] {
			b[i][j] = uint8(i + j)
		}
	}
	return b
}

func main() {
	pic.Show(Pic)
}
