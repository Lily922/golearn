package main

import (
	"github.com/golang/go/src/fmt"
	"github.com/golang/go/src/io/ioutil"
)

func readFile(fileName string){
	b,err:=ioutil.ReadFile(fileName)
	if err ==nil{
		fmt.Printf("%s\n",b)
	}else{
		fmt.Println(err)
	}
}

func eval(s string){
	switch s{
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	default:
		fmt.Println("lalala")
	}
}

func grade (g int){
	switch{
	case g>90:
		fmt.Println("prefect")
	case g>60:
		fmt.Println("ok")
	case g>0:
		fmt.Println("angry")
	default:
		panic(fmt.Sprintf("%s","error happend"))
	}
}

func testSwitch(g int){
	switch{
	case g>100:
		panic(fmt.Sprintf("grade is %d",g))
	case g>90:
		fmt.Println("perfect")
	case g>60:
		fmt.Println("ok")
	case g>0:
		fmt.Println("angry")
	default:
		panic(fmt.Sprintf("grade is %d",g))
	}
}
func main(){
	readFile("test")
	eval("a")
	eval("test")
	grade(80)
	//grade(-20)
	testSwitch(-10)
	testSwitch(80)
}