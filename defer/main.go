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
	fmt.Println("Second", result)

	return
}

func calc() int {

	result := 0
	fmt.Println("First", result)

	show := func() {
		result = result + 10
		fmt.Println("defer",result)
	}
	defer show()

	result = 5
	fmt.Println("Second", result)

	return result
}

func main() {
	a := calculate()
	fmt.Println("main first", a)

	b := calc()
	fmt.Println("main second", b)
}