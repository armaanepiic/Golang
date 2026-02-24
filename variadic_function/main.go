package main

import "fmt"


func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main() {
	print(1,2,3,4,5,6,7,8) // we can send any number of data
}
