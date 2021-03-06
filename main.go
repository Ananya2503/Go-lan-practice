package main //main.go run first

import (
	"fmt"
	ifstatement "gobasic/ifStatement"
	logicoperator "gobasic/logicOperator"
	mathoperator "gobasic/mathOperator"
	"gobasic/variable"
)

// Structure
type Product struct {
	name     string
	price    float64
	category string
	discount int
}

//compile and run first
func main() {
	fmt.Println("1. Variable declaration")
	fmt.Println("2. Math Operator")
	fmt.Println("3. Logic Operator")
	fmt.Println("4. If statement")
	fmt.Println("5. Array")
	fmt.Println("6. Slice")
	fmt.Println("7. Map")
	fmt.Println("8. For loop")
	fmt.Println("9. Function return multiple value")
	fmt.Println("10. Variadic function")
	fmt.Println("11. Structure")

	// Switch case & Scanf
	var choice int
	fmt.Print("\nEnter number: ")
	fmt.Scanf("%d", &choice)
	fmt.Println()

	switch choice {
	case 1:
		variable.DataType()
		variable.VariableDeclaration()
	case 2:
		mathoperator.MathOperator()
		add := mathoperator.Add(5, 3)
		subtract := mathoperator.Subtract(5, 3)
		multiples := mathoperator.Multiples(5, 3)
		divided, _ := mathoperator.Divided(5, 3)
		mod := mathoperator.Modulo(5, 3)

		fmt.Println("5 + 3 =", add)
		fmt.Println("5 - 3 =", subtract)
		fmt.Println("5 * 3 =", multiples)
		fmt.Println("5 / 3 =", divided)
		fmt.Println("5 % 3 =", mod)
	case 3:
		logicoperator.LogicOperator()
		fmt.Println("5 == 3:", logicoperator.Equal(5, 3))
		fmt.Println("5 != 3:", logicoperator.NotEqual(5, 3))
		fmt.Println("5 > 3:", logicoperator.MoreThan(5, 3))
		fmt.Println("5 < 3:", logicoperator.LessThan(5, 3))
		fmt.Println("5 >= 3:", logicoperator.MoreThanOrEqual(5, 3))
		fmt.Println("5 <= 3:", logicoperator.LessThanOrEqual(5, 3))
	case 4:
		ifstatement.IfStatement()
		fmt.Printf("Grade is %s", ifstatement.Grade(80))
	case 5:
		array()
	case 6:
		slice()
	case 7:
		mapType()
	case 8:
		forLoop()
	case 9:
		total, status := returnMultipleValue(5, 3)
		fmt.Println("5 + 3 =", total, "and", total, "is", status)
	case 10:
		total := variadicFunction(10, 2, 5, 3)
		fmt.Println(total)
		total = variadicFunction(2, 6)
		fmt.Println(total)
	case 11:
		useStruct()
	default:
		fmt.Println("Incorrect number")
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

	// ?????????????????????????????????????????????????????? -> slice_name[start_index:end_index + 1]
	fmt.Println("Element from index 1 to 5: ", numbers1[1:])
	fmt.Println("Element from index 0 to 2: ", numbers1[:3])
	fmt.Println("Element from index 2 to 3: ", numbers1[2:4])
}

func mapType() {
	// declaration
	country := map[string]string{"TH": "Thailand", "JP": "Japan"}
	country["CN"] = "China"

	fmt.Println(country)
	fmt.Println("Key is TH:", country["TH"])

	// check value in map
	/*
		value = value in key
		check = bool return true if has key in map
	*/
	value, check := country["JP"]
	if check {
		fmt.Println(value)
	} else {
		fmt.Println("No data")
	}
}

func forLoop() {
	fmt.Println("---for loop with break and continue---")
	for i := 0; i < 11; i++ {
		if i == 9 {
			break
		}
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// for range: if want to ignore "index:, use "_"
	fmt.Printf("\n---for range---\n")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Println("index", index, "has value", value)
	}

	fmt.Printf("\n---for range (ignore index)---\n")
	language := map[string]string{"TH": "Thai", "EN": "English"}
	for _, value := range language {
		fmt.Println(value)
	}
}

func returnMultipleValue(num1, num2 int) (int, string) {
	total := num1 + num2
	status := ""

	if total%2 == 0 {
		status = "even"
	} else {
		status = "odd"
	}
	return total, status
}

// Variadic function: function ?????????????????? parameter ?????????????????????????????????????????????????????????
func variadicFunction(numbers ...int) int {
	// numbers is slice
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func useStruct() {
	pen := Product{name: "Pen", price: 20.5, category: "stationary", discount: 5}
	fmt.Println(pen)
	fmt.Println("Price is", pen.price)
}
