package NtmDataStorage

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"github.com/manyminds/api2go"
	"net/http"
)

// NtmProductConfigStorage stores all of the tasty chocolate, needs to be injected into
// ProductConfig Resource. In the real world, you would use a database for that.
type NtmProductConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmProductConfigStorage) NewProductConfigStorage(args []BmDaemons.BmDaemon) *NtmProductConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmProductConfigStorage{mdb}
}

// GetAll of the chocolate
func (s NtmProductConfigStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.ProductConfig {
	in := NtmModel.ProductConfig{}
	var out []NtmModel.ProductConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.ProductConfig
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
func (s NtmProductConfigStorage) GetOne(id string) (NtmModel.ProductConfig, error) {
	in := NtmModel.ProductConfig{ID: id}
	out := NtmModel.ProductConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ProductConfig for id %s not found", id)
	return NtmModel.ProductConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmProductConfigStorage) Insert(c NtmModel.ProductConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmProductConfigStorage) Delete(id string) error {
	in := NtmModel.ProductConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ProductConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *NtmProductConfigStorage) Update(c NtmModel.ProductConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ProductConfig with id does not exist")
	}

	return nil
}

func (s *NtmProductConfigStorage) Count(req api2go.Request, c NtmModel.ProductConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
