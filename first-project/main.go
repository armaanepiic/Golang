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
}