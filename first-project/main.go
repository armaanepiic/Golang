package main

import "fmt"

var a = 20
var b = 30

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	var p = 30
	var q = 40
	add(p, q)
	add(a, b)
	add(a, p)
}