package main

import "fmt"

// anonymous function = function without name
// IIFE = immediately invoked function expression

func main() {
	func (a, b int) {
		fmt.Println("I am a anonymous function")
		c := a + b
		fmt.Println(c)
	}(5, 9)
}