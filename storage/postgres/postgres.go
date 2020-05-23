package postgres

import "github.com/mrasif/exploring-golang/models"

// Db is a storage
type Db map[int]models.Person

// Save will store
func (pg Db) Save(n int, p models.Person){
	pg[n]=p
}

// Retrive will fetch
func (pg Db) Retrive(n int) models.Person{
	return pg[n]
}

// Accessor for data operation
type Accessor interface {
	Save(n int, p models.Person)
	Retrive(n int) models.Person
}