package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	id := uuid.New()
	fmt.Println(id)
	fmt.Println(id.String())
	fmt.Printf("%s\n", id)

	// การรับค่าจากผู้ใช้ Scan, Scanln, Scanf
	// 1. การใช้ fmt.Scan()
	// func Scan(a...any) (n int, err error) // รับค่าแบบ space หรือ enter
	// n คือ จำนวนตัวอักษรที่รับเข้ามา
	// err คือ ข้อผิดพลาด
	var name string // ประกาศตัวแปร name แบบ string

	fmt.Print("Enter your name: ") // แสดงข้อความ Enter your name:
	_, err := fmt.Scan(&name)      // รับค่าจากผู้ใช้ และเก็บไว้ในตัวแปร name
	if err != nil {                // ถ้ามีข้อผิดพลาด
		fmt.Println(err) // แสดงข้อผิดพลาด
	} else { // ถ้าไม่มีข้อผิดพลาด
		fmt.Println("Your name is", name) // แสดงข้อความ Your name is และค่า name
	}

	// 2. การใช้ fmt.Scanln()
}
