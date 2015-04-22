package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string = "abc"
	var (
		a int
		c int
	)
	fmt.Printf("%v %v %v %q %v %v", i, f, b, s, a, c)
}
