package NtmResource

import (
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmDataStorage"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"reflect"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type NtmSalesReportResource struct {
	NtmSalesReportStorage               *NtmDataStorage.NtmSalesReportStorage
	NtmHospitalSalesReportStorage       *NtmDataStorage.NtmHospitalSalesReportStorage
	NtmRepresentativeSalesReportStorage *NtmDataStorage.NtmRepresentativeSalesReportStorage
	NtmProductSalesReportStorage        *NtmDataStorage.NtmProductSalesReportStorage
}

func (c NtmSalesReportResource) NewSalesReportResource(args []BmDataStorage.BmStorage) *NtmSalesReportResource {
	var sr  *NtmDataStorage.NtmSalesReportStorage
	var hsr *NtmDataStorage.NtmHospitalSalesReportStorage
	var rsp *NtmDataStorage.NtmRepresentativeSalesReportStorage
	var psr *NtmDataStorage.NtmProductSalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmSalesReportStorage" {
			sr = arg.(*NtmDataStorage.NtmSalesReportStorage)
		} else if tp.Name() == "NtmHospitalSalesReportStorage" {
			hsr = arg.(*NtmDataStorage.NtmHospitalSalesReportStorage)
		} else if tp.Name() == "NtmRepresentativeSalesReportStorage" {
			rsp = arg.(*NtmDataStorage.NtmRepresentativeSalesReportStorage)
		} else if tp.Name() == "NtmProductSalesReportStorage" {
			psr = arg.(*NtmDataStorage.NtmProductSalesReportStorage)
		}
	}
	return &NtmSalesReportResource{
		NtmSalesReportStorage : sr,
		NtmHospitalSalesReportStorage: hsr,
		NtmRepresentativeSalesReportStorage: rsp,
		NtmProductSalesReportStorage: psr,
	}
}

// FindAll SalesConfigs
func (c NtmSalesReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.Salesreport


	models := c.NtmSalesReportStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, *model)
	}
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmSalesReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmSalesReportStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmSalesReportResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Salesreport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmSalesReportStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmSalesReportResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmSalesReportStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmSalesReportResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Salesreport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmSalesReportStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
