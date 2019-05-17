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

// NtmGeneralPerformanceResultStorage stores all of the tasty modelleaf, needs to be injected into
// GeneralPerformanceResult and GeneralPerformanceResult Resource. In the real world, you would use a database for that.
type NtmGeneralPerformanceResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmGeneralPerformanceResultStorage) NewGeneralPerformanceResultStorage(args []BmDaemons.BmDaemon) *NtmGeneralPerformanceResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmGeneralPerformanceResultStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmGeneralPerformanceResultStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.GeneralPerformanceResult {
	in := NtmModel.GeneralPerformanceResult{}
	var out []NtmModel.GeneralPerformanceResult
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(&iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s NtmGeneralPerformanceResultStorage) GetOne(id string) (NtmModel.GeneralPerformanceResult, error) {
	in := NtmModel.GeneralPerformanceResult{ID: id}
	out := NtmModel.GeneralPerformanceResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("GeneralPerformanceResult for id %s not found", id)
	return NtmModel.GeneralPerformanceResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmGeneralPerformanceResultStorage) Insert(c NtmModel.GeneralPerformanceResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmGeneralPerformanceResultStorage) Delete(id string) error {
	in := NtmModel.GeneralPerformanceResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("GeneralPerformanceResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmGeneralPerformanceResultStorage) Update(c NtmModel.GeneralPerformanceResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("GeneralPerformanceResult with id does not exist")
	}

	return nil
}
