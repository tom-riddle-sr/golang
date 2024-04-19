package main

import (
	"fmt"
	"math"
	"unsafe"
)

// func main() {
// day4_1()
// day4_2()
// day4_3()
// day4_4()
// day4_5()
// day4_6()
// }

// "1 byte = 8 bits" 是描述電腦數據存儲的基本單位的一種方式。

// ❤️整數❤️
// 整數型別可分為2類
// 初始值0
// -有符號:int8 int16 int32 int64 正數,負數,0
// -無符號:unit8 uint16 uint32 uint64 正數,0

//int的Max值<unit, 因為最左邊的bit是符號位元(正,負)

// 將 1 左移 7 位元再 - 1 (不 - 1 會錯誤)
// unsafe.Sizeof(a)取得記憶體大小
func day4_1(){
var a int8 = 1<<7 - 1
fmt.Printf("int8 長度= %d byte, 最大值= %d\n\n", unsafe.Sizeof(a), a)
}

// ❤️浮點數❤️
// 浮點數型別可分為2類
// 初始值0
// float32
// float64

// %d：用於格式化整數並以十進制方式輸出
// %f：用於格式化浮點數，並以十進制方式輸出
// %e：用於格式化浮點數，並以科學記數法輸出
// %T：用於顯示變數的型別

func day4_2(){
var b float32=math.MaxFloat32
fmt.Printf("float32 長度= %d byte\n最大值(科學記號法)= %e\n最大值(十進制)= %f\n\n", unsafe.Sizeof(b), b, b)}

// ❤️布林❤️
// 初始值false

func day4_3(){
 ded := true
fmt.Println(ded)
}

// ❤️字串❤️
// 初始值""
// go語言的字元是用整數來表達,只有字串型別沒有字元型別

func day4_4(){
 string1:="西西cute"
//  用反引號(`)定義多行文字
 string2:=`西西cute
 smart
 talented`
	fmt.Println(string1,string2)
}

//  ❤️型別的轉換❤️
func day4_5(){
	// 宣告一個16位元整數，轉成8位元時發生截斷，因為8位元整數最大值是127
var a int16 = 128
fmt.Printf("%d \n", int8(a))

// 將浮點數轉成整數，會損失小數點後的精度
var b float32 = math.Pi
fmt.Printf("整數: %d, 浮點數: %f \n", int(b), b)
}

// ❤️型別的別名(Type Alias)❤️
// type關鍵字本身就是用於定義型別,別名的語法必須加上等號( = )來區分這是在定義別名
// go語言本身也有內建的型別別名
// type byte = uint8
// type rune = int32

func day4_6(){
	type IntAlias = int

	var a byte
	fmt.Printf("%T \n", a)

	var b rune
	fmt.Printf("%T \n", b)

	var c IntAlias
	fmt.Printf("%T \n", c)
}