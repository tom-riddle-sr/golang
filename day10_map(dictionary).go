package main

// 映射(map)在一些語言裡稱為字典(dictionary) //有點像{}
// 指的是以鍵-值(key-value)存取的資料容器
// 映射也是一種型別

import "fmt"

// func main() {
// page10_1()
// page10_2()
// page10_3()
// page10_4()
// }

func page10_1(){
var a= map[string]int{"aa":0,"bb":1}
fmt.Printf("len=%d, value=%v \n", len(a), a)
}

func page10_2(){
var arr=map[string]int{"aa":0,"bb":1,"cc":2}
aa:=arr["aa"]
fmt.Println(aa,len(arr))

for i,v:=range arr{
	// 映射巡覽的順序不一定會和初始值或資料插入的順序一致，因為它的索引值是不連續的。
	fmt.Println(i,v)
}

// 檢查是否存在
 ii,vv:=arr["dd"]

if vv{
	fmt.Println(ii)
}else{
	fmt.Println("不存在")
}
}

func page10_3(){
	// 動態建立映射
	var arr = make(map[string]int)
	arr["aa"]=0
	fmt.Println(arr)
	// 刪除映射
	// 映射裡的資料是分散的,無法像陣列可以用元素覆蓋的方式刪除資料
	// go語言提供了delete() 函式
	// *刪除不存在的key不會造成錯誤
	delete(arr,"aa")
	fmt.Println(arr)
}

// 映射裡可以放函式
func page10_4(){
map1:=make(map[string]func(a ,b int)int)

map1["add"]=func(a,b int)int{
	return a+b
}

map1["minor"]=func(a,b int)int{
	return a-b
}

map1["multiply"]=func(a,b int)int{
	return a*b
}

map1["divide"]=func(a,b int)int{
	return a/b
}

for i,v:=range map1{
	fmt.Printf("%s=%d\n",i,v(10,2))
}
}