package main

import "fmt"

func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean的值为：%f\n", mean)

	str := "123"
	fmt.Printf(str)
	var num int
	num = 0
	for index,c:= range str {
		n := int(c)-'0'
		fmt.Println(index, n)
		num = num*10 + n
	}
	fmt.Println(num)

}
