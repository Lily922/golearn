package main

import (
	"fmt"
	"os"
	"strings"
	"io"
)

// 本函数主要练习输入输出
// https://www.jianshu.com/p/abc396787a32

/*
type Reader interface {
    Read(p []byte) (n int, err error)
}
 */
func readTest(i int){
	var b []byte
	var err error
	switch i{
	case 1:
		b,err = readFrom(strings.NewReader("huhuhu"),3)
	case 2:
		b,err = readFrom(os.Stdin,3)
	case 3:
		f,err:=os.Open("/Users/lily/golang/src/github.com/LilyFaFa/golearn/basicgrammar/course2/test")
		if err !=nil {
			panic(err)
		}
		b,err=readFrom(f,3)
	}
	if err == nil{
		fmt.Println(string(b))
	}
}

func writeTest(w io.Writer,b []byte ,i int){
	switch i {
	case 1:
		err:=writeTo(w,b)
		if err !=nil{
			fmt.Println(err)
		}
	}
}

func readFrom(r io.Reader,num int)([]byte,error){
	p:=make([]byte,num)
	n,err :=r.Read(p)
	if n>0 {
		return p[0:n],err
	}
	return p,err

}

func writeTo(w io.Writer,b []byte)error{
	_,err:=w.Write(b)
	fmt.Println(b)
	return err
}

func main(){
	readTest(1)
	readTest(2)
	readTest(3)
	writer:=os.Stdout
	writeTest(writer,[]byte{'d','d'},1)
}