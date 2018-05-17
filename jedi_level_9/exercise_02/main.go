package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("My name is", p.name, "and I am", p.age, "years old.")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "Tony",
		age:  25,
	}
	//saySomething(p)
	saySomething(&p)
}
