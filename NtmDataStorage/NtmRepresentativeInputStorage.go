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

// NtmRepresentativeinputStorage stores all of the tasty modelleaf, needs to be injected into
// Representativeinput and Representativeinput Resource. In the real world, you would use a database for that.
type NtmRepresentativeinputStorage struct {
	images  map[string]*NtmModel.Representativeinput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmRepresentativeinputStorage) NewRepresentativeinputStorage(args []BmDaemons.BmDaemon) *NtmRepresentativeinputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmRepresentativeinputStorage{make(map[string]*NtmModel.Representativeinput), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmRepresentativeinputStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.Representativeinput {
	in := NtmModel.Representativeinput{}
	var out []NtmModel.Representativeinput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.Representativeinput
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
func (s NtmRepresentativeinputStorage) GetOne(id string) (NtmModel.Representativeinput, error) {
	in := NtmModel.Representativeinput{ID: id}
	out := NtmModel.Representativeinput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Representativeinput for id %s not found", id)
	return NtmModel.Representativeinput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmRepresentativeinputStorage) Insert(c NtmModel.Representativeinput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmRepresentativeinputStorage) Delete(id string) error {
	in := NtmModel.Representativeinput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Representativeinput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmRepresentativeinputStorage) Update(c NtmModel.Representativeinput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Representativeinput with id does not exist")
	}

	return nil
}

func (s *NtmRepresentativeinputStorage) Count(req api2go.Request, c NtmModel.Representativeinput) int {
	r, _ := s.db.Count(req, &c)
	return r
}