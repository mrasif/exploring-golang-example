package main

import "fmt"

type person struct {
	first string
}

type mongo map[int]person
type postgres map[int]person

func (m mongo) save(n int, p person){
	m[n]=p
}

func (m mongo) retrive(n int) person {
	return m[n]
}

func (pg postgres) save(n int, p person){
	pg[n]=p
}

func (pg postgres) retrive(n int) person{
	return pg[n]
}

type accessor interface {
	save(n int, p person)
	retrive(n int) person
}

func put(a accessor, n int, p person){
	a.save(n,p)
}

func get(a accessor, n int) person {
	return a.retrive(n)
}

func main() {

	dbm:=mongo{}
	dbp:=postgres{}

	p1 := person{
		first: "Asif",
	}

	p2 := person{
		first: "Sam",
	}
	
	put(dbm,1,p1)
	put(dbm,2,p2)
	fmt.Println(get(dbm,1))
	fmt.Println(get(dbm,2))

	put(dbp,1,p1)
	put(dbp,2,p2)
	fmt.Println(get(dbp,1))
	fmt.Println(get(dbp,2))
	
}