package logicoperator

import "fmt"

func LogicOperator() {
	fmt.Println("***** Logic operator *****")
	fmt.Println("---------------------------------")
	fmt.Println("| Operator |     Description    |")
	fmt.Println("|-------------------------------|")
	fmt.Println("|    ==    |        Equal       |")
	fmt.Println("|-------------------------------|")
	fmt.Println("|    !=    |      Not equal     |")
	fmt.Println("|-------------------------------|")
	fmt.Println("|    >     |      More than     |")
	fmt.Println("|-------------------------------|")
	fmt.Println("|    <     |      Less than     |")
	fmt.Println("|-------------------------------|")
	fmt.Println("|    >=    | More than or equal |")
	fmt.Println("|-------------------------------|")
	fmt.Println("|    <=    | Less than or equal |")
	fmt.Println("---------------------------------")
	fmt.Println()
}

func Equal(a, b int) bool {
	return a == b
}

func NotEqual(a, b int) bool {
	return a != b
}

func MoreThan(a, b int) bool {
	return a > b
}

func LessThan(a, b int) bool {
	return a < b
}

func MoreThanOrEqual(a, b int) bool {
	return a >= b
}

func LessThanOrEqual(a, b int) bool {
	return a <= b
}
