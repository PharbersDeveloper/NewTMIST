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

type NtmHospitalStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmHospitalStorage) NewHospitalStorage(args []BmDaemons.BmDaemon) *NtmHospitalStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmHospitalStorage{mdb}
}

func (s NtmHospitalStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.Hospital {
	in := NtmModel.Hospital{}
	var out []NtmModel.Hospital
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.Hospital
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

func (s NtmHospitalStorage) GetOne(id string) (NtmModel.Hospital, error) {
	in := NtmModel.Hospital{ID: id}
	out := NtmModel.Hospital{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Hospital for id %s not found", id)
	return NtmModel.Hospital{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *NtmHospitalStorage) Insert(c NtmModel.Hospital) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmHospitalStorage) Delete(id string) error {
	in := NtmModel.Hospital{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Hospital with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *NtmHospitalStorage) Update(c NtmModel.Hospital) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Hospital with id does not exist")
	}

	return nil
}

func (s *NtmHospitalStorage) Count(req api2go.Request, c NtmModel.Hospital) int {
	r, _ := s.db.Count(req, &c)
	return r
}