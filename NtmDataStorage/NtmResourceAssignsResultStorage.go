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

// NtmResourceAssignsResultStorage stores all of the tasty modelleaf, needs to be injected into
// ResourceAssignsResult and ResourceAssignsResult Resource. In the real world, you would use a database for that.
type NtmResourceAssignsResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmResourceAssignsResultStorage) NewResourceAssignsResultStorage(args []BmDaemons.BmDaemon) *NtmResourceAssignsResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmResourceAssignsResultStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmResourceAssignsResultStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.ResourceAssignsResult {
	in := NtmModel.ResourceAssignsResult{}
	var out []NtmModel.ResourceAssignsResult
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
func (s NtmResourceAssignsResultStorage) GetOne(id string) (NtmModel.ResourceAssignsResult, error) {
	in := NtmModel.ResourceAssignsResult{ID: id}
	out := NtmModel.ResourceAssignsResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ResourceAssignsResult for id %s not found", id)
	return NtmModel.ResourceAssignsResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmResourceAssignsResultStorage) Insert(c NtmModel.ResourceAssignsResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmResourceAssignsResultStorage) Delete(id string) error {
	in := NtmModel.ResourceAssignsResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ResourceAssignsResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmResourceAssignsResultStorage) Update(c NtmModel.ResourceAssignsResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ResourceAssignsResult with id does not exist")
	}

	return nil
}
