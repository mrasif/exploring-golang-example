package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("from a person - this is my first name ",p.first)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I am secret agent - this is my name ",sa.first)
}

func main() {
	p1:= person{
		first: "Asif",
	}
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	var x,y human
	x=p1
	y=sa1
	x.speak()
	y.speak()
	
}