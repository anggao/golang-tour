package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n2 := 0 // f_n-2
	n1 := 1 // f_n-1
	return func() int {
		result := n1
		n1 = n1 + n2
		n2 = result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
