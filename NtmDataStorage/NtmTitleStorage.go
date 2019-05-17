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

// NtmTitleStorage stores all of the tasty modelleaf, needs to be injected into
// Title and Title Resource. In the real world, you would use a database for that.
type NtmTitleStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmTitleStorage) NewTitleStorage(args []BmDaemons.BmDaemon) *NtmTitleStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmTitleStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmTitleStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.Title {
	in := NtmModel.Title{}
	var out []NtmModel.Title
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(&iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s NtmTitleStorage) GetOne(id string) (NtmModel.Title, error) {
	in := NtmModel.Title{ID: id}
	out := NtmModel.Title{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Title for id %s not found", id)
	return NtmModel.Title{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmTitleStorage) Insert(c NtmModel.Title) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmTitleStorage) Delete(id string) error {
	in := NtmModel.Title{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Title with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmTitleStorage) Update(c NtmModel.Title) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Title with id does not exist")
	}

	return nil
}
