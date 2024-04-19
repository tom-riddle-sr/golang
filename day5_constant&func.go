package main

import "fmt"

// func main(){
// page5_1()
// page5_2()
// page5_3()
// page5_4()
// page5_5()
// page5_6()
// page5_7()
// }

// ❤️常數❤️
const (
	pi=3.1415926
	e1=2.71828
)
// 常數可以用表達式指定內容，以及在陣列宣告時指定大小
func page5_1(){
const pageSize = 5
const totalPage = 20
const totalSize = pageSize * totalPage

var arr [totalSize]int
fmt.Println(len(arr),arr)
//印出: 100
}

// 列舉（Enum）是一種特殊的資料類型，它允許你定義一組命名的整數常數
// go語言中沒有定義列舉(enum)的功能，但是可以用常數和iota關鍵字達到相似的效果
// iota關鍵字是一個常數計數器，第一個常數從 0 開始，往下逐漸的累加。而iota也可以加入表達式
func page5_2(){
type Weekday int
const (
	Sunday Weekday = iota*2+1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

// ❤️函數❤️
// 函式有4個部份：函式名稱 參數 回傳值 主體
// 第二個 string 是這個函數的return type
func foo(a int, b string) string {
	return fmt.Sprintf("%d %s", a, b)// %s 字符串格式指定符
	// 根據格式字符串和參數列表生成一個新的字符串，但不會將這個字符串輸出到螢幕
}

func page5_3() {
	s := foo(2020, "Iron Man")
	fmt.Println(s)
}
// go函式的回傳值可宣告多個，但要用小括號( )包起來，在函式呼叫的左側就必須放上相同數量的變數

func foo1(c int, d string)(int ,string){
return c,d
}
func page5_4(){
	_,b:=foo1(2024,"cc")
	fmt.Println(b)
}

// Println&Printf的區別
// Println: 會自動換行
// Printf: 不會自動換行,需要格式化指定符

// go函式還有一個特色，就是回傳值可以命名
// 被命名的回傳值就等同於宣告一個變數，在函式結束時用return可以自動回傳變數裡的內容

func foo3(a int ,b string)(c string, d int ){
c=b+b
d=a
return
}

func page5_5(){
	fmt.Println(foo3(1,"cc"))
}

// ❤️函式變數❤️
// go函式屬於一級公民(first class)
// -被當作參數傳遞
// -被賦值給變數
// -作為返回值
// -在執行時創建
// -函式也是一種資料型別

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func page5_6() {
	var foo func(a, b int) int

	foo = minus
	foo = add

	a := foo(1, 2)
	fmt.Printf("%d\n", a)
}

// ❤️匿名函式❤️
// foo函式接收2個整數和一個函式變數
func foo4(a, b int, f func(a, b int) int) int {
	return f(a, b)
}

func page5_7() {
	var add = func(a, b int) int {
		return a + b
	}

	a := foo4(1, 2, add)
	fmt.Printf("%d\n", a)
}