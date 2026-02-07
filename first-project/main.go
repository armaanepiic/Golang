package main

import "fmt"

// print a value
func add (num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}

// return value
func findProduct (num1 int, num2 int) int {
	result := num1 * num2
	return result
}

func getNumbers (num1 int, num2 int) (int, int) {
	sum := num1 + num2
	product := num1 * num2
	return sum , product
}

func printSomething () {
	fmt.Println("Education must be free")
}

func sayHello (name string) {
	fmt.Println("Welcome to the Golang course, ", name)
}

func main() {

	// function
	var a int = 10
	b := 20
	add (a, b)
	product := findProduct(a, b)
	fmt.Println(product)

	p, q := getNumbers(a, b)
	fmt.Println(p)
	fmt.Println(q)
	printSomething()
	sayHello("Arman")
}