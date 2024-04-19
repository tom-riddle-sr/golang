package main

import "fmt"

// 宣告一個名叫a的整數變數
var a int

// 宣告一個名叫b的字串變數
var b string

// 宣告一個名叫c的浮點數陣列
var c []float32

// 宣告一個名叫d的方法變數並且回傳bool型別
var d func() bool

// 若型別相同，可以這樣寫
var e, f, g int

// 若型別不同，可以這樣寫
var (
	h int
	i string
	j []float32
	k func() bool
)

// 函式以外的地方宣告變數、方法、結構等等，只要是名稱開頭是大寫的，表示可以被匯出(exported)

var year int = 2020
var name string = "Iron man"

// 短變數宣告:變數精簡寫法,同時宣告和初始化,只能在函式內使用
func main1() {
	year := 2020
	name := "Iron man"

	fmt.Println(year,name)

}

func main2() {
	var a, b = 1, 2

	// 傳統swap寫法
	var temp int
	temp = a
	a = b
	b = temp

	fmt.Println(a, b) // 這將輸出 "2 1"
}

// 匿名佔位符(anonymous placeholder)
// 忽略第二個return值
func getHeroName() (string, string) {
	return "Iron man", "Dr.strange"
}

func main3() {
	var name, _ = getHeroName()
	fmt.Println("I like " + name)
}
