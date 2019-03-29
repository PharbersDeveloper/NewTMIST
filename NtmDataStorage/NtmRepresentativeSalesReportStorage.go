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

// NtmRepresentativeSalesReportStorage stores all of the tasty modelleaf, needs to be injected into
// Representativesalesreport and Representativesalesreport Resource. In the real world, you would use a database for that.
type NtmRepresentativeSalesReportStorage struct {
	SalesConfigs  map[string]*NtmModel.Representativesalesreport
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmRepresentativeSalesReportStorage) NewRepresentativeSalesReportStorage(args []BmDaemons.BmDaemon) *NtmRepresentativeSalesReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmRepresentativeSalesReportStorage{make(map[string]*NtmModel.Representativesalesreport), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmRepresentativeSalesReportStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.Representativesalesreport {
	in := NtmModel.Representativesalesreport{}
	var out []NtmModel.Representativesalesreport
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
func (s NtmRepresentativeSalesReportStorage) GetOne(id string) (NtmModel.Representativesalesreport, error) {
	in := NtmModel.Representativesalesreport{ID: id}
	out := NtmModel.Representativesalesreport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Representativesalesreport for id %s not found", id)
	return NtmModel.Representativesalesreport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmRepresentativeSalesReportStorage) Insert(c NtmModel.Representativesalesreport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmRepresentativeSalesReportStorage) Delete(id string) error {
	in := NtmModel.Representativesalesreport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Representativesalesreport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmRepresentativeSalesReportStorage) Update(c NtmModel.Representativesalesreport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Representativesalesreport with id does not exist")
	}

	return nil
}
