package main

import "github.com/golang/go/src/fmt"

func unumber(){
	const(
		java =0
		python=1
		clang=2
		golang=3
	)
	const (
		a int =12
		b int =13
		c int =14
	)
	const(
		f = 1<<(10*iota)
		h
		j
		k
	)
	fmt.Println(a,b,c)
	fmt.Println(java,python,clang,golang)
	fmt.Println(f,h,j,k)
}
func main(){
	unumber()
}