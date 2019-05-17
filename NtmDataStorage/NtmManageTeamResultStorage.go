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

// NtmManageTeamResultStorage stores all of the tasty modelleaf, needs to be injected into
// ManageTeamResult and ManageTeamResult Resource. In the real world, you would use a database for that.
type NtmManageTeamResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmManageTeamResultStorage) NewManageTeamResultStorage(args []BmDaemons.BmDaemon) *NtmManageTeamResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmManageTeamResultStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmManageTeamResultStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.ManageTeamResult {
	in := NtmModel.ManageTeamResult{}
	var out []NtmModel.ManageTeamResult
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
func (s NtmManageTeamResultStorage) GetOne(id string) (NtmModel.ManageTeamResult, error) {
	in := NtmModel.ManageTeamResult{ID: id}
	out := NtmModel.ManageTeamResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ManageTeamResult for id %s not found", id)
	return NtmModel.ManageTeamResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmManageTeamResultStorage) Insert(c NtmModel.ManageTeamResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmManageTeamResultStorage) Delete(id string) error {
	in := NtmModel.ManageTeamResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ManageTeamResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmManageTeamResultStorage) Update(c NtmModel.ManageTeamResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ManageTeamResult with id does not exist")
	}

	return nil
}
