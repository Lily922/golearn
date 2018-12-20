package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(i int)string{
	var s string
	for ;i>0; i=i/2 {
		s=strconv.Itoa(i%2)+s
	}
	return s
}

func sum(){
	add:=0
	for i:=0;i<10;i++{
		add+=i
	}
	fmt.Println(add)
}

func printFile(file string){
    f,err:=os.Open(file)
	if err !=nil {
		panic(err)
	}
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main(){
	result:=convertToBin(1)
	sum()
	printFile("/Users/lily/golang/src/github.com/LilyFaFa/golearn/basicgrammar/course5/test")
	fmt.Println(result)
}