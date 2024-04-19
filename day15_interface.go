package main

import "fmt"

// func main() {
// page15_1()
// page15_2()
// }

//  ❤️介面(interface)❤️
// 不需要明確地定義實作介面，而是採用隱性實作，只要符合方法簽章即可
// 用於程式與模組之間的解藕

// 定義一個訊息發送者介面
type MessageSender interface {
	Send(content string)
}

// 定義一個簡訊發送者結構
type SmsSender struct {
	Mobile string
}

// 實作Send方法於SmsSender
func (s *SmsSender) Send(content string) {
	fmt.Println("Send:", content, "to", s.Mobile)
}

func page15_1() {
	var sender MessageSender

	sender = &SmsSender{Mobile: "0900000000"}
	sender.Send("hello,sms")

	// 我創建了一個MessageSender類型的變數sender
	// 將SmsSender的實例賦值給一個符合MessageSender類型的變數(我們通常會說一個 struct 實現了一個 interface)
	// 並通過sender調用Send方法



	//型別斷言(type assertion)
// 需要將介面轉成原來的型別用他原本的東西
smsSender, ok := sender.(*SmsSender)
fmt.Println(smsSender, ok)
//或是要檢查一個介面變數是否實現了某個介面
if ok {
    fmt.Println("sender is a *SmsSender")
} else {
    fmt.Println("sender is not a *SmsSender")
}
}


// 型別分支(type switch)

// 定義一個函式，接收任何型別，並且格式化輸出值
// interface{}(空介面) 是一種特殊的介面，它不包含任何方法，因此所有類型都實現了它

func printAnyType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("case int: %d \n", v)
	case string:
		fmt.Printf("case string: %s \n", v)
	default:
		fmt.Printf("default: %v \n", v)
	}
}

func page15_2() {
	printAnyType(2020)
	printAnyType("Iron Man")
	printAnyType(0.25)
}

