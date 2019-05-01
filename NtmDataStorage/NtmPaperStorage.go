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

type NtmPaperStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmPaperStorage) NewPaperStorage(args []BmDaemons.BmDaemon) *NtmPaperStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmPaperStorage{mdb}
}

func (s NtmPaperStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.Paper {
	in := NtmModel.Paper{}
	var out []NtmModel.Paper
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.Paper
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

func (s NtmPaperStorage) GetOne(id string) (NtmModel.Paper, error) {
	in := NtmModel.Paper{ID: id}
	out := NtmModel.Paper{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Paper for id %s not found", id)
	return NtmModel.Paper{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *NtmPaperStorage) Insert(c NtmModel.Paper) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmPaperStorage) Delete(id string) error {
	in := NtmModel.Paper{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Paper with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *NtmPaperStorage) Update(c NtmModel.Paper) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Paper with id does not exist")
	}

	return nil
}

func (s *NtmPaperStorage) Count(req api2go.Request, c NtmModel.Paper) int {
	r, _ := s.db.Count(req, &c)
	return r
}