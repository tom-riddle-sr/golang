package main

import "fmt"

func main() {
// page16_1()
page16_2()
}

// go語言的函式是可以當作變數使用
// 函式也是一種型別，函式的一種進階用法
// ❤️閉包(Closures)❤️
// 函式變數
// 函式中使用外部環境的區域變數
// 延長變數的生命週期
func page16_1(){
	var foo = func() func() {
		// 為了確保i可以保存修改值並不被外部影響
    var i int
    return func() {
        i++
        fmt.Println(i)
    }
}()
foo()
foo()
foo()
}

// 閉包與結構
// 閉包可以做很類似物件導向的事，可以同時封裝資料和行為
// 閉包不一定要回傳函式，也可以回傳結構

func page16_2(){
	type counter struct {
		add   func()
		minus func()
		print func()
	}

	count := func() counter {
    i := 0
    return counter{
        func() {
            i++
        },
        func() {
            i--
        },
        func() {
            fmt.Println("i =", i)
        },
    }
}()

count.add()
count.add()
count.minus()
count.print()

}