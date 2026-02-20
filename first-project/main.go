package main

import "fmt"

// arameter vs argument
// sends as an arguments
// recieves as a parameter

func add(a int, b int) { // parameter => a, b
	c := a + b
	fmt.Println(c)
}
func main() {
	add(2, 5) // argumants => 2, 5
}

/*
1. parameter vs argument
2. first order function
	i. standard function or named function
	ii. anonymous function
	iii. IIFE
	iv. function expression
3.
*/