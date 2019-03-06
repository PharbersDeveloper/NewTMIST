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

// NtmRepresentativeInputStorage stores all of the tasty modelleaf, needs to be injected into
// RepresentativeInput and RepresentativeInput Resource. In the real world, you would use a database for that.
type NtmRepresentativeInputStorage struct {
	images  map[string]*NtmModel.RepresentativeInput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmRepresentativeInputStorage) NewRepresentativeInputStorage(args []BmDaemons.BmDaemon) *NtmRepresentativeInputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmRepresentativeInputStorage{make(map[string]*NtmModel.RepresentativeInput), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmRepresentativeInputStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.RepresentativeInput {
	in := NtmModel.RepresentativeInput{}
	var out []NtmModel.RepresentativeInput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.RepresentativeInput
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
func (s NtmRepresentativeInputStorage) GetOne(id string) (NtmModel.RepresentativeInput, error) {
	in := NtmModel.RepresentativeInput{ID: id}
	out := NtmModel.RepresentativeInput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("RepresentativeInput for id %s not found", id)
	return NtmModel.RepresentativeInput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmRepresentativeInputStorage) Insert(c NtmModel.RepresentativeInput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmRepresentativeInputStorage) Delete(id string) error {
	in := NtmModel.RepresentativeInput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("RepresentativeInput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmRepresentativeInputStorage) Update(c NtmModel.RepresentativeInput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("RepresentativeInput with id does not exist")
	}

	return nil
}

func (s *NtmRepresentativeInputStorage) Count(req api2go.Request, c NtmModel.RepresentativeInput) int {
	r, _ := s.db.Count(req, &c)
	return r
}