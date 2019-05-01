package NtmDataStorage

import (
	"fmt"
	"errors"
	"Ntm/NtmModel"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
)

// NtmManagerinputStorage stores all of the tasty modelleaf, needs to be injected into
// Managerinput and Managerinput Resource. In the real world, you would use a database for that.
type NtmManagerinputStorage struct {
	images  map[string]*NtmModel.Managerinput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmManagerinputStorage) NewManagerinputStorage(args []BmDaemons.BmDaemon) *NtmManagerinputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmManagerinputStorage{make(map[string]*NtmModel.Managerinput), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmManagerinputStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.Managerinput {
	in := NtmModel.Managerinput{}
	var out []NtmModel.Managerinput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.Managerinput
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
func (s NtmManagerinputStorage) GetOne(id string) (NtmModel.Managerinput, error) {
	in := NtmModel.Managerinput{ID: id}
	out := NtmModel.Managerinput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Managerinput for id %s not found", id)
	return NtmModel.Managerinput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmManagerinputStorage) Insert(c NtmModel.Managerinput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmManagerinputStorage) Delete(id string) error {
	in := NtmModel.Managerinput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Managerinput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmManagerinputStorage) Update(c NtmModel.Managerinput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Managerinput with id does not exist")
	}

	return nil
}

func (s *NtmManagerinputStorage) Count(req api2go.Request, c NtmModel.Managerinput) int {
	r, _ := s.db.Count(req, &c)
	return r
}