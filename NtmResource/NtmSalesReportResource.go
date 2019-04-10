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

	models := c.NtmSalesReportStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, *model)
	}
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmSalesReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	modelRoot, err := c.NtmSalesReportStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	modelRoot.HospitalSalesReport = []*NtmModel.HospitalSalesReport{}
	r.QueryParams["ids"] = modelRoot.HospitalSalesReportIDs
	for _, hospitalSalesReport := range c.NtmHospitalSalesReportStorage.GetAll(r, -1,-1) {
		modelRoot.HospitalSalesReport = append(modelRoot.HospitalSalesReport, hospitalSalesReport)
	}

	modelRoot.RepresentativeSalesReport = []*NtmModel.RepresentativeSalesReport{}
	r.QueryParams["ids"] = modelRoot.RepresentativeSalesReportIDs
	for _, representativeSalesReport := range c.NtmRepresentativeSalesReportStorage.GetAll(r, -1,-1) {
		modelRoot.RepresentativeSalesReport = append(modelRoot.RepresentativeSalesReport, representativeSalesReport)
	}

	modelRoot.ProductSalesReport = []*NtmModel.ProductSalesReport{}
	r.QueryParams["ids"] = modelRoot.ProductSalesReportIDs
	for _, productSalesReport := range c.NtmProductSalesReportStorage.GetAll(r, -1,-1) {
		modelRoot.ProductSalesReport = append(modelRoot.ProductSalesReport, productSalesReport)
	}

	return &Response{Res: modelRoot}, err
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
