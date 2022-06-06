# Go-lan-practice

## Start project
```
go mod init {module name/project name}
```

## Run project
```
go run {file name}
```

## Programming Structure
- func main(): กลุ่มคำสั่งในฟังก์ชันนี้จะถูกทำงานเป็นลำดับแรกเสมอ
- {} ใช้บอกขอบเขตการทำงานของฟังก์ชัน
- ใช้ // หรือ /**/ เพื่อ comment
- import "package_name": เรียกใช้คำสั่งพื้นฐานใน package_name

## Data type
|Data type|Description|Zero value (default value)|
|:-------:|:---------:|:------------------------:|
|bool|Logic value (true/false)|false|
|int (8, 16, 32, 64)<br />unit (8, 16, 32, 64): ไม่คิดเครื่องหมาย<br />unitprt|integer|0|
|float32 float64|floating number|0|
|string|messege/set of characters|""|

## Variable
- static-type: ประกาศตัวแปรก่อนเรียกใช้งาน
- การนิยามมี 2 รูปแบบ
    1. **Manual Type Declaration**: ประกาศแบบระบุ data type มีโครงสร้าง คือ var {var name} {data type}
    2. **Type Inference**: ประกาศแบบไม่ระบุ data type มีโครงสร้าง คือ {var name} := {value}
- constant variable: ตัวแปรที่ไม่สามารถเปลี่ยนค่าได้

## Math Operator
|Operator|Description|
|:------:|:---------:|
|+|add|
|-|minus|
|*|times|
|/|divided|
|%|modulo|

## Logic Operator
|Operator|Description|
|:------:|:---------:|
|==|equal|
|!=|not equal|
|>|more than|
|<|less than|
|>=|more than or equal|
|<=|less than or equal|

## Array
- Element in array = same type
- index start from 0
- fixed size
- can use ... when array is infinite size
```
// Manual type declaration
var var_name [size] element_type = [size] element_type {element}

// Type inference
var_name := [size] element_type {element}
```

## Slice
- คล้ายกับ Array
- index start from 0
- สามารถเข้าถึง element แบบระบุช่วงได้
- dynamic size = สามารถเพิ่ม element ได้ด้วย append()
```
// Manual type declaration
var var_name [] element_type = [] element_type {element}

// Type inference
var_name := [] element_type {element}
```

## Map
- เก็บข้อมูลแบบคู่ (key, value)
- ใช้ key ในการเข้าถึงข้อมูล
- key กับ value สามารถมี variable type คนละชนิดได้
```
map_name := map[key_type] value_type {key:value}
```

## Break & Continue
- break: หยุดและออกจาก loop ทันที
- continue: หยุดการทำงานในรอบนั้น แล้วเริ่มต้นรอบใหม่

## Structure (Struct)
- นำข้อมูลที่มีชนิดต่างกันมารวมเข้าไว้ด้วยกัน
- สามารถสร้างชนิดตัวแปรของตัวเองได้
```
type struct_name struct {
    var1 var1_type;
    var2 var2_type;
}
```

## Package
- fmt: จัดการเรื่องรูปแบบ input/output

## ข้อสังเกต
- เมื่อ import package เข้ามาแล้วต้องมีการเรียกใช้เท่านั้น
- เมื่อนิยามตัวแปรแล้วต้องเรียกใช้เสมอ
- ไม่จำเป็นต้องใส่ ;

## Credit
Thank you for YouTube channel [KongRuksiam Official](https://youtu.be/pytqhPDTjnQ)