package main

import "fmt"

// func main() {
// 	page14_1()
// page14_2()
// page14_3()
// }

// 結構該如何共享程式碼
// go語言沒有繼承的特性，但我們能用組合的方式來共享程式碼
// go語言還提供一種優於組合的語法特性，稱作內嵌。

// ❤️組合(composition)❤️
// 多用組合少用繼承

// Is-A: 繼承關係，表示一個物件也是另一個物件。
// Has-A: 組合關係，表示一個物件擁有另一個物件。

// %+v用於輸出一個值的所有欄位
func page14_1(){
type Person struct{
	Name string
	}
type Hero struct{
	Person *Person
	HP ,Rank int
}
var a =&Hero{&Person{"Iron"}, 100, 1}
fmt.Printf("%+v\n", a)
}

// ❤️內嵌(embedding)❤️
// 內嵌允許我們在結構內組合其他結構時，不需要定義欄位名稱
// 內嵌的結構欄位還是會有名稱，就是和結構本身的名稱同名
// 直接透過該結構叫用欄位或方法

func page14_2(){
type Person struct{
	Name string
}


type Hero struct{
*Person
age,power int

}


var a =&Hero{
	&Person{"cc"},23,23,
}

fmt.Print(a.Name)
fmt.Print(a)

}

// 內嵌與方法
type Hero struct {
	*Person
	HeroName string
	HeroRank int
}

// 英雄都會飛
func (*Hero) Fly() {
	fmt.Println("I can fly.")
}

// 定義一個正常人結構
type Person struct {
	Name string
}

// 正常人會走路
func (*Person) Walk() {
	fmt.Println("I can walk.")
}

func page14_3(){
	// 定義一個英雄結構
var tony = &Hero{
		Person:   &Person{"Tony Stark"},
		HeroName: "Iron Man",
		HeroRank: 1}

	tony.Walk()
	tony.Fly()

}

// 只要是內嵌的都可以省略名稱直接使用方法
// 2個內嵌有相同名稱欄位時 要明確指定
// 可以被內嵌的型別不只有結構，也可以是基本型別
// 基本型別被內嵌之後，欄位名稱就是型別的原始名稱，ex: int, string, ...。