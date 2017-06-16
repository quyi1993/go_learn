package main

import "fmt"

/*
This function intSeq returns another function,
which we define anonymously in the body of intSeq.
The returned function closes over the variable i to form a closure.
 */
func intSeq() (func() int,string) {
	i := 0
	return func() int {
		i += 1
		return i
	},"result"
}

func main() {
	nextInt, str := intSeq()
	fmt.Println(nextInt(), str)
	fmt.Println(nextInt(), str)
	fmt.Println(nextInt(), str)

	newNextInt, newStr := intSeq()
	fmt.Println(newNextInt(), newStr)



}