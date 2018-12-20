package main

import "github.com/golang/go/src/fmt"

func arrTest(){
	arr :=[5]int{}
	arr[4]=1
	fmt.Println(arr)
	arr2 :=[...]int{1,2,3,4,5}
	fmt.Println(arr2)
	modify(arr)
	for i,j:=range arr{
		fmt.Printf("%d,%d,",i,j)
	}
}

func sliceTest(){
	s:=[]int{}
	s=append(s,1)
	s=append(s,2)
	fmt.Println(s[0:2])
	arr:=[...]int{0,1,2,3,4,5,6,7,8,9}
	// arr是s1的底层数组
	s1:=arr[2:6]
	//slice可以向后扩展，不可以向前扩展
	//s1[i]不可以超越len(s1)，向后扩展不可以超过底层数组cap(s1)
	// arr是s2的底层数组
	s2:=s1[3:5]
	//8,4 (2,3,4,5,6,7,8,9)
	fmt.Println(cap(s1),len(s1))
	//5,2 (5,6,7,8,9)
	fmt.Println(cap(s2),len(s2))
	s2=append(s2,3)
	//5,3 (5,6,3,8,9)
	fmt.Println(cap(s2),len(s2))
	s3:=[]int{1,2,3}
	fmt.Println(len(s3),cap(s3))
	//添加的元素超过cap,那么就会分配一个新的底层数组
}

func sliceTest2(){
	var s1 []int
	for i:=0;i<100;i++{
		s1=append(s1,i*2+1)
		//cap 会每次*2倍的扩容
		fmt.Printf("len:%d,cap:%d\n",len(s1),cap(s1))
	}
	//指向同一个底层数组
	s2:=s1
	s2[1]=9090
	fmt.Println(s1[1])
	//分配一个新的底层数组
	s3 :=make([]int,100)
	copy(s3,s2)
	s3[1]=9
	fmt.Println(s1[1])

}

func mapTest(){
	// map是无序的
	m:=map[string]string{
		"a":"b",
		"c":"d",
	}
	m["hello"]="world"
	fmt.Println(m)
	//empty map
	m2:=make(map[string]int)
	//nil
	var m3 map[int]int

	m3=make(map[int]int)
	m3[1]=2

	// 查找
	size1,ok:=m3[1]
	fmt.Println(size1,ok)
	fmt.Println(m2,m3)
	//删除
	delete(m,"c")
	fmt.Println(m)

	for key,value:=range m{
		fmt.Println(key,value)
	}
}

func stringTest(){

}
func modify(inputs [5]int){
	inputs[0]=90
}

func main(){
	//arrTest()
	//sliceTest()
	//sliceTest2()
	mapTest()
}