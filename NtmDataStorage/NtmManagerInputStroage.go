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

// NtmManagerInputStorage stores all of the tasty modelleaf, needs to be injected into
// ManagerInput and ManagerInput Resource. In the real world, you would use a database for that.
type NtmManagerInputStorage struct {
	images  map[string]*NtmModel.ManagerInput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmManagerInputStorage) NewManagerInputStorage(args []BmDaemons.BmDaemon) *NtmManagerInputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmManagerInputStorage{make(map[string]*NtmModel.ManagerInput), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmManagerInputStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.ManagerInput {
	in := NtmModel.ManagerInput{}
	var out []NtmModel.ManagerInput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.ManagerInput
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
func (s NtmManagerInputStorage) GetOne(id string) (NtmModel.ManagerInput, error) {
	in := NtmModel.ManagerInput{ID: id}
	out := NtmModel.ManagerInput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ManagerInput for id %s not found", id)
	return NtmModel.ManagerInput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmManagerInputStorage) Insert(c NtmModel.ManagerInput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmManagerInputStorage) Delete(id string) error {
	in := NtmModel.ManagerInput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ManagerInput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmManagerInputStorage) Update(c NtmModel.ManagerInput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ManagerInput with id does not exist")
	}

	return nil
}

func (s *NtmManagerInputStorage) Count(req api2go.Request, c NtmModel.ManagerInput) int {
	r, _ := s.db.Count(req, &c)
	return r
}