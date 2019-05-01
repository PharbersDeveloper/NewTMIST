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

// NtmScenarioStorage stores all of the tasty modelleaf, needs to be injected into
// Scenario and Scenario Resource. In the real world, you would use a database for that.
type NtmScenarioStorage struct {
	Policies map[string]*NtmModel.Scenario
	idCount  int

	db *BmMongodb.BmMongodb
}

func (s NtmScenarioStorage) NewScenarioStorage(args []BmDaemons.BmDaemon) *NtmScenarioStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmScenarioStorage{make(map[string]*NtmModel.Scenario), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmScenarioStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.Scenario {
	in := NtmModel.Scenario{}
	var out []NtmModel.Scenario
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
func (s NtmScenarioStorage) GetOne(id string) (NtmModel.Scenario, error) {
	in := NtmModel.Scenario{ID: id}
	out := NtmModel.Scenario{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Scenario for id %s not found", id)
	return NtmModel.Scenario{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmScenarioStorage) Insert(c NtmModel.Scenario) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmScenarioStorage) Delete(id string) error {
	in := NtmModel.Scenario{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Scenario with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmScenarioStorage) Update(c NtmModel.Scenario) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Scenario with id does not exist")
	}

	return nil
}
