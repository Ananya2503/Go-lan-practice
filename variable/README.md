# Variable
## Data type
|Data type|Description|Zero value (default value)|
|:-------:|:---------:|:------------------------:|
|bool|Logic value (true/false)|false|
|int (8, 16, 32, 64)<br />unit (8, 16, 32, 64): ไม่คิดเครื่องหมาย<br />uintprt|integer|0|
|float32 float64|floating number|0|
|string|messege/set of characters|""|

## Variable
- static-type: ประกาศตัวแปรก่อนเรียกใช้งาน
- การนิยามมี 2 รูปแบบ
    1. **Manual Type Declaration**: ประกาศแบบระบุ data type มีโครงสร้าง คือ var {var name} {data type}
    2. **Type Inference**: ประกาศแบบไม่ระบุ data type มีโครงสร้าง คือ {var name} := {value}
- constant variable: ตัวแปรที่ไม่สามารถเปลี่ยนค่าได้