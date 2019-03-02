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

type NtmRegionStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmRegionStorage) NewRegionStorage(args []BmDaemons.BmDaemon) *NtmRegionStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmRegionStorage{mdb}
}

func (s NtmRegionStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.Region {
	in := NtmModel.Region{}
	var out []NtmModel.Region
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.Region
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

func (s NtmRegionStorage) GetOne(id string) (NtmModel.Region, error) {
	in := NtmModel.Region{ID: id}
	out := NtmModel.Region{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Region for id %s not found", id)
	return NtmModel.Region{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *NtmRegionStorage) Insert(c NtmModel.Region) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmRegionStorage) Delete(id string) error {
	in := NtmModel.Region{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Region with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *NtmRegionStorage) Update(c NtmModel.Region) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Region with id does not exist")
	}

	return nil
}

func (s *NtmRegionStorage) Count(req api2go.Request, c NtmModel.Region) int {
	r, _ := s.db.Count(req, &c)
	return r
}