package main

import "fmt"

// func main (){
// page7_1()
// page7_2()
// page7_3()
// page7_4()
// }

// 如果要在程式裡，想要動態配置記憶體的話，都要使用指標來存放位置
// 用來操作記憶體空間，可以透過記憶體位址直接存取資料，取代資料搬移和複製的動作，使程式擁有較高的執行效能

// ❤️指標變數❤️
// 可以當作變數使用的都是一種型別
// %p 用來顯示記憶體位址
// nil  空指標, 沒有指向任何記憶體位址
// *2種用法:
// 1. 當它出現在一個類型前面時，它表示一個指標類型
// 2. 取出指標變數所指向的記憶體位址的值

func page7_1(){
	str := "Iron Man"
	var ptr *string  // 宣告一個指向字串的指標變數
	ptr = &str  //取出str的記憶體位址給ptr
	fmt.Printf("%p \n", ptr)
	fmt.Printf("%s \n", *ptr)
}

// 如果只想修改指向的值而非原本的值呢


// swap函式
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}


func page7_2() {
	a, b := 1, 2
	swap(&a, &b)
	println(a, b)
}

//new()
func page7_3(){
	ptr := new(string)  // 動態配置一個字串型別的的變數,記憶體位置不固定

*ptr = "Iron Man"

fmt.Printf("%s \n", *ptr)
}



func page7_4(){
	var a int =10
	fmt.Println(a, &a)
	//值,a的記憶體位址
	var b*int=&a
	fmt.Println(b,*b, &b)
	//a的記憶體位址,a的值,b的記憶體位址
	fmt.Printf("%p = %p\n",&a,b)
	fmt.Printf("%d = %d ",a,*b)

	}

