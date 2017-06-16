package main

func Factorial(x int) (result int) {
	if x==0 {
		return 1
	}
	result = x * Factorial(x-1)
	return result
}

func fibonaci(n int) int {
	if n==1 {
		return 1
	} else if n==2 {
		return 1
	}
	return fibonaci(n-1) * fibonaci(n-2)
}
