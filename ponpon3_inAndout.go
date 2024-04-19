package main

import "fmt"

// func main(){
// 	pon3_1()
// }

func pon3_1(){
	var a int
	var b int
fmt.Print("請輸入2個數字：")
fmt.Print("1.")
fmt.Scanln(&a)
fmt.Print("2.")
fmt.Scanln(&b)
fmt.Printf("%d+%d=%d\n", a, b, a+b)

var c int
var d int
fmt.Print("請輸入第二組2個數字： 用空格格開")
fmt.Scanln(&c, &d)
var result int = c + d
fmt.Printf("%d+%d=%d\n", c, d, result)
}