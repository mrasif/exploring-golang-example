package architecture

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
	db Accessor
}

func NewPersonService(db Accessor) PersonService{
	return PersonService{
		db:db,
	}
}

// Put for save
func (s PersonService) Put(n int, p models.Person){
	s.db.Save(n,p)
}

// Get for retrive
func (s PersonService) Get(n int) models.Person {
	return s.db.Retrive(n)
}

// Put for save
func Put(a Accessor, n int, p models.Person){
	a.Save(n,p)
}

// Get for retrive
func Get(a Accessor,n int) models.Person {
	return a.Retrive(n)
}