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

// NtmPaperinputStorage stores all of the tasty modelleaf, needs to be injected into
// Paperinput and Paperinput Resource. In the real world, you would use a database for that.
type NtmPaperinputStorage struct {
	images  map[string]*NtmModel.Paperinput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmPaperinputStorage) NewPaperinputStorage(args []BmDaemons.BmDaemon) *NtmPaperinputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmPaperinputStorage{make(map[string]*NtmModel.Paperinput), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmPaperinputStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.Paperinput {
	in := NtmModel.Paperinput{}
	var out []NtmModel.Paperinput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.Paperinput
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
func (s NtmPaperinputStorage) GetOne(id string) (NtmModel.Paperinput, error) {
	in := NtmModel.Paperinput{ID: id}
	out := NtmModel.Paperinput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Paperinput for id %s not found", id)
	return NtmModel.Paperinput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmPaperinputStorage) Insert(c NtmModel.Paperinput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmPaperinputStorage) Delete(id string) error {
	in := NtmModel.Paperinput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Paperinput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmPaperinputStorage) Update(c NtmModel.Paperinput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Paperinput with id does not exist")
	}

	return nil
}

func (s *NtmPaperinputStorage) Count(req api2go.Request, c NtmModel.Paperinput) int {
	r, _ := s.db.Count(req, &c)
	return r
}