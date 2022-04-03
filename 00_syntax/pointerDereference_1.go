package main

import (
	"fmt"
)

func main(){
	x := 7
	y := &x //& obtain from (a pointer) the address of a data item held in another location.
	fmt.Println(x,y)
	//get x value
	fmt.Println(*y)// x = 7 = *y
	//change x value
	*y = 8
	fmt.Println(x,y)
}
