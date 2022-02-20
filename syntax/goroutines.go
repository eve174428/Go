package main

//goroutine :  flutter async /  swift dispatchQueue.main.async
import (
	"fmt"
	"reflect"
	"sync"
	"time"
)
func step1(){
	for i := 0 ; i < 1; i ++{
		fmt.Println("1")
	}
}
func step2(){
	for i := 0 ; i < 3; i ++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("2 - %d \n",i+1)
	}
}

func main(){
	//waitgroup sc
	await := &sync.WaitGroup{}
	fmt.Println(reflect.TypeOf(await))

	fmt.Println("Steps")
		//add -> done ->  wait
	await.Add(2)
	go func(){
		step1()
		await.Done()
	}()

	go func ()  {
		step2()
		await.Done()
	}()

	await.Wait()//Wait blocks until the WaitGroup counter is zero.
	fmt.Println("Done!")
}
