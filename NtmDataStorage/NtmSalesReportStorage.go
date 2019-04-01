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
// SalesReport and SalesReport Resource. In the real world, you would use a database for that.
type NtmSalesReportStorage struct {
	SalesConfigs  map[string]*NtmModel.SalesReport
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmSalesReportStorage) NewSalesReportStorage(args []BmDaemons.BmDaemon) *NtmSalesReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmSalesReportStorage{make(map[string]*NtmModel.SalesReport), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmSalesReportStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.SalesReport {
	in := NtmModel.SalesReport{}
	var out []NtmModel.SalesReport
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.SalesReport
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
func (s NtmSalesReportStorage) GetOne(id string) (NtmModel.SalesReport, error) {
	in := NtmModel.SalesReport{ID: id}
	out := NtmModel.SalesReport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("SalesReport for id %s not found", id)
	return NtmModel.SalesReport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmSalesReportStorage) Insert(c NtmModel.SalesReport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmSalesReportStorage) Delete(id string) error {
	in := NtmModel.SalesReport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("SalesReport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmSalesReportStorage) Update(c NtmModel.SalesReport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("SalesReport with id does not exist")
	}

	return nil
}

func (s *NtmSalesReportStorage) Count(req api2go.Request, c NtmModel.SalesReport) int {
	r, _ := s.db.Count(req, &c)
	return r
}
