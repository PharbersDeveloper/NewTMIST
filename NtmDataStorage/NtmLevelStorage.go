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

// NtmLevelStorage stores all of the tasty modelleaf, needs to be injected into
// Level and Level Resource. In the real world, you would use a database for that.
type NtmLevelStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmLevelStorage) NewLevelStorage(args []BmDaemons.BmDaemon) *NtmLevelStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmLevelStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmLevelStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.Level {
	in := NtmModel.Level{}
	var out []NtmModel.Level
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
func (s NtmLevelStorage) GetOne(id string) (NtmModel.Level, error) {
	in := NtmModel.Level{ID: id}
	out := NtmModel.Level{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Level for id %s not found", id)
	return NtmModel.Level{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmLevelStorage) Insert(c NtmModel.Level) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmLevelStorage) Delete(id string) error {
	in := NtmModel.Level{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Level with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmLevelStorage) Update(c NtmModel.Level) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Level with id does not exist")
	}

	return nil
}
