package main

import "fmt"

type DivideError struct {
	dividee int
	divider int
}

func (de DivideError) error() string{
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
`
	return strFormat
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider==0 {
		dData := DivideError{varDividee,varDivider}
		errorMsg = dData.error()
		return
	}else {
		result = varDividee / varDivider
		return
	}
}

func main() {
	if result, errorMsg:=Divide(100,10) ; errorMsg=="" {
		fmt.Println("100/10=", result)
	}
	if _, errorMsg:=Divide(100,0) ; errorMsg!="" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}