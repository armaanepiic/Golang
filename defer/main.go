package main

import "fmt"

func a() {
	i := 0

	fmt.Println("first", i)

	defer fmt.Println("second", i)

	i = i + 1
	
	fmt.Println("third", i)
	
	defer fmt.Println("fourth", i)
	

}

func sum(a int, b int) (result int) {
	result = a + b
	return
}

func calculate() (result int) {

	fmt.Println("First", result)

	show := func() {
		result = result + 10
		fmt.Println("defer",result)
	}
	defer show()

	result = 5

	p := func(a int) {
		fmt.Println("ami", a)
	}

	defer p(result)

	defer fmt.Println(result)
	
	fmt.Println("Second", result)

	defer fmt.Println(5)

	return
}


func main() {
	a := calculate()
	fmt.Println("main first", a)
}

/*

=> named return type
	1. all codes executes
	2. defer function stored in the magic box
	3. return -> all defer func executed
	4. return values of named variable
=> normal return type
	1. all codes executes
	2. defer function stored in the magic box
	3. return values are evaluated at this time
	4. all defer functions are executed
*/