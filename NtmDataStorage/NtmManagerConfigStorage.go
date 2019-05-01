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

// NtmManagerConfigStorage stores all of the tasty modelleaf, needs to be injected into
// ManagerConfig and ManagerConfig Resource. In the real world, you would use a database for that.
type NtmManagerConfigStorage struct {
	images  map[string]*NtmModel.ManagerConfig
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmManagerConfigStorage) NewManagerConfigStorage(args []BmDaemons.BmDaemon) *NtmManagerConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmManagerConfigStorage{make(map[string]*NtmModel.ManagerConfig), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmManagerConfigStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.ManagerConfig {
	in := NtmModel.ManagerConfig{}
	out := []NtmModel.ManagerConfig{}
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
func (s NtmManagerConfigStorage) GetOne(id string) (NtmModel.ManagerConfig, error) {
	in := NtmModel.ManagerConfig{ID: id}
	out := NtmModel.ManagerConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ManagerConfig for id %s not found", id)
	return NtmModel.ManagerConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmManagerConfigStorage) Insert(c NtmModel.ManagerConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmManagerConfigStorage) Delete(id string) error {
	in := NtmModel.ManagerConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ManagerConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmManagerConfigStorage) Update(c NtmModel.ManagerConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ManagerConfig with id does not exist")
	}

	return nil
}
