package main

import (
	"fmt"
)

func main() {
	// คำสั่ง print และการตรวจสอบชนิดของข้อมูล
	// 1. การใช้ fmt.Print() และ fmt.Println()
	fmt.Print("Hello, ")
	fmt.Println("World!")

	println("--------------------")

	// 2. การตรวจสอบชนิดของข้อมูล
	// 3. การใช้ fmt.Printf()
	fmt.Printf("Hello, %s\n", "World!") // %s คือ format specifier สำหรับ string
	fmt.Printf("Hello, %d\n", 100)      // %d คือ format specifier สำหรับ integer
	fmt.Printf("Hello, %f\n", 100.123)  // %f คือ format specifier สำหรับ float
	fmt.Printf("Hello, %t\n", true)     // %t คือ format specifier สำหรับ boolean
	fmt.Printf("Hello, %v\n", "World!") // %v คือ format specifier สำหรับ value ใดๆ ที่เป็นไปได้

	println("--------------------")

	// 4. การใช้ fmt.Sprintf()
	_ = fmt.Sprint("Hello, ", "World!") // คำสั่งนี้จะ return ค่าเป็น string แต่ไม่ได้ใช้ผลลัพธ์

	s := fmt.Sprintf("Hello, %s", "World!") // คำสั่งนี้จะ return ค่าเป็น string และใช้ผลลัพธ์

	fmt.Println(s)

	println("--------------------")

	// รหัสของรูปแบบต่างๆ ของ format specifier
	// %v คือ format specifier สำหรับ value ใดๆ ที่เป็นไปได้
	// %T คือ format specifier สำหรับชนิดของข้อมูล
	// %t คือ format specifier สำหรับ boolean
	// %d คือ format specifier สำหรับ integer
	// %b คือ format specifier สำหรับ binary
	// %o คือ format specifier สำหรับ octal
	// %x คือ format specifier สำหรับ hexadecimal
	// %X คือ format specifier สำหรับ hexadecimal ใหญ่
	// %f คือ format specifier สำหรับ float
	// %e คือ format specifier สำหรับ scientific notation
	// %E คือ format specifier สำหรับ scientific notation ใหญ่
	// %s คือ format specifier สำหรับ string
	// %q คือ format specifier สำหรับ quoted string
	// %p คือ format specifier สำหรับ pointer
	// %c คือ format specifier สำหรับ character
	// %U คือ format specifier สำหรับ Unicode
	// %v คือ format specifier สำหรับ value ใดๆ ที่เป็นไปได้

	// example
	fmt.Printf("%v\n", 100)             // 100
	fmt.Printf("%T\n", 100)             // int
	fmt.Printf("%t\n", true)            // true
	fmt.Printf("%d\n", 100)             // 100
	fmt.Printf("%b\n", 100)             // 1100100
	fmt.Printf("%o\n", 100)             // 144
	fmt.Printf("%x\n", 100)             // 64
	fmt.Printf("%X\n", 100)             // 64
	fmt.Printf("%f\n", 100.123)         // 100.123000
	fmt.Printf("%e\n", 100.123)         // 1.001230e+02
	fmt.Printf("%E\n", 100.123)         // 1.001230E+02
	fmt.Printf("%s\n", "Hello, World!") // Hello, World!
	fmt.Printf("%q\n", "Hello, World!") // "Hello, World!"
	fmt.Printf("%p\n", &s)              // 0xc0000b6010
	fmt.Printf("%c\n", 100)             // d
	fmt.Printf("%U\n", 100)             // U+0064
}
