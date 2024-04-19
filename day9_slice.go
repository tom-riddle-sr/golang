package main

import "fmt"

// ❤️切片❤️
// 對於陣列的抽象介面
// 過程中並不會產生新的資料，因為切片只是一個介面
// 當我們對切片修改資料時，也會修改原始陣列的資料

// func main(){
// page9_1()
// page9_2()
// page9_3()
// page9_4()
// page9_5()
// page9_6()
// }

func page9_1(){
var arr =[...]string{"Vivi","Jerry","Alice","Fifi","Tom"}

slice:=arr[1:4] // [1:4] 代表從陣列的第1個元素開始，到第4個元素結束，不包含第4個元素

for i,v:=range slice{
	fmt.Printf("%d %s",i,v )
	fmt.Println(len(slice) )
}

}

// 切片省略

func forFunc (aa []string){
for i,v:=range aa{
	fmt.Printf("%d %s",i,v)
	fmt.Println()
}

}

func page9_2(){
arr:=[...]string{"Vivi","Jerry","Alice","Fifi","Tom"}

a:=arr[:]
b:=arr[:3]

forFunc(a)
fmt.Println("-----------")
forFunc(b)
}

// 切面內部
// 包含 開始位置、長度(length)、容量(capacity)
// 切片會從開始位置找到連續空間中的一個記憶體位址，然後用長度存取固定資料，而容量則表示還能再切多少元素
// 長度定義: 從 切片開頭位置 到 切片結束位置 的長度。
// 容量定義: 從 切片開頭位置 到 原始陣列結束位置 的長度。

func page9_3(){
arr:=[...]string{"Vivi","Jerry","Alice","Fifi","Tom","Bob","Cindy","David"}

a:=arr[1:6]
fmt.Println(a,len(a),cap(a))
// 切片可以再切
a=a[2:]
fmt.Println(a,len(a),cap(a))

}

//切片宣告
func page9_4(){
// 宣告一個字串切片
var strSlice []string
fmt.Printf("len=%d, cap=%d, values=%v \n", len(strSlice), cap(strSlice), strSlice)

// 宣告一個整數切片
var intSlice []int
fmt.Printf("len=%d, cap=%d, values=%v \n", len(intSlice), cap(intSlice), strSlice)

// 宣告並初始化一個字串切片
strSlice2 := []string{"Iron Man"}
fmt.Printf("len=%d, cap=%d, values=%v \n", len(strSlice2), cap(strSlice2), strSlice2)

// 宣告並初始化一個整數切片
intSlice2 := []int{2019, 2020}
fmt.Printf("len=%d, cap=%d, values=%v \n", len(intSlice2), cap(intSlice2), intSlice2)
}

//動態切片
// 可以用 make( ) 動態建立一個切片
// 建立時可以指定長度和容量，格式有2種:
// ([]T, len, cap)
func page9_5(){
// 格式1:長度 容量相同
a:=make([]string,2)
a[1]="Chase"
fmt.Println(a)
// 格式2:長度 容量不同
b:=make([]string,2,5)
b[1]="sdsd"
fmt.Println(b)
}

// 切片擴增
// 當我們要動態增加資料時，可以用 append( ) 來附加一個新元素
// 如果現有空間不足時，切片就會自動擴增容量

func page9_6(){
aa:=make([]string,0)
fmt.Printf("len=%d, cap=%d, value=%v\n", len(aa), cap(aa),aa)
aa=append(aa, "pp")
fmt.Printf("len=%d, cap=%d, value=%v\n", len(aa), cap(aa),aa)
aa=append(aa, "qq")
fmt.Printf("len=%d, cap=%d, value=%v\n", len(aa), cap(aa),aa)
aa=append(aa, "rr")
fmt.Printf("len=%d, cap=%d, value=%v\n", len(aa), cap(aa),aa)



}
