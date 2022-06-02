package main //main.go run first

import "fmt"

//compile and run first
func main() {
	//constant variable
	const name string = "Ananya"

	//manual type declaration
	var age int = 22

	//type inferrence
	score := 81.6
	var isPass = true

	fmt.Println("My name is", name) //print with new line
	fmt.Println("Age:", age)
	fmt.Println("Score:", score)
	fmt.Println("Pass exam:", isPass)

	fmt.Printf("\nName: %v\n", name)       //%v = print value
	fmt.Printf("Type of name: %T\n", name) //%T = print type

	// Math Operator
	var num1, num2 = 10, 3

	fmt.Println("add =", num1+num2)
	fmt.Println("minus =", num1-num2)
	fmt.Println("times =", num1*num2)
	fmt.Println("divided =", num1/num2)
	fmt.Println("modulo =", num1%num2)

	// Logic Operator
	fmt.Println(num1, "equal", num2, ":", num1 == num2)
	fmt.Println(num1, "not equal", num2, ":", num1 != num2)
	fmt.Println(num1, "more than", num2, ":", num1 > num2)
	fmt.Println(num1, "less than", num2, ":", num1 < num2)
	fmt.Println(num1, "more than or equal", num2, ":", num1 >= num2)
	fmt.Println(num1, "less than or equal", num2, ":", num1 <= num2)
}
