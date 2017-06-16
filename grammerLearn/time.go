package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Print(time.Now())
	var s string
	s = time.Now().String()
	fmt.Println(s)
}