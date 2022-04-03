package main

import (
	"archive/zip"
	"reflect"
	"fmt"
)

func main(){
	var x int
	xLoc := &x
	fmt.Println(*xLoc)//7
	z := *int
	xLoc = z
	fmt.Println(x)
}
