package main
import "fmt"

func main() {
	// integer
	var x int = 100
	var y = 200
	z := 300
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// float
	var a float32 = 23.34
	fmt.Println(a)

	// string
	message := "I love you"
	fmt.Println(message)

	// boolean
	isOk := true
	fmt.Println(isOk)
	isOk = false
	fmt.Println(isOk)
	// isOk = "Hello" => error because at first declaration isOk was boolean

	// const
	const pi = 3.1416
	fmt.Println(pi)

	// condition
	age := 20
	sex := "male"

	if(age > 60 || sex == "male") {
		fmt.Println("You are ready to marry")
	}

	isPretty := false
	if (!isPretty) {
		fmt.Println("You are pretty")
	}

	day := "monday"

	switch day {
	case "friday" , "saturday": fmt.Println("Holiday")
	case "monday" : fmt.Println("half office day")
	default:
		fmt.Printf("Office day")
	}

	// if(age > 18) {
	// 	fmt.Println("You are ready for marriage")
	// } else if (age == 18) {
	// 	fmt.Println("you are just 18")
	// } else {
	// 	fmt.Println("you are teenager")
	// }
}