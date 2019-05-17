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

// NtmManageTimeResultStorage stores all of the tasty modelleaf, needs to be injected into
// ManageTimeResult and ManageTimeResult Resource. In the real world, you would use a database for that.
type NtmManageTimeResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmManageTimeResultStorage) NewManageTimeResultStorage(args []BmDaemons.BmDaemon) *NtmManageTimeResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmManageTimeResultStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmManageTimeResultStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.ManageTimeResult {
	in := NtmModel.ManageTimeResult{}
	var out []NtmModel.ManageTimeResult
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
func (s NtmManageTimeResultStorage) GetOne(id string) (NtmModel.ManageTimeResult, error) {
	in := NtmModel.ManageTimeResult{ID: id}
	out := NtmModel.ManageTimeResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ManageTimeResult for id %s not found", id)
	return NtmModel.ManageTimeResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmManageTimeResultStorage) Insert(c NtmModel.ManageTimeResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmManageTimeResultStorage) Delete(id string) error {
	in := NtmModel.ManageTimeResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ManageTimeResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmManageTimeResultStorage) Update(c NtmModel.ManageTimeResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ManageTimeResult with id does not exist")
	}

	return nil
}
