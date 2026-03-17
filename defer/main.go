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

func main() {
	res := sum(3, 4)
	fmt.Println(res)
}