package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	delta := 0.000001
	z := 1.0
	z_prev := 0.0
	for math.Abs(z-z_prev) > delta {
		z_prev = z
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
