package main

import (
	"fmt"
	"reflect"
)

//interface 物件導向 -> 共同的部份
//A variable of interface type can store a value of any type with a method set that is any superset of the interface.
//"interface{}" 很像swift type any
func anytype(any interface{}){
	fmt.Println(any,"'s type is :",reflect.TypeOf(any))
}
//example 1
type Cat  struct{
	name string
}
type Dog struct{
	name string
}
//make common function as interface ,call same func  do (eat)  different things
type Action interface{
	eat()
}
func (c Cat) eat(){
	fmt.Println(c.name," eat fish. ")
	fmt.Println(reflect.TypeOf(c))//main.Cat, basicType
}
func (d *Dog) eat(){
	fmt.Println(d.name,"eat meat.")
	fmt.Println(reflect.TypeOf(d))//*main.Dog, pointerType
}
func act(a Action){
	a.eat()
}

//example2
type Person struct {
	fname string
	lname string
}
type Agent struct{
	Person
	license bool
}
type human interface{
	speak()
}
func (p Person) speak(){
	fmt.Println("person speaking!")
}
func (s Agent) speak(){
	fmt.Println("Agent speaking")
}
func whoIsSpeaking(h human){
	h.speak()
}

func main(){

	anytype("string")
	anytype(1)
	anytype(struct{}{})
	anytype(struct{name string}{"Eve"})

	//物件導向 -> 共同的部份
	fmt.Println("-- interface 物件導向...共同的部份 --")
	c := Cat{"Kitty"}
	d := Dog{"Snoopy"}
	act(c)
	act(&d)

	//example2
	p := Person{
		fname: "Eve",
		lname: "Tung",
	}
	a := Agent{
		p,
		true,
	}
	whoIsSpeaking(p)
	whoIsSpeaking(a)

}



