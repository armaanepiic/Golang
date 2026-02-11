package main

import (
	"fmt"
	"example.com/mathlib"
)

var a = 20
var b = 30


func main() { 
	fmt.Println("Showing custom package")
	mathlib.Add(10, 20)
	// mathlib.Money
}

// go mod init example.com