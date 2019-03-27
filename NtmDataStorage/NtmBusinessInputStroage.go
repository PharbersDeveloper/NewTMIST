package NtmDataStorage

import (
	"fmt"
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
)

// NtmBusinessinputStorage stores all of the tasty modelleaf, needs to be injected into
// Businessinput and Businessinput Resource. In the real world, you would use a database for that.
type NtmBusinessinputStorage struct {
	images  map[string]*NtmModel.Businessinput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmBusinessinputStorage) NewBusinessinputStorage(args []BmDaemons.BmDaemon) *NtmBusinessinputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmBusinessinputStorage{make(map[string]*NtmModel.Businessinput), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmBusinessinputStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.Businessinput {
	in := NtmModel.Businessinput{}
	var out []NtmModel.Businessinput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.Businessinput
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s NtmBusinessinputStorage) GetOne(id string) (NtmModel.Businessinput, error) {
	in := NtmModel.Businessinput{ID: id}
	out := NtmModel.Businessinput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Businessinput for id %s not found", id)
	return NtmModel.Businessinput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmBusinessinputStorage) Insert(c NtmModel.Businessinput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmBusinessinputStorage) Delete(id string) error {
	in := NtmModel.Businessinput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Businessinput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmBusinessinputStorage) Update(c NtmModel.Businessinput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Businessinput with id does not exist")
	}

	return nil
}

func (s *NtmBusinessinputStorage) Count(req api2go.Request, c NtmModel.Businessinput) int {
	r, _ := s.db.Count(req, &c)
	return r
}