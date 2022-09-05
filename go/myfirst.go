package main // go จะให้ความสำคัญกับ package main ก่อน (เอาง่ายๆคือรันก่อน)

import "fmt"

type Product struct {
	name     string
	price    int
	discount float64
	category string
}

// Run code in func main first !!
func main() {
	product1 := Product{name: "macbook", price: 100000, discount: 10000.50, category: "Computer"}
	fmt.Println(product1.name, "ราคา", product1.name, "ส่วนลด", product1.discount)

	showMessage("Mrjane")

	// การประกาศตัวแปร
	myVariable()

	// ตัวดำเนินการทางคณิตศาสตร์
	operator()

	// if else
	ifelse()

	// Switch case
	mySwitchCase()

	// ตัวอย่างการเก็บข้อมูลใน array
	allArray()

	//map, Loop map array
	mapArray()

}

func showMessage(yourname string) {
	//fmt คือ build in func ส่วน Println คือ method
	fmt.Println("Hello !!", yourname)
}

func myVariable() {
	// variable
	var name string = "mrjane"
	var age int = 35
	var height int = 165
	fmt.Println(name, "is", age, "years old", "he height", height)

	catname := "Sukar"
	catname = "Gafeild" // เปลี่ยนค่าตัวแปลได้
	fmt.Println(catname)

	// ดูขนิดหรือประเภทตัวแปล อารมณื type of js
	fmt.Printf("Type name = %T\n", name) // %v แสดง value ออกมา \n ขึ้นบรรทัดใหม่
}

func operator() {
	// ตัวดำเนินการทางคณิตศาสตร์
	var number01 = 10
	var number02 = 3
	fmt.Println("result = ", number01+number02)
	fmt.Println("result = ", number01%number02)
}

func ifelse() {
	// if else
	score := 90
	// fmt.Print("กรุณาใส่คะแนนของคุณ") //รับ input เหมือน promt js
	// fmt.Scanf("%d", &score)

	if score >= 50 {
		fmt.Println("ยินดีด้วย คุณสอบผ่าน")
	} else {
		fmt.Println("เสียใจด้วย คุณสอบไม่ผ่าน")
	}
}

func mySwitchCase() {
	// switch case
	var number int
	fmt.Print("กรุณากรอกข้อมูล")
	fmt.Scanf("%d", &number)
	switch number {
	case 1:
		fmt.Println("Bangkok")
	case 2:
		fmt.Println("Samutprakan")
	case 3:
		fmt.Println("Chonburi")
	default:
		fmt.Println("คุณกรอกข้อมูลไม่ถูกต้อง")
	}
}

func allArray() {
	// การประกาศ Array
	/* ต้องระบุจำนวนสมาชิกใน Array และ Type of ต้อง ขนิดเดียวกันเท่านั้น*/
	var array1 [3]int = [3]int{100, 200, 300}
	fmt.Println(array1)

	array2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11} //ประกาศ Array แบบไม่จำกัดจำนวนสมาชิก (สามารถเพิ่มได้เรื่อยๆ)
	var arrayLength = len(array2)
	fmt.Println("จำนวนสมาชิกใน Array2 คือ ", arrayLength)

	// Slice เหมือน Array ทุกประการ เพียงแต่ไม่ต้องประกาศจำนวนสมาชิก !! (ไม่จำกัดสมาชิก)
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	slice1 = append(slice1, 40) // เพิ่มจำนวนสมาชิกใน slice
	fmt.Println(slice1)
	slice1 = append(slice1, 500)
	fmt.Println(slice1[:])  // เรียกสมาชิกทุกตัวใน slice
	fmt.Println(slice1[2:]) // เรียกสมาชิกตั้งแต่ตัว index ที่ 2 ถึงตัวสุดท้าย
	fmt.Println(slice1[:3]) //เรียกสมาชิกใน slice ตั้งแต่ตัวแรกถึงตัว index ที่ 3
	// Loop array2
	for i := 0; i < len(array2); i++ {
		fmt.Println("index ที่ ", i, "มีค่าเท่ากับ", array2[i])
	}
	// for loop มี loop เดียว
	for count := 1; count <= 10; count++ {
		fmt.Println(count)
		if count == 5 {
			//ใส่เงื่อนไข
			fmt.Println("เข้าเงื่อนไข count == 5")
			break // เมื่อเจอ Break จะออกจาก loop แต่ถ้า continue จะข้ามรอบนั้นไป (ทวน)
		}
	}
	fmt.Println("จบโปรแกรม")

}

func mapArray() {
	// map = การกำหนดให้เรียกสมาชิกด้วย Key กำหนดสมาชิก key:value, ...

	exampleMap := map[string]string{"TH": "Thailand", "US": "America", "EN": "England"}
	fmt.Println(exampleMap)       //map[EN:England TH:Thailand US:America]
	fmt.Println(exampleMap["EN"]) // England ค้นหา value ด้วย key

	// Loop map array
	for key, value := range exampleMap {
		fmt.Println("Key = ", key, "value =", value)
	}
}
