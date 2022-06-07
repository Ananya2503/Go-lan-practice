package variable

import "fmt"

func DataType() {
	fmt.Println("***** Data type *****")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("|     Data Type        |                Description               | Zero value |")
	fmt.Println("|------------------------------------------------------------------------------|")
	fmt.Println("| bool                 | Logic value                              |    false   |")
	fmt.Println("|------------------------------------------------------------------------------|")
	fmt.Println("| int (8, 16, 32, 64)  | Integer                                  |      0     |")
	fmt.Println("| uint (8, 16, 32, 64) | Unsigned integer                         |      0     |")
	fmt.Println("| uintprt              | Integer representation of memory address |      0     |")
	fmt.Println("|------------------------------------------------------------------------------|")
	fmt.Println("| float32 float64      | Floating number                          |      0     |")
	fmt.Println("|------------------------------------------------------------------------------|")
	fmt.Println("| string               | Message                                  |     \"\"     |")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()
}

func VariableDeclaration() {
	fmt.Println("***** Variable declaration *****")
	fmt.Println("เป็น static-type ต้องประกาศตัวแปรก่อน จึงจะสามารถเรียกใช้งานได้")
	fmt.Println()

	fmt.Println("1. Manual type declaration")
	fmt.Println("structure: var var_name var_type = value")
	fmt.Println("Example: var country string = \"Thailand\"")
	var country string = "Thailand"
	fmt.Println("Value in \"country\":", country)
	fmt.Println("----------")
	fmt.Println("structure (more than 1 var): var var_name1, var_name2 var_type = value1, value2")
	fmt.Println("Example: var age, height int = 22, 164")
	var age, height int = 22, 164
	fmt.Println("Value in \"age\":", age)
	fmt.Println("Value in \"height\":", height)
	fmt.Println()

	fmt.Println("2. Type inference")
	fmt.Println("structure 1: var_name := value")
	fmt.Println("Example: score := 81.4")
	score := 81.4
	fmt.Println("Value in \"score\":", score)
	fmt.Println("----------")
	fmt.Println("structure 2: var var_name = value")
	fmt.Println("Example: var isPass = true")
	var isPass = true
	fmt.Println("Value in \"isPass\":", isPass)
	fmt.Println()

	fmt.Println("-- Constant variable เป็นค่าคงที่ที่ไม่สามารถเปลี่ยนค่าได้ --")
	fmt.Println("structure: const var_name var_type = value")
	fmt.Println("Example: const name string = \"Ananya\"")
	const name string = "Ananya"
	fmt.Println("Const variable \"name\":", name)
	fmt.Println()

	fmt.Println("Check variable type")
	fmt.Println("use fmt.Printf(\"% T\"), var")
	fmt.Printf("Type of \"name\": %T", name)
	fmt.Println()
}
