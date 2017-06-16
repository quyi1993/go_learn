package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	if a < 20 && a > 16{
		fmt.Printf("16<a<20\n")
	}else {
		fmt.Printf("a<=16或者a>=20\n")
	}
	fmt.Printf("a 的值为：%d \n", a)


	if (a==100) {
		if b==200 {
			fmt.Printf("a 的值为 100且b的值为 200")
		}
	}

}
