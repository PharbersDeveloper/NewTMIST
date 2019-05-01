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

// NtmHospitalSalesReportStorage stores all of the tasty modelleaf, needs to be injected into
// HospitalSalesReport and HospitalSalesReport Resource. In the real world, you would use a database for that.
type NtmHospitalSalesReportStorage struct {
	SalesConfigs  map[string]*NtmModel.HospitalSalesReport
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmHospitalSalesReportStorage) NewHospitalSalesReportStorage(args []BmDaemons.BmDaemon) *NtmHospitalSalesReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmHospitalSalesReportStorage{make(map[string]*NtmModel.HospitalSalesReport), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmHospitalSalesReportStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.HospitalSalesReport {
	in := NtmModel.HospitalSalesReport{}
	var out []NtmModel.HospitalSalesReport
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.HospitalSalesReport
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
func (s NtmHospitalSalesReportStorage) GetOne(id string) (NtmModel.HospitalSalesReport, error) {
	in := NtmModel.HospitalSalesReport{ID: id}
	out := NtmModel.HospitalSalesReport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("HospitalSalesReport for id %s not found", id)
	return NtmModel.HospitalSalesReport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmHospitalSalesReportStorage) Insert(c NtmModel.HospitalSalesReport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmHospitalSalesReportStorage) Delete(id string) error {
	in := NtmModel.HospitalSalesReport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("HospitalSalesReport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmHospitalSalesReportStorage) Update(c NtmModel.HospitalSalesReport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("HospitalSalesReport with id does not exist")
	}

	return nil
}
