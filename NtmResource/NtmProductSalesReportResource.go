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

type NtmProductSalesReportResource struct {
	NtmProductSalesReportStorage       *NtmDataStorage.NtmProductSalesReportStorage
	NtmSalesReportStorage               *NtmDataStorage.NtmSalesReportStorage
}

func (c NtmProductSalesReportResource) NewProductSalesReportResource(args []BmDataStorage.BmStorage) *NtmProductSalesReportResource {
	var psr  *NtmDataStorage.NtmProductSalesReportStorage
	var sr *NtmDataStorage.NtmSalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmProductSalesReportStorage" {
			psr = arg.(*NtmDataStorage.NtmProductSalesReportStorage)
		} else if tp.Name() == "NtmSalesReportStorage" {
			sr = arg.(*NtmDataStorage.NtmSalesReportStorage)
		}
	}
	return &NtmProductSalesReportResource{
		NtmProductSalesReportStorage: psr,
		NtmSalesReportStorage: sr,
	}
}

// FindAll SalesConfigs
func (c NtmProductSalesReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	salesReportsID, dcok := r.QueryParams["salesReportsID"]

	if dcok {
		modelRootID := salesReportsID[0]
		modelRoot, err := c.NtmSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		r.QueryParams["ids"] = modelRoot.ProductSalesReportIDs

		model := c.NtmProductSalesReportStorage.GetAll(r, -1,-1)


		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []*NtmModel.ProductSalesReport
	result = c.NtmProductSalesReportStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmProductSalesReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmProductSalesReportStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmProductSalesReportResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ProductSalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmProductSalesReportStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmProductSalesReportResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmProductSalesReportStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmProductSalesReportResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ProductSalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmProductSalesReportStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
