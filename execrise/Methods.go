package main

import "fmt"

type rect struct {
	width,height int
}

func (r rect) getArea() (area int){
	return r.width * r.height
}

func (r *rect) getPerim() (perim int){
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{8,9}
	fmt.Println("area: ", r.getArea())
	fmt.Println("perim:", r.getPerim())


}