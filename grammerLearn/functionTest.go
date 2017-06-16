package main

import (
	"fmt"
	"math"
)

func main() {
	num := max(4,9)
	fmt.Printf("%d" ,num)
	s1,s2 := swap("abc","xyz")
	fmt.Printf("%s %s",s1,s2)



	num1:=1
	num2:=6
	swap_int(&num1,&num2)
	fmt.Printf("%d %d",num1, num2)


	getSquareRoot := func(x float64) float64{
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func swap(str1,str2 string) (string, string) {
	return str2, str1
}

//引用传递值
func swap_int(a, b *int)  {
	temp:=*a
	*a = *b
	*b = temp
}

