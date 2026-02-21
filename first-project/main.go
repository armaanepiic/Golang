package main

import "fmt"
const a = 10
var p = 100 // data segment

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age =", age)

	show := func() {
		money = money + a + p 
		fmt.Println(money)
	}
	return show // in this line money and show func will be stored in heap
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("I am init function")
}

/*


2 phases =>
	1. compilation phase (compile time)
	2. execution phase (run time)

	1** compile phase **
	** code segment **
	a = 10
	outer = func() {...}
	outer anonymous1 = func() {...}
	call = func() {...}
	main = func() {...}
	init = func() {...}

	2** execution phase **
*/