# Go-lan-practice

## Start project
```
go mod init _module name/project name_
```

## Run project
```
go run _file name_
```

## Structure
- func main(): กลุ่มคำสั่งในฟังก์ชันนี้จะถูกทำงานเป็นลำดับแรกเสมอ
- {} ใช้บอกขอบเขตการทำงานของฟังก์ชัน
- ใช้ // หรือ /**/ เพื่อ comment
- import "package_name": เรียกใช้คำสั่งพื้นฐานใน package_name

## Data type
|Data type|Description|Zero value (default value)|
|:-------:|:---------:|:------------------------:|
|bool|Logic value (true/false)|false|
|int (8, 16, 32, 64)<br />unit (8, 16, 32, 64): เลขจำนวนเต็มแบบไม่คิดเครื่องหมาย<br />unitprt|integer|0|
|float32 float64|floating number|0|
|string|messege/set of characters|""|

## Package
- fmt: จัดการเรื่องรูปแบบ input/output

## ข้อสังเกต
- เมื่อ import package เข้ามาแล้วต้องมีการเรียกใช้เท่านั้น
- ไม่จำเป็นต้องใส่ ;

## Credit
Thank you for YouTube channel [KongRuksiam Official](https://youtu.be/pytqhPDTjnQ)