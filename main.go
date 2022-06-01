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
	isPass := true

	fmt.Println("My name is", name) //print with new line
	fmt.Println("Age:", age)
	fmt.Println("Score:", score)
	fmt.Println("Pass exam:", isPass)

	fmt.Printf("\nName: %v\n", name)       //%v = print value
	fmt.Printf("Type of name: %T\n", name) //%T = print type
}
