package main

import "fmt"

func main() {
	//位运算符
	var a uint = 60
	var b uint = 13
	var c uint = 0

	c = a & b
	fmt.Printf("%d\n",c)

	c = a << 2
	fmt.Printf("%d\n", c)

	//逻辑运算符
	var aa bool = true
	var bb bool = false
	if(aa && bb) {
		fmt.Println("aa && bb = true")
	}
	if(aa || bb) {
		fmt.Println("aa || bb = true")
	}
}