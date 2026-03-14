package main

import "fmt"

func main() {
	a := 5
	j := 10.23343
	fmt.Printf("%.2f\n", j)
	fmt.Println(j)

	r := '❤'
	fmt.Println(r)
	fmt.Printf("%c\n", r)

	fmt.Println(a)
	flag := false
	fmt.Printf("%t\n", flag)

	var s string = "My name is arman"
	fmt.Printf("%s\n", s)
	fmt.Printf("Type if variable s = %T\n", s)
	fmt.Printf("Type if variable flag = %T\n", flag)
}