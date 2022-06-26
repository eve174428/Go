package main

import "fmt"

func main() {
	var m = map[string]string{}
	m["Name"] = "eve"
	fmt.Println(m["Name"])//eve
	var mm = make(map[string]string)
	mm["Last"] = "tong"
	fmt.Println(mm["Last"])

	name, isExist := m["Name"]
	fmt.Println("Is name Exist ?", isExist,",",name)
	last, isExist := m["last"]
	fmt.Println("Is last name Exist ?", isExist,last)
	if lastName, isExist := mm["Last"]; isExist {
		fmt.Println(m["Name"],"", lastName)
	}//; then()
}
