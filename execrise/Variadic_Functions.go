package main

import "fmt"

//Here’s a function that will take an arbitrary number of ints as arguments.

func sum(nums ...int) int {
	var sum int
	for _,num:= range nums {
		sum += num
	}
	fmt.Println(sum)
	return sum
}

func main(){
	sum(1,2,3)

	nums := []int{1,2,3,4}
	sum(nums...)
}