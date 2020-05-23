package main

import (
	"fmt"

	"github.com/mrasif/exploring-golang/models"
	"github.com/mrasif/exploring-golang/services"
	"github.com/mrasif/exploring-golang/storage/mongo"
	"github.com/mrasif/exploring-golang/storage/postgres"
)




func main() {

	dbm:=mongo.Db{}
	dbp:=postgres.Db{}

	p1 := models.Person{
		First: "Asif",
	}

	p2 := models.Person{
		First: "Sam",
	}

	s := services.PersonService{
		DB:dbm,
	}
	
	s.Put(1,p1)
	s.Put(2,p2)
	fmt.Println(s.Get(1))
	fmt.Println(s.Get(2))

	s.DB=dbp

	s.Put(1,p1)
	s.Put(2,p2)
	fmt.Println(s.Get(1))
	fmt.Println(s.Get(2))
	
}