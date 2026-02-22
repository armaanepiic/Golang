package main

import "fmt"

type User struct { 
	Name string
	Age int
}

// reciever function
func (usr User) printDetails () {
	fmt.Println("Name=", usr.Name)
	fmt.Println("Age=", usr.Age)
}

func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main() {
	var user1 User
	// instantiate
	user1 = User{
		Name: "Arman",
		Age: 30,
	}

	// printUserDetails(user1)
	user1.printDetails()
	

	// instantiate
	user2 := User{ // Instance
		Name: "Nusrat",
		Age: 28,
	}
	// printDetails(user2)
	user2.printDetails()

	user1.call(19)

	
}


/*

*** code segment ***

	User = type User struct {...}
	printUserDetails = func() {...}
	call = func(int) {...}
	main = func() {...}


	*** reciever function cannot work without struct

*/