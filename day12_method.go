package main

import "fmt"

// func main() {
// page12_1()
// page12_2()
// }
// ❤️method❤️
// 雖然go語言沒有類別，但方法也不是寫在結構內，而是寫在外面
// 並且會透過接受者(receiver)來指定方法給一個結構型別

// 	type Person struct{
// 		name string
// 		blood ,attackValue int

// 		}

// func (p1 *Person) Attack(p2 *Person){
// finalBlood:=p2.blood-p1.blood
// fmt.Printf("%s攻擊%s, %s剩下%d血量\n", p2.name, p1.name, p1.name, finalBlood)

// }


func page12_1(){
// var Loti = &Person{"Loti", 80, 80}
// var Vivi = &Person{"Vivi", 100, 100}


// Loti.Attack(Vivi)
}

// 接收者
// this 或 self

// 其他型別與方法
// go語言中除了結構以外，其他的型別也是可以定義方法。
// 但是go規定接收者只能是同個package裡的型別
// 像int string 是內建package的型別，所以不能定義方法
// 因此我們就必須使用型別別名


// 定義浮點數型別的別名
type MyFloat float64
func (a *MyFloat) Add(b float64) {
	*a = *a + MyFloat(b)
}
func page12_2(){
	var a = MyFloat(2.2)
	a.Add(1.1)
	fmt.Println(a)
}

//Pass by value
//Pass by pointer
// 參考pointer func 7-4
