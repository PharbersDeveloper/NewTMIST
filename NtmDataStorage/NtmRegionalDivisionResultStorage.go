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

// NtmRegionalDivisionResultStorage stores all of the tasty modelleaf, needs to be injected into
// RegionalDivisionResult and RegionalDivisionResult Resource. In the real world, you would use a database for that.
type NtmRegionalDivisionResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmRegionalDivisionResultStorage) NewRegionalDivisionResultStorage(args []BmDaemons.BmDaemon) *NtmRegionalDivisionResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmRegionalDivisionResultStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmRegionalDivisionResultStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.RegionalDivisionResult {
	in := NtmModel.RegionalDivisionResult{}
	var out []NtmModel.RegionalDivisionResult
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
func (s NtmRegionalDivisionResultStorage) GetOne(id string) (NtmModel.RegionalDivisionResult, error) {
	in := NtmModel.RegionalDivisionResult{ID: id}
	out := NtmModel.RegionalDivisionResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("RegionalDivisionResult for id %s not found", id)
	return NtmModel.RegionalDivisionResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmRegionalDivisionResultStorage) Insert(c NtmModel.RegionalDivisionResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmRegionalDivisionResultStorage) Delete(id string) error {
	in := NtmModel.RegionalDivisionResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("RegionalDivisionResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmRegionalDivisionResultStorage) Update(c NtmModel.RegionalDivisionResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("RegionalDivisionResult with id does not exist")
	}

	return nil
}
