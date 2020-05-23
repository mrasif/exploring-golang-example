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

type personService struct {
	db accessor
}

func (s personService) put(n int, p person){
	s.db.save(n,p)
}

func (s personService) get(n int) person {
	return s.db.retrive(n)
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

	s := personService{
		db:dbm,
	}
	
	s.put(1,p1)
	s.put(2,p2)
	fmt.Println(s.get(1))
	fmt.Println(s.get(2))

	s.db=dbp

	s.put(1,p1)
	s.put(2,p2)
	fmt.Println(s.get(1))
	fmt.Println(s.get(2))
	
}