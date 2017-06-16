package main

import (
	"fmt"
)

type phone interface {
	call()
	send(message string)
	sell()
}

type Iphone struct {
	iphoneName int64
	price float32
}

type Androidphone struct {
	androidName string
}
func (iphone Iphone) call(phoneNum int64) (n int){
	fmt.Printf("%d 呼叫 %d\n" , iphone.iphoneName, phoneNum)
	return 1
}
func (iphone Iphone) send(message string) {
	fmt.Printf("%d send message %s\n", iphone.iphoneName, message)
}
func (iphone Iphone) sell() {
	fmt.Println(iphone.price)
}

func main() {
	var p Iphone = Iphone{18823230296,4800}
	fmt.Println(p.call(13439590465))
	p.sell()
	p.send("我现在很忙")
}

