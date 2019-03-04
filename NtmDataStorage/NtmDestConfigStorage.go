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

// NtmDestConfigStorage stores all of the tasty chocolate, needs to be injected into
// DestConfig Dest. In the real world, you would use a database for that.
type NtmDestConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmDestConfigStorage) NewDestConfigStorage(args []BmDaemons.BmDaemon) *NtmDestConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmDestConfigStorage{mdb}
}

// GetAll of the chocolate
func (s NtmDestConfigStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.DestConfig {
	in := NtmModel.DestConfig{}
	var out []NtmModel.DestConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.DestConfig
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
func (s NtmDestConfigStorage) GetOne(id string) (NtmModel.DestConfig, error) {
	in := NtmModel.DestConfig{ID: id}
	out := NtmModel.DestConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("DestConfig for id %s not found", id)
	return NtmModel.DestConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmDestConfigStorage) Insert(c NtmModel.DestConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmDestConfigStorage) Delete(id string) error {
	in := NtmModel.DestConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("DestConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *NtmDestConfigStorage) Update(c NtmModel.DestConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("DestConfig with id does not exist")
	}

	return nil
}

func (s *NtmDestConfigStorage) Count(req api2go.Request, c NtmModel.DestConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
