package main

import "fmt" 

// arameter vs argument
// sends as an arguments
// recieves as a parameter

// as it recieves a function => higher order function
// as it returns a function => higher order function

func processOperation (a int, b int, op func(p int, q int)) func (x int, y int) {
	op(a, b)
	return add
}

func call () func (x int, y int) {
	return add
}
func add(x int, y int) { // parameter => a, b
	z := x + y
	fmt.Println(z)
}

func main() {
	sum := call() // func expression
	sum(10, 20)
	sum2 := processOperation(1, 3, add)
	sum2(4, 8)
}

/*
1. parameter vs argument
2. first order function
	i. standard function or named function
	ii. anonymous function
	iii. IIFE
	iv. function expression
3. higher order function
	// as it recieves a function => higher order function
	// as it returns a function => higher order function
4. callback function => when a func is passed to a higher order func
*/