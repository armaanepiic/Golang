package main

import "fmt"

// init function
// you can't call this function
// computer call this automatically

func main() {
	fmt.Println("Hello init function")
	// init() xxxx
}

func init() {
	fmt.Println("I am the first function that is executed first")
}