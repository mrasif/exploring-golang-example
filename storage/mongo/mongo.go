package mongo

import "github.com/mrasif/exploring-golang/models"

// Db for storage
type Db map[int]models.Person

// Save will store
func (m Db) Save(n int, p models.Person){
	m[n]=p
}

// Retrive will fetch
func (m Db) Retrive(n int) models.Person {
	return m[n]
}

// Accessor for data operation
type Accessor interface {
	Save(n int, p models.Person)
	Retrive(n int) models.Person
}