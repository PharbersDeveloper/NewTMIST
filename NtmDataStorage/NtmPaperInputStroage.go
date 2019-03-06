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

// NtmPaperInputStorage stores all of the tasty modelleaf, needs to be injected into
// PaperInput and PaperInput Resource. In the real world, you would use a database for that.
type NtmPaperInputStorage struct {
	images  map[string]*NtmModel.PaperInput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmPaperInputStorage) NewPaperInputStorage(args []BmDaemons.BmDaemon) *NtmPaperInputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmPaperInputStorage{make(map[string]*NtmModel.PaperInput), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmPaperInputStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.PaperInput {
	in := NtmModel.PaperInput{}
	var out []NtmModel.PaperInput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.PaperInput
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
func (s NtmPaperInputStorage) GetOne(id string) (NtmModel.PaperInput, error) {
	in := NtmModel.PaperInput{ID: id}
	out := NtmModel.PaperInput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("PaperInput for id %s not found", id)
	return NtmModel.PaperInput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmPaperInputStorage) Insert(c NtmModel.PaperInput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmPaperInputStorage) Delete(id string) error {
	in := NtmModel.PaperInput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("PaperInput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmPaperInputStorage) Update(c NtmModel.PaperInput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("PaperInput with id does not exist")
	}

	return nil
}

func (s *NtmPaperInputStorage) Count(req api2go.Request, c NtmModel.PaperInput) int {
	r, _ := s.db.Count(req, &c)
	return r
}