package main

import "fmt"

// func main(){
// page6_1()
// page6_2()
// page6_3()
// page6_4()
// }

// 流程控制分為2種:
// -分支語法
// -循環語法

// ❤️分支語法❤️
// 分支是指用條件來選擇執行某一段程式碼
// if/else 和 switch case

// if/else
func page6_1(){
myAge := 30
if myAge < 13 {
    fmt.Println("Child")
} else if myAge >= 13 && myAge < 20 {
    fmt.Println("Teen")
} else {
    fmt.Println("Adult")
}
}

// go語言的 if 還有一個特殊寫法，可以在條件之前加上一行表達式
func page6_2(){
myAge := 30
if myAge = myAge - 15; myAge < 20 {
    fmt.Println("Teen")
} else {
    fmt.Println("Adult")
}
}

// switch case
// 除了case帶條件以外，還可以帶多個值
func page6_3(){
myAge := 30
switch {
case myAge < 13:
    fmt.Println("Child")
case myAge >= 13 && myAge < 20:
    fmt.Println("Teen")
case myAge >= 20 && myAge < 50:
    fmt.Println("Adult")
default:
    fmt.Println("Elder")
}
}


// ❤️循環語法❤️
// for loop和遞迴

//for loop
//go 只有for loop，沒有while loop
func page6_4(){
	for i := 1; i <= 9; i++ {
    for j := 1; j <= i; j++ {
    	fmt.Printf("%d*%d=%d ", i, j, i*j)
    }
    fmt.Println()
}
}