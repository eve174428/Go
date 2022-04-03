package main

import (
	"fmt"
	"strings"
)

func main(){
	variables()
	basicTypes()
}

func variables(){

	//var
	fmt.Println("--- var---")
	var age int
	age = 33

	//var mutiple variables
	var n1, n2 int =  1, 2
	var (
		year = 77
		bc = 1911
	)
	//Variables with initializers :
	month := 2
	month = n1 + n2 // month += 1
	n1,n2,n3 := 5,5,4
	var date int = n1 + n2 + n3
	fmt.Println("i am ", age , " years old. bron in ", year + bc, "/", month, "/" ,date)
}

func basicTypes(){
	//Types
	fmt.Println("========== Type ==========")
	//int32 vs int64
	fmt.Println("--- int32 vs int64 ---")
	var n1 int32
	var n2 int64
	fmt.Println(int64(n1) + n2) //0 , because default value is 0

	//string
	fmt.Println("--- String ---")
	var name string
	name = "Eve"
	country := "Taiwan"
	fmt.Println(name, "is come from", country)
	fmt.Println(strings.Contains(name, country))//false
	fmt.Println(strings.ReplaceAll(name, "e", "a"))//Eva

	//to array
	fmt.Println(strings.Split(name,""))//[E v e]
	fmt.Println(strings.Split(name," "))//[Eve]
	arrSplited := strings.Split(name,"")
	arrSplited1 := strings.Split(name," ")
	fmt.Println(arrSplited[0])
	fmt.Println(arrSplited1[0])

}

