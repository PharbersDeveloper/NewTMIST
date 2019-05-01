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

// NtmImageStorage stores all of the tasty modelleaf, needs to be injected into
// Image and Image Resource. In the real world, you would use a database for that.
type NtmImageStorage struct {
	images  map[string]*NtmModel.Image
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmImageStorage) NewImageStorage(args []BmDaemons.BmDaemon) *NtmImageStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmImageStorage{make(map[string]*NtmModel.Image), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmImageStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.Image {
	in := NtmModel.Image{}
	var out []NtmModel.Image
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
func (s NtmImageStorage) GetOne(id string) (NtmModel.Image, error) {
	in := NtmModel.Image{ID: id}
	out := NtmModel.Image{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Image for id %s not found", id)
	return NtmModel.Image{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmImageStorage) Insert(c NtmModel.Image) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmImageStorage) Delete(id string) error {
	in := NtmModel.Image{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Image with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmImageStorage) Update(c NtmModel.Image) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Image with id does not exist")
	}

	return nil
}
