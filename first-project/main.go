package main

import "fmt"

const a = 10
var p = 100


func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}

	add(5, 6)
	add(p, a)
}

func main() {
	call()
	// a := 12 // shadowing
	fmt.Println(a)
}

func init () {
	fmt.Println("I am init function")
}

/*
go file run process:
	2 phases
	  1. compilation phase -> binary file
	  2. execution phase


	  go run main.go => compile it => main => ./main
	  go build main.go => compile it => main
*/

/*
	*** Binary File ***

*/

/*
	*** compilation phase ***
	*** code segment ***

		const a = 10
		call function () {...}
		add function () {...}
		main function () {...}
		init function () {...}
		
*/