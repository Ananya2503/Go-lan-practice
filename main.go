package main //main.go run first

import "fmt"

//compile and run first
func main() {

	// Variable declaration
	fmt.Println("-----Variable declaration-----")
	variableDeclaration()

	// Math Operator
	fmt.Printf("\n-----Math Operator-----\n")
	mathOperator(10, 3)

	// Logic Operator
	fmt.Printf("\n-----Logic Operator-----\n")
	logicOperator(10, 3)

	// Scanf
	fmt.Printf("\n-----Get input from keyboard-----\n")
	getInputFromKeyboard()
}

func variableDeclaration() {
	//constant variable
	const name string = "Ananya"

	//manual type declaration
	var age, height = 22, 164

	//type inferrence
	score := 81.6
	var isPass = true

	fmt.Println("My name is", name) //print with new line
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Score:", score)
	fmt.Println("Pass exam:", isPass)

	fmt.Printf("\nName: %v\n", name)       //%v = print value
	fmt.Printf("Type of name: %T\n", name) //%T = print type
}

func mathOperator(num1 int, num2 int) {
	fmt.Println("add =", num1+num2)
	fmt.Println("minus =", num1-num2)
	fmt.Println("times =", num1*num2)
	fmt.Println("divided =", num1/num2)
	fmt.Println("modulo =", num1%num2)
}

func logicOperator(num1 int, num2 int) {
	fmt.Println(num1, "equal", num2, ":", num1 == num2)
	fmt.Println(num1, "not equal", num2, ":", num1 != num2)
	fmt.Println(num1, "more than", num2, ":", num1 > num2)
	fmt.Println(num1, "less than", num2, ":", num1 < num2)
	fmt.Println(num1, "more than or equal", num2, ":", num1 >= num2)
	fmt.Println(num1, "less than or equal", num2, ":", num1 <= num2)
}

func getInputFromKeyboard() {
	// fmt.Scanf(string_format, address_list)
	var name string
	fmt.Print("Enter name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello,", name)
}
