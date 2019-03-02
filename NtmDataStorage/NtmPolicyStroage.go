package NtmDataStorage

import (
	"fmt"
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
)

// NtmPolicyStorage stores all of the tasty modelleaf, needs to be injected into
// Policy and Policy Resource. In the real world, you would use a database for that.
type NtmPolicyStorage struct {
	Policies  map[string]*NtmModel.Policy
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmPolicyStorage) NewPolicyStorage(args []BmDaemons.BmDaemon) *NtmPolicyStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmPolicyStorage{make(map[string]*NtmModel.Policy), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmPolicyStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.Policy {
	in := NtmModel.Policy{}
	var out []NtmModel.Policy
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
func (s NtmPolicyStorage) GetOne(id string) (NtmModel.Policy, error) {
	in := NtmModel.Policy{ID: id}
	out := NtmModel.Policy{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Policy for id %s not found", id)
	return NtmModel.Policy{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmPolicyStorage) Insert(c NtmModel.Policy) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmPolicyStorage) Delete(id string) error {
	in := NtmModel.Policy{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Policy with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmPolicyStorage) Update(c NtmModel.Policy) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Policy with id does not exist")
	}

	return nil
}
