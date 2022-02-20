package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main(){
	realine := os.Args[1]
	fmt.Println(os.Args[0])//the program you've executed,starting with the program name
	fmt.Println(os.Args[1])
	html := fmt.Sprint(`
	<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8">
	<title>03_os-Args</title>
	</head>
	<body>
	<h1>`+ realine +`</h1>
	</body>
	</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil{
		log.Fatal("error create file : ",err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(html))

	//$go run 03_os-Args.go test
}
