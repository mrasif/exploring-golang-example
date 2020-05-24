package architecture


import (
	"testing"

	architecture "github.com/mrasif/exploring-golang"
	"github.com/mrasif/exploring-golang/models"
	"github.com/mrasif/exploring-golang/storage/mongo"
)



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

func TestPut(t *testing.T){
	dbm := mongo.Db{}

	p := models.Person{
		First: "Asif",
	}

	architecture.Put(dbm,1,p)

	got:=architecture.Put(dbm,1)
	if got != p {
		t.Fatalf("Want %v, Got %v",p,got)
	}
}