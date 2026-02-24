package main

import "fmt"

func main() {
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
	var s []int // len=0, cap=0
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

/*
	
	slice type => 6
		1. slice from an existing array 		=> s := arr[1:4]
		2. slice from a slice					=> s1 := s[1:2]
		3. slice literal						=> s:= []int{1, 2, 5}
		4. make func with len					=> s := make([]int, 3) 
		5. make func with len and capacity		=> s := make([]int, 3, 5) 
		6. emply slice / nil slice				=> var s []int
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