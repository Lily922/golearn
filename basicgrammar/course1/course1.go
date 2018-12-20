package  main

import (
	"github.com/golang/go/src/fmt"
)

var (
	aa = 90
	bb = 80
	ss = "hello"
)
func variableZeroValue(){
	var a ,b int
	var s string
	fmt.Printf("%d,%d,%s \n",a,b,s)
}

func variableInitialValue(){
	var a,b int =3,4
	var s string ="hello"
	fmt.Println(a,b,s)
}

func variableDeductionValue(){
	a,b:=1,2
	s := "hello"
	fmt.Println(a,b,s)
}

func main(){
	variableZeroValue()
	variableInitialValue()
	variableDeductionValue()
	fmt.Println("Hello World")
	fmt.Println(aa,bb,ss)
}