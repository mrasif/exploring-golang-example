package main

import (
	"fmt"

	architecture "github.com/mrasif/exploring-golang"
	"github.com/mrasif/exploring-golang/models"
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

	ps:=architecture.NewPersonService(dbm)

	ps.Put(1,p1)
	ps.Put(2,p2)
	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(2))

	architecture.Put(dbp,1,p1)
	architecture.Put(dbp,2,p2)
	fmt.Println(architecture.Get(dbp,1))
	fmt.Println(architecture.Get(dbp,2))
	
}