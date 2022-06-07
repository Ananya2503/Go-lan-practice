package mathoperator

import "fmt"

func MathOperator() {
	fmt.Println("***** Math operator *****")
	fmt.Println("--------------------------")
	fmt.Println("| Operator | Description |")
	fmt.Println("|------------------------|")
	fmt.Println("|    +     |     Add     |")
	fmt.Println("|------------------------|")
	fmt.Println("|    -     |   Subtract  |")
	fmt.Println("|------------------------|")
	fmt.Println("|    *     |  Multiples  |")
	fmt.Println("|------------------------|")
	fmt.Println("|    /     |   Divided   |")
	fmt.Println("|------------------------|")
	fmt.Println("|    %     |    Modulo   |")
	fmt.Println("--------------------------")
	fmt.Println()
}

func Add(num1, num2 int) int {
	return num1 + num2
}

func Subtract(num1, num2 int) int {
	return num1 - num2
}

func Multiples(num1, num2 int) int {
	return num1 * num2
}

func Divided(num1, num2 int) (int, string) {
	if num2 == 0 {
		return 0, "divided cannot be 0"
	} else {
		return num1 / num2, ""
	}
}

func Modulo(num1, num2 int) int {
	return num1 % num2
}
