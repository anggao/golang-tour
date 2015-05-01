package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]unit8 {
	slices := make([][]uint8, dy)
	for i := range slices {
		slices[i] = make([]uint8, dx)
	}
	return slices
}

func main() {
	pic.Show(Pic)
}
