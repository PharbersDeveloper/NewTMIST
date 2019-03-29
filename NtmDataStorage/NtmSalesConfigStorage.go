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

// NtmSalesConfigStorage stores all of the tasty modelleaf, needs to be injected into
// SalesConfig and SalesConfig Resource. In the real world, you would use a database for that.
type NtmSalesConfigStorage struct {
	SalesConfigs  map[string]*NtmModel.SalesConfig
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmSalesConfigStorage) NewSalesConfigStorage(args []BmDaemons.BmDaemon) *NtmSalesConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmSalesConfigStorage{make(map[string]*NtmModel.SalesConfig), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmSalesConfigStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.SalesConfig {
	in := NtmModel.SalesConfig{}
	var out []NtmModel.SalesConfig
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
func (s NtmSalesConfigStorage) GetOne(id string) (NtmModel.SalesConfig, error) {
	in := NtmModel.SalesConfig{ID: id}
	out := NtmModel.SalesConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("SalesConfig for id %s not found", id)
	return NtmModel.SalesConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmSalesConfigStorage) Insert(c NtmModel.SalesConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmSalesConfigStorage) Delete(id string) error {
	in := NtmModel.SalesConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("SalesConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmSalesConfigStorage) Update(c NtmModel.SalesConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("SalesConfig with id does not exist")
	}

	return nil
}
