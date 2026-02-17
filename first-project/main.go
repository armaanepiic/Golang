package main

import "fmt"

// init function
// you can't call this function
// computer call this automatically
var a = 10
func main() {
	fmt.Println(a)
	// init() xxxx
}

func init() {
	fmt.Println(a)
	a = 20
}