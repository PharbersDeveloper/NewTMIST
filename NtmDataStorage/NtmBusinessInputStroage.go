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

// NtmBusinessInputStorage stores all of the tasty modelleaf, needs to be injected into
// BusinessInput and BusinessInput Resource. In the real world, you would use a database for that.
type NtmBusinessInputStorage struct {
	images  map[string]*NtmModel.BusinessInput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmBusinessInputStorage) NewBusinessInputStorage(args []BmDaemons.BmDaemon) *NtmBusinessInputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmBusinessInputStorage{make(map[string]*NtmModel.BusinessInput), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmBusinessInputStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.BusinessInput {
	in := NtmModel.BusinessInput{}
	var out []NtmModel.BusinessInput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.BusinessInput
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
func (s NtmBusinessInputStorage) GetOne(id string) (NtmModel.BusinessInput, error) {
	in := NtmModel.BusinessInput{ID: id}
	out := NtmModel.BusinessInput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("BusinessInput for id %s not found", id)
	return NtmModel.BusinessInput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmBusinessInputStorage) Insert(c NtmModel.BusinessInput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmBusinessInputStorage) Delete(id string) error {
	in := NtmModel.BusinessInput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("BusinessInput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmBusinessInputStorage) Update(c NtmModel.BusinessInput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("BusinessInput with id does not exist")
	}

	return nil
}

func (s *NtmBusinessInputStorage) Count(req api2go.Request, c NtmModel.BusinessInput) int {
	r, _ := s.db.Count(req, &c)
	return r
}