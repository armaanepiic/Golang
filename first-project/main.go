package main

import "fmt" 

// arameter vs argument
// sends as an arguments
// recieves as a parameter

// as it recieves a function => higher order function
func processOperation(a int, b int, op func (p int, q int)) {
	op(a, b)
}

func add(x int, y int) { // parameter => a, b
	z := x + y
	fmt.Println(z)
}

func main() {
	processOperation(2, 5, add)
}

/*
1. parameter vs argument
2. first order function
	i. standard function or named function
	ii. anonymous function
	iii. IIFE
	iv. function expression
3. higher order function
*/