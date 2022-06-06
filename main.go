package main //main.go run first

import "fmt"

//compile and run first
func main() {
	fmt.Println("1. Variable declaration")
	fmt.Println("2. Math Operator")
	fmt.Println("3. Logic Operator")
	fmt.Println("4. If statement")
	fmt.Println("5. Array")
	fmt.Println("6. Slice")

	// switch case
	var choice int
	fmt.Print("\nEnter number: ")
	fmt.Scanf("%d", &choice)
	fmt.Println()

	switch choice {
	case 1:
		variableDeclaration()
	case 2:
		var num1, num2 int
		fmt.Scanf("%d %d", &num1, &num2)
		mathOperator(num1, num2)
	case 3:
		var num1, num2 int
		fmt.Scanf("%d %d", &num1, &num2)
		logicOperator(num1, num2)
	case 4:
		var score int
		fmt.Scanf("%d", &score)
		ifStatement(score)
	case 5:
		array()
	case 6:
		slice()
	default:
		fmt.Println("Incorrect number")
	}
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

func ifStatement(score int) {
	/*
		if condition {
			statement
		} else if {
			statement
		} else {
			statement
		}
	*/

	if score >= 50 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func array() {
	// Array declaration
	var numbers1 [3]int = [3]int{100, 200, 300} // fixed size
	numbers2 := [...]int{1, 2, 3, 4, 5}         // infinite array
	numbers3 := [3]int{10}

	fmt.Println(numbers1)
	fmt.Println(numbers2)
	fmt.Println(numbers3)
	fmt.Println("Size of array 1:", len(numbers1))
	fmt.Println("Size of array 2:", len(numbers2))
}

func slice() {
	var numbers1 []int = []int{0, 1, 2, 3, 4}
	numbers2 := []int{10, 20}

	fmt.Println(numbers1)
	fmt.Println(numbers2)

	// append
	numbers1 = append(numbers1, 5)
	fmt.Println("After append 5 in numbers1: ", numbers1)

	// เข้าถึงแบบเป็นช่วง -> slice_name[start_index:end_index + 1]
	fmt.Println("Element from index 1 to 5: ", numbers1[1:])
	fmt.Println("Element from index 0 to 2: ", numbers1[:3])
	fmt.Println("Element from index 2 to 3: ", numbers1[2:4])
}
