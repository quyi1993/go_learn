package main

import "fmt"

func main() {
	var x interface{}
	x= true
	switch i := x.(type) {
	case nil: fmt.Printf("x 的类型：%T", i)
	case int: fmt.Printf("x is int")
	case string:fmt.Printf("x is string")
	case bool:fmt.Printf("x is bool")
	case float64:fmt.Printf("x is float64")
	default:
		fmt.Printf("unknown")
	}
}
