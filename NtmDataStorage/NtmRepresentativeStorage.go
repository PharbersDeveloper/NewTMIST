package NtmDataStorage

import (
	"errors"
	"fmt"
	"Ntm/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
	"net/http"
)

type NtmRepresentativeStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmRepresentativeStorage) NewRepresentativeStorage(args []BmDaemons.BmDaemon) *NtmRepresentativeStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmRepresentativeStorage{mdb}
}

func (s NtmRepresentativeStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.Representative {
	in := NtmModel.Representative{}
	var out []NtmModel.Representative
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.Representative
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

func (s NtmRepresentativeStorage) GetOne(id string) (NtmModel.Representative, error) {
	in := NtmModel.Representative{ID: id}
	out := NtmModel.Representative{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Representative for id %s not found", id)
	return NtmModel.Representative{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *NtmRepresentativeStorage) Insert(c NtmModel.Representative) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmRepresentativeStorage) Delete(id string) error {
	in := NtmModel.Representative{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Representative with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *NtmRepresentativeStorage) Update(c NtmModel.Representative) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Representative with id does not exist")
	}

	return nil
}

func (s *NtmRepresentativeStorage) Count(req api2go.Request, c NtmModel.Representative) int {
	r, _ := s.db.Count(req, &c)
	return r
}
