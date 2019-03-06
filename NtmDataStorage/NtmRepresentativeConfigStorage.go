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

// NtmRepresentativeConfigStorage stores all of the tasty chocolate, needs to be injected into
// RepresentativeConfig Resource. In the real world, you would use a database for that.
type NtmRepresentativeConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmRepresentativeConfigStorage) NewRepresentativeConfigStorage(args []BmDaemons.BmDaemon) *NtmRepresentativeConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmRepresentativeConfigStorage{mdb}
}

// GetAll of the chocolate
func (s NtmRepresentativeConfigStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.RepresentativeConfig {
	in := NtmModel.RepresentativeConfig{}
	var out []NtmModel.RepresentativeConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.RepresentativeConfig
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
func (s NtmRepresentativeConfigStorage) GetOne(id string) (NtmModel.RepresentativeConfig, error) {
	in := NtmModel.RepresentativeConfig{ID: id}
	out := NtmModel.RepresentativeConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("RepresentativeConfig for id %s not found", id)
	return NtmModel.RepresentativeConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmRepresentativeConfigStorage) Insert(c NtmModel.RepresentativeConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmRepresentativeConfigStorage) Delete(id string) error {
	in := NtmModel.RepresentativeConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("RepresentativeConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *NtmRepresentativeConfigStorage) Update(c NtmModel.RepresentativeConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("RepresentativeConfig with id does not exist")
	}

	return nil
}

func (s *NtmRepresentativeConfigStorage) Count(req api2go.Request, c NtmModel.RepresentativeConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
