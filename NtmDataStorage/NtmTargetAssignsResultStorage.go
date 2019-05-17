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

// NtmTargetAssignsResultStorage stores all of the tasty modelleaf, needs to be injected into
// TargetAssignsResult and TargetAssignsResult Resource. In the real world, you would use a database for that.
type NtmTargetAssignsResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmTargetAssignsResultStorage) NewTargetAssignsResultStorage(args []BmDaemons.BmDaemon) *NtmTargetAssignsResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmTargetAssignsResultStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmTargetAssignsResultStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.TargetAssignsResult {
	in := NtmModel.TargetAssignsResult{}
	var out []NtmModel.TargetAssignsResult
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
func (s NtmTargetAssignsResultStorage) GetOne(id string) (NtmModel.TargetAssignsResult, error) {
	in := NtmModel.TargetAssignsResult{ID: id}
	out := NtmModel.TargetAssignsResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("TargetAssignsResult for id %s not found", id)
	return NtmModel.TargetAssignsResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmTargetAssignsResultStorage) Insert(c NtmModel.TargetAssignsResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmTargetAssignsResultStorage) Delete(id string) error {
	in := NtmModel.TargetAssignsResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("TargetAssignsResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmTargetAssignsResultStorage) Update(c NtmModel.TargetAssignsResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("TargetAssignsResult with id does not exist")
	}

	return nil
}
