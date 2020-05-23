package services

import (
	"github.com/mrasif/exploring-golang/models"
)


// Accessor for data operation
type Accessor interface {
	Save(n int, p models.Person)
	Retrive(n int) models.Person
}

// PersonService contains db
type PersonService struct {
	DB Accessor
}

// Put for save
func (s PersonService) Put(n int, p models.Person){
	s.DB.Save(n,p)
}

// Get for retrive
func (s PersonService) Get(n int) models.Person {
	return s.DB.Retrive(n)
}