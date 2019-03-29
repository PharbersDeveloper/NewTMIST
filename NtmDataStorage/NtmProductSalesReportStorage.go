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

// NtmProductSalesReportStorage stores all of the tasty modelleaf, needs to be injected into
// Productsalesreport and Productsalesreport Resource. In the real world, you would use a database for that.
type NtmProductSalesReportStorage struct {
	SalesConfigs  map[string]*NtmModel.Productsalesreport
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmProductSalesReportStorage) NewProductSalesReportStorage(args []BmDaemons.BmDaemon) *NtmProductSalesReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmProductSalesReportStorage{make(map[string]*NtmModel.Productsalesreport), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmProductSalesReportStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.Productsalesreport {
	in := NtmModel.Productsalesreport{}
	var out []NtmModel.Productsalesreport
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
func (s NtmProductSalesReportStorage) GetOne(id string) (NtmModel.Productsalesreport, error) {
	in := NtmModel.Productsalesreport{ID: id}
	out := NtmModel.Productsalesreport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Productsalesreport for id %s not found", id)
	return NtmModel.Productsalesreport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmProductSalesReportStorage) Insert(c NtmModel.Productsalesreport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmProductSalesReportStorage) Delete(id string) error {
	in := NtmModel.Productsalesreport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Productsalesreport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmProductSalesReportStorage) Update(c NtmModel.Productsalesreport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Productsalesreport with id does not exist")
	}

	return nil
}
