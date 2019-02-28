package NtmDataStorage

import (
	"errors"
	"fmt"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
	"net/http"
)

type NtmProductStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmProductStorage) NewProductStorage(args []BmDaemons.BmDaemon) *NtmProductStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmProductStorage{mdb}
}

func (s NtmProductStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.Product {
	in := NtmModel.Product{}
	var out []NtmModel.Product
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.Product
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

func (s NtmProductStorage) GetOne(id string) (NtmModel.Product, error) {
	in := NtmModel.Product{ID: id}
	out := NtmModel.Product{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Product for id %s not found", id)
	return NtmModel.Product{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *NtmProductStorage) Insert(c NtmModel.Product) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmProductStorage) Delete(id string) error {
	in := NtmModel.Product{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Product with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *NtmProductStorage) Update(c NtmModel.Product) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Product with id does not exist")
	}

	return nil
}

func (s *NtmProductStorage) Count(req api2go.Request, c NtmModel.Product) int {
	r, _ := s.db.Count(req, &c)
	return r
}