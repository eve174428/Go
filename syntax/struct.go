package main

import (
	"fmt"
	"reflect"
)

//struct
type Student struct{
	no int
	name string
}

func (s1 Student) getName() string{ //return string type value
	return s1.name
}

func (s Student) speak(){
	fmt.Println("Student no.",s.no,s.name, `says, "Goood morning, Eve"`)
}

type Person struct{
	fname string
	lname string
}

type SecretAgent struct{
	Person
	license bool
}

func (sa SecretAgent) speak(){
	fmt.Println(sa.fname,sa.lname,`says, "i wanna be a Go Dev. in 2022."`)
}
func (s Person) speak(){
	fmt.Printf("Hi, my name is %s %s.", s.fname,s.lname)
}

func main(){
	//struct
	s1 := Student{1, "Chi"}
	s2 := Student{
		no: 2,
		name: "Eve",
	}
	fmt.Println(struct{class string}{"computer science"})
	fmt.Println(s1,s2)
	fmt.Println(s1.getName())
	fmt.Println(reflect.TypeOf(s1))
	s1.speak()//Student no. 1 Chi says, "Goood morning, Eve"

	sal := SecretAgent{
		Person{
			"Eve",
			"Tung",
		},
		true,
	}

	sal.speak()
	sal.Person.speak()
}
