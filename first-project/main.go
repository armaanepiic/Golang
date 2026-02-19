package main

import "fmt"

// anonymous function = function without name
// IIFE = immediately invoked function expression

// function expression or assign function in variable

func main() {
	add := func (a, b int) {
		fmt.Println("I am a anonymous function")
		c := a + b
		fmt.Println(c)
	}
	add(3, 8)
}

func init() {
	fmt.Println("i am init function")
}