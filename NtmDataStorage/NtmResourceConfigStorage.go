package NtmDataStorage

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"Ntm/NtmModel"
	"github.com/manyminds/api2go"
	"net/http"
)

// NtmResourceConfigStorage stores all of the tasty chocolate, needs to be injected into
// ResourceConfig Resource. In the real world, you would use a database for that.
type NtmResourceConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmResourceConfigStorage) NewResourceConfigStorage(args []BmDaemons.BmDaemon) *NtmResourceConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmResourceConfigStorage{mdb}
}

// GetAll of the chocolate
func (s NtmResourceConfigStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.ResourceConfig {
	in := NtmModel.ResourceConfig{}
	var out []NtmModel.ResourceConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.ResourceConfig
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Student)
	}
}

// GetOne
func (s NtmResourceConfigStorage) GetOne(id string) (NtmModel.ResourceConfig, error) {
	in := NtmModel.ResourceConfig{ID: id}
	out := NtmModel.ResourceConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ResourceConfig for id %s not found", id)
	return NtmModel.ResourceConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmResourceConfigStorage) Insert(c NtmModel.ResourceConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmResourceConfigStorage) Delete(id string) error {
	in := NtmModel.ResourceConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ResourceConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *NtmResourceConfigStorage) Update(c NtmModel.ResourceConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ResourceConfig with id does not exist")
	}

	return nil
}

func (s *NtmResourceConfigStorage) Count(req api2go.Request, c NtmModel.ResourceConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
