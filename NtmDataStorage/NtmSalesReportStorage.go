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

// NtmSalesReportStorage stores all of the tasty modelleaf, needs to be injected into
// Salesreport and Salesreport Resource. In the real world, you would use a database for that.
type NtmSalesReportStorage struct {
	SalesConfigs  map[string]*NtmModel.Salesreport
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmSalesReportStorage) NewSalesReportStorage(args []BmDaemons.BmDaemon) *NtmSalesReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmSalesReportStorage{make(map[string]*NtmModel.Salesreport), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmSalesReportStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.Salesreport {
	in := NtmModel.Salesreport{}
	var out []NtmModel.Salesreport
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.Salesreport
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Student)
	}
}

// GetOne tasty modelleaf
func (s NtmSalesReportStorage) GetOne(id string) (NtmModel.Salesreport, error) {
	in := NtmModel.Salesreport{ID: id}
	out := NtmModel.Salesreport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Salesreport for id %s not found", id)
	return NtmModel.Salesreport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmSalesReportStorage) Insert(c NtmModel.Salesreport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmSalesReportStorage) Delete(id string) error {
	in := NtmModel.Salesreport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Salesreport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmSalesReportStorage) Update(c NtmModel.Salesreport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Salesreport with id does not exist")
	}

	return nil
}

func (s *NtmSalesReportStorage) Count(req api2go.Request, c NtmModel.Salesreport) int {
	r, _ := s.db.Count(req, &c)
	return r
}
