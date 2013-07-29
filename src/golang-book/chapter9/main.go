package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) sayHello() string {
	return "Hello " + p.name
}

func main(){
	var p Person
	p.name = "jon"
	p.age = 27
	fmt.Println(p)	
	fmt.Println(p.sayHello())
}
