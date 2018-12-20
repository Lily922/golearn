package main

import (
	"fmt"

	"github.com/LilyFaFa/golearn/interfacegrammar/person"
)

func main() {
	persons := []person.Person{}
	persons = append(persons,
		&person.Student{
			Name: "Tom",
			Age:  12,
			Job:  "Student",
		},
		&person.Teacher{
			Name: "Sam",
			Age:  34,
			Job:  "Teacher",
		})

	for _, p := range persons {
		Inspect(p)
		if CheckTeacher(p) {
			fmt.Printf("check teacher, %s is a teacher \n", p.GetName())
		}
		if FindTom(p) {
			fmt.Printf("find Tom , he is a %s \n", p.GetJob())
		}
	}
}

//FindTom check whose name is Tom
func FindTom(p person.Person) bool {
	fmt.Printf("%T,%v \n", p, p)
	if p.GetName() == "Tom" {
		return true
	}
	return false
}

func Inspect(p person.Person) {

	switch p.(type) {
	case *person.Teacher:
		fmt.Println("teacher")
	case *person.Student:
		fmt.Println("student")
	}
}

func CheckTeacher(p person.Person) bool {
	if _, ok := p.(*person.Teacher); ok {
		return true
	}
	return false
}
