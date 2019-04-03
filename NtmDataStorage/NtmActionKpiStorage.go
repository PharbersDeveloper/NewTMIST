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

// NtmActionKpiStorage stores all of the tasty modelleaf, needs to be injected into
// ActionKpi and ActionKpi Resource. In the real world, you would use a database for that.
type NtmActionKpiStorage struct {

	db *BmMongodb.BmMongodb
}

func (s NtmActionKpiStorage) NewActionKpiStorage(args []BmDaemons.BmDaemon) *NtmActionKpiStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmActionKpiStorage{ mdb}
}

// GetAll of the modelleaf
func (s NtmActionKpiStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.ActionKpi {
	in := NtmModel.ActionKpi{}
	var out []*NtmModel.ActionKpi
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s NtmActionKpiStorage) GetOne(id string) (NtmModel.ActionKpi, error) {
	in := NtmModel.ActionKpi{ID: id}
	out := NtmModel.ActionKpi{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ActionKpi for id %s not found", id)
	return NtmModel.ActionKpi{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmActionKpiStorage) Insert(c NtmModel.ActionKpi) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmActionKpiStorage) Delete(id string) error {
	in := NtmModel.ActionKpi{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ActionKpi with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmActionKpiStorage) Update(c NtmModel.ActionKpi) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ActionKpi with id does not exist")
	}

	return nil
}
