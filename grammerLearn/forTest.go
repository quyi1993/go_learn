package main

import "fmt"

func main() {
	var b int = 15
	var a int

	numbers := [6]int{1,2,3,5}

	//打印1-10
	for i:=0; i<10; i++ {
		fmt.Printf("%d,",i)
	}

	//打印1-b
	for a<b {
		fmt.Printf("%d,",a)
		a ++
	}

	//打印数组numbers
	for i,x:= range numbers {
		fmt.Printf("d第%d位x的值%d\n",i,x)
	}
}
