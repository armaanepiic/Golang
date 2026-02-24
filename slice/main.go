package main

import "fmt"

func changeSlice(p []int) []int {
	p[0] = 10
	p = append(p, 11)
	return p
}

func main() {

	// var x []int 		// [], len=0, cap=0
	// x = append(x, 1)	// [1], len=1, cap=1
	// x = append(x, 2)
	// x = append(x, 3)
	// // fmt.Println(x)  

	// y := x
	// x = append(x, 4)
	// y = append(y, 5)
	// x[0] = 10
	// fmt.Println(x) // [10 2 3 5]
	// fmt.Println(y) // [10 2 3 5]


	x := []int{1, 2, 3, 4, 5}

	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changeSlice(a)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x[0:8])












	// arr := [6]string{"This", "is", "a", "go", "interview", "course"}
	// fmt.Println(arr)

	// // slice from an existing array
	// s := arr[1:4]
	// fmt.Println(s)

	// // slice from a slice
	// s1 := s[1:2]
	// fmt.Println(s1)
	// fmt.Println(len(s1))
	// fmt.Println(cap(s1))

	// slice from slice literal
	// s:= []int{1, 2, 5} // slice literal
	// fmt.Println("slice:",s, "len:", len(s), "cap:", cap(s))

	// make function with length
	// s := make([]int, 3) // [0, 0, 0], len = 3, cap = 3
	// s[0] = 5
	// fmt.Println("slice:",s, "len:", len(s), "cap:", cap(s))

	// s := make([]int, 3, 5) // [0, 0, 0], len = 3, cap = 3
	// s[0] = 5				// [5, 0, 0], len = 3, cap = 3
	// // s[3] = 3	// index out of range 
	// fmt.Println("slice:",s, "len:", len(s), "cap:", cap(s))

	// empty slice or nil slice
	// var s []int // len=0, cap=0
	// s = append(s, 1, 2, 3)
	// fmt.Println(s)
}

/*
	
	slice type => 6
		1. slice from an existing array 		=> s := arr[1:4]
		2. slice from a slice					=> s1 := s[1:2]
		3. slice literal						=> s:= []int{1, 2, 5}
		4. make func with len					=> s := make([]int, 3) 
		5. make func with len and capacity		=> s := make([]int, 3, 5) 
		6. emply slice / nil slice				=> var s []int
	
	rule of expanding space: slice underlying array rule => till 1024 (100% increase)
	after 1024 it will increase by 25%
*/



/*

2 phases =>
	1. compilation phase (compile time)
	2. execution phase (run time)

	1** compile phase **
	** code segment **
	main = func() {...}


	2** execution phase **
*/