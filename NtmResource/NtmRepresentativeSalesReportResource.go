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

type NtmRepresentativeSalesReportResource struct {
	NtmRepresentativeSalesReportStorage *NtmDataStorage.NtmRepresentativeSalesReportStorage
	NtmSalesReportStorage               *NtmDataStorage.NtmSalesReportStorage
}

func (c NtmRepresentativeSalesReportResource) NewRepresentativeSalesReportResource(args []BmDataStorage.BmStorage) *NtmRepresentativeSalesReportResource {
	var psr  *NtmDataStorage.NtmRepresentativeSalesReportStorage
	var sr *NtmDataStorage.NtmSalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmRepresentativeSalesReportStorage" {
			psr = arg.(*NtmDataStorage.NtmRepresentativeSalesReportStorage)
		} else if tp.Name() == "NtmSalesReportStorage" {
			sr = arg.(*NtmDataStorage.NtmSalesReportStorage)
		}
	}
	return &NtmRepresentativeSalesReportResource{
		NtmRepresentativeSalesReportStorage: psr,
		NtmSalesReportStorage: sr,
	}
}

// FindAll SalesConfigs
func (c NtmRepresentativeSalesReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	salesReportsID, dcok := r.QueryParams["salesReportsID"]

	if dcok {
		modelRootID := salesReportsID[0]
		modelRoot, err := c.NtmSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		r.QueryParams["ids"] = modelRoot.RepresentativeSalesReportIDs

		model := c.NtmRepresentativeSalesReportStorage.GetAll(r, -1,-1)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []*NtmModel.RepresentativeSalesReport
	result = c.NtmRepresentativeSalesReportStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmRepresentativeSalesReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmRepresentativeSalesReportStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmRepresentativeSalesReportResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.RepresentativeSalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmRepresentativeSalesReportStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmRepresentativeSalesReportResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmRepresentativeSalesReportStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmRepresentativeSalesReportResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.RepresentativeSalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmRepresentativeSalesReportStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
