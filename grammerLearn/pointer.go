package main

import "fmt"

func main() {
	var ptr *int
	fmt.Printf("ptr的值为：%x\n", ptr)
	if(ptr!=nil) {
		fmt.Print("ptr不是空指针")
	}


	var a int =20
	var ip *int
	ip = &a
	fmt.Printf("a 变量的地址是：%x\n", &a)

	fmt.Printf("ip 变量存储的指针地址：%x\n", ip)

	fmt.Printf("*ip变量的值：%d\n", *ip)

}