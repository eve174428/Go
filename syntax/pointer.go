package main

import (
	"fmt"
	"reflect"
)

func main(){

	//Pointer type (which include memory address"&", and value "*", it depens on memory address.)
	m1 := 2 //BaseType
	ptr := &m1 //& memory place(編號)
	fmt.Println(ptr)
	fmt.Println(reflect.TypeOf(ptr))//*int , PointerType (var m1 *int)
	fmt.Println(*ptr)//2, * get value
	fmt.Println(reflect.TypeOf(*ptr))//int

	criminal_1,crimianl_2 := "Ben","John"
	fmt.Println(criminal_1,crimianl_2)
	swap(&criminal_1,&crimianl_2)//交換"&"(編號)
	fmt.Println(criminal_1,crimianl_2)

	fmt.Println("-- with struct --")
	type Food struct{
		fluit string
	}

	var f *Food
	f = &Food{"Banana"}
	fmt.Println(f.fluit)

}

//pointer
func swap(p1, p2 *string){ //交換值
	tmp := *p2
	*p2 = *p1
	*p1 = tmp
}

