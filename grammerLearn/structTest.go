package main

import "fmt"

type person struct{
	name string
	age int
}

func main() {
	fmt.Println(person{"quyi",18})
	fmt.Println(&person{"mary",20})

	s := person{"bob",15}
	fmt.Println(s.name, s.age)
}