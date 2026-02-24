package main

import "fmt"

// pointer

type User struct{
	Name string
	Age int
	Salary float32
}

func print(numbers *[3]int) {
	fmt.Println(numbers)
}

func printObj(user *User) {
	fmt.Println(user)
}

func main() {
	// x := 20
	// x = 30
	// p := &x // address of x
	// *p = 40
	// fmt.Println(x)
	// fmt.Println("Address:", p)
	// fmt.Println("Value at the address:", *p)\

	// arr := [3]int{1, 2, 3}
	// print(&arr)

	obj := User{
		Name: "Arman",
		Age: 30,
		Salary: 300.34,
	}
	printObj(&obj)
}

/*

2 phases =>
	1. compilation phase (compile time)
	2. execution phase (run time)

	1** compile phase **
	** code segment **
	print = func() {...}
	main = func() {...}

	2** execution phase **
*/