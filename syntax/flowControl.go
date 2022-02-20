package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main(){

		//control structures
		fmt.Println("--- if ...else ---")
		var isTrue *bool
		t := true
		isTrue = &t

		if isTrue == nil{
			fmt.Println("no value")
		}else if !*isTrue{
			fmt.Println("False")
		}else{
			fmt.Println("True")
		}

		fmt.Println("===== for loops ====")
		for i := 0 ; i < 10; i++ {
			fmt.Println(i)//1~9
		}

		fmt.Println("-- for range --") //swift- for...in ;flutter- forEach
		fluits := []string{"apple","banana","orange"}
		for i, v := range fluits{
			fmt.Println(i,".",v)
		}

		//swift dictionary; flutter maps
		dic := make(map[string]interface{})
		dic["台灣"] = "Taiwan"
		dic["美國"] = "USA"

		for k, v := range dic{
			fmt.Printf("key: %s and value: %s \n", k ,v )
		}

		fmt.Println("--- for continue --")
		var even []int
		//odd
		for i := 0; i < 10; i++{
			// even
			if i%2 == 0{
				even = append(even, i)
				continue
			}else if i == 9{
				break
			}
			fmt.Println(i)//1,3,5,7,9
		}
		fmt.Println(even)//0~4

		fmt.Println("--- switch ---")

		weekNo := []int{1,2,3,4,5,6,7}
		for _, v := range weekNo{ //_ , same as swift "_" omit this var
			switch v {
			case 1:
				fmt.Println("Mon.")
			case 2:
				fmt.Println("Tue.")
			case 3:
				fmt.Println("Wed.")
			case 4:
				fmt.Println("Thur.")
			case 5:
				fmt.Println("Fri.")
			case 6:
				fmt.Println("Sat")
			default:
				break
			}
		}

		err := errors.New("404")
		fmt.Println(err)
		fmt.Println(reflect.TypeOf(err))
		switch err.Error(){
		case "404":
			fmt.Printf("%s Not Found.\n", err.Error())
		default:
			break
		}

}
