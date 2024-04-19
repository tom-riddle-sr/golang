package main

import "fmt"

// func main(){
// page8_1()
// page8_2()
// }

func page8_1(){
	// var arr[3]string
// var arr = [3]string{"Iron Man", "Dr.Stange", "Thor"} //指定初始值
// 如果陣列有設定初始值時，可以省略陣列長度
var arr = [...]string{"Iron Man", "Dr.Stange", "Thor"} //指定初始值
arr[0]="Jerry"
arr[1]="Alice"
arr[2]="Fifi"



	fmt.Println(arr)
	fmt.Println(len(arr))
}

// foreach Array
func page8_2(){
	arr := [...]int{10, 20, 30, 40, 50}
	for i,v := range arr {
		fmt.Println(i,v)
	}
}