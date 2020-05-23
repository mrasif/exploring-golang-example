package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("from a person - this is my first name ",p.first)
}

func main() {
	p1:= person{
		first: "Asif",
	}
	fmt.Printf("%T\n",p1)

	var x human
	x=p1
	fmt.Printf("%T\n",x)
	
}