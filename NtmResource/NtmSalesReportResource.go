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
	NtmPaperStorage						*NtmDataStorage.NtmPaperStorage
	NtmSalesConfigStorage				*NtmDataStorage.NtmSalesConfigStorage
}

func (c NtmSalesReportResource) NewSalesReportResource(args []BmDataStorage.BmStorage) *NtmSalesReportResource {
	var sr  *NtmDataStorage.NtmSalesReportStorage
	var hsr *NtmDataStorage.NtmHospitalSalesReportStorage
	var rsp *NtmDataStorage.NtmRepresentativeSalesReportStorage
	var psr *NtmDataStorage.NtmProductSalesReportStorage
	var ps	*NtmDataStorage.NtmPaperStorage
	var sc	*NtmDataStorage.NtmSalesConfigStorage

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
		} else if tp.Name() == "NtmPaperStorage" {
			ps = arg.(*NtmDataStorage.NtmPaperStorage)
		} else if tp.Name() == "NtmSalesConfigStorage" {
			sc = arg.(*NtmDataStorage.NtmSalesConfigStorage)
		}
	}
	return &NtmSalesReportResource{
		NtmSalesReportStorage : sr,
		NtmHospitalSalesReportStorage: hsr,
		NtmRepresentativeSalesReportStorage: rsp,
		NtmProductSalesReportStorage: psr,
		NtmPaperStorage: ps,
		NtmSalesConfigStorage: sc,
	}
}

// FindAll SalesConfigs
func (c NtmSalesReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	papersID, dcok := r.QueryParams["papersID"]

	//salesConfigsID, scok := r.QueryParams["salesConfigsID"]

	var result []NtmModel.SalesReport


	if dcok {
		modelRootID := papersID[0]
		modelRoot, err := c.NtmPaperStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.SalesReportIDs

		result := c.NtmSalesReportStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	//if scok {
	//	modelRootID := salesConfigsID[0]
	//	modelRoot, err := c.NtmSalesConfigStorage.GetOne(modelRootID)
	//
	//	if err != nil {
	//		return &Response{}, nil
	//	}
	//
	//	result, err := c.NtmSalesReportStorage.GetOne(modelRoot.SalesReportID)
	//
	//	if err != nil {
	//		return &Response{}, nil
	//	}
	//
	//	return &Response{Res: result}, nil
	//}

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
	choc, ok := obj.(NtmModel.SalesReport)
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
	choc, ok := obj.(NtmModel.SalesReport)
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
