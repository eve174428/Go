package main

import (
	"fmt"
)

func main(){
	c := make(chan string,2)//wont block until the channel is full
	//sending "hello" across channel
	c <- "hello"
	c <- "World"
	/*sender will block until something is ready to receive.*/
	//recive value from channel "c <-"hello"
	msg := <- c
	fmt.Println(msg)
	//recive value from channel "c <-"World"
	msg = <- c
	fmt.Println(msg)
}
