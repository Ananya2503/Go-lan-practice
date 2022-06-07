package ifstatement

import "fmt"

func IfStatement() {
	fmt.Println("***** If statement structure *****")
	fmt.Println("if condition {")
	fmt.Println("   statement")
	fmt.Println("} else if condition {")
	fmt.Println("   statement")
	fmt.Println("} else {")
	fmt.Println("   statement")
	fmt.Println("}")
	fmt.Println()
}

func Grade(score int) string {
	if score >= 80 {
		return "A"
	} else if score < 80 && score >= 71 {
		return "B"
	} else if score < 70 && score >= 61 {
		return "C"
	} else if score < 60 && score >= 51 {
		return "D"
	} else {
		return "F"
	}
}
