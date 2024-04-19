package main

import "fmt"

// func main() {
// page11_1()
// page11_2()
// page11_3()
// page11_4()
// }

// 物件導向風格的語言
// 繼承雖然有很好的共享實作及多型的特性，
// 但它也帶來了不少問題，因為繼承類別屬於侵入式依賴，
// go語言主張以組合和介面來取代繼承
// ❤️結構(struct)❤️


func page11_1() {
	type Hero struct {
	name string
	age, power  int
}

	var andy Hero
fmt.Println(andy)

// 宣告一個英雄，並初始化所有欄位
var tony = Hero{"IronMan", 30, 666}
fmt.Println(tony)

// 宣告一個英雄，並初始化部份欄位，其他欄位使用預設值要明確指定是宣告哪個值
var stephen = Hero{name: "Dr.Strange"}
fmt.Println(stephen)
}

// 結構與指標
// 避免資料不會被修改的問題
func page11_2(){
	type Hero struct{
		name string
		age,power int
	}

var a = &Hero{"IronMan", 30, 666} //
fmt.Println(*a)

a.name="Loti"
a.age=20
a.power=999

fmt.Println(*a)

}

// new( )
// 效果就和 &結構名 一樣,創建一個新的指標,但是用 new 就不能同時初始化結構裡的欄位
func page11_3(){
	type Hero struct{
		name string
		age,power int
	}
a:=new(Hero)

a.name="IronMan"
fmt.Println(*a)

}

// 匿名結構
// 不用事先定義好結構型別，在變數宣告時產生一次性的結構型別
func page11_4(){

	stephen := &struct {
    name       string
    age, power int
}{name: "Dr.Strange"}

fmt.Println(*stephen)

}