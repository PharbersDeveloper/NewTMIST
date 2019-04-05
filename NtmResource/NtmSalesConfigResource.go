package NtmResource

import (
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmDataStorage"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"net/http"
	"reflect"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type NtmSalesConfigResource struct {
	NtmSalesConfigStorage       *NtmDataStorage.NtmSalesConfigStorage
	NtmPaperStorage				*NtmDataStorage.NtmPaperStorage
	NtmSalesReportStorage 		*NtmDataStorage.NtmSalesReportStorage
	NtmHospitalSalesReportStorage *NtmDataStorage.NtmHospitalSalesReportStorage
}

func (c NtmSalesConfigResource) NewSalesConfigResource(args []BmDataStorage.BmStorage) *NtmSalesConfigResource {
	var sc *NtmDataStorage.NtmSalesConfigStorage
	var ps *NtmDataStorage.NtmPaperStorage
	var srs *NtmDataStorage.NtmSalesReportStorage
	var hsrs *NtmDataStorage.NtmHospitalSalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmSalesConfigStorage" {
			sc = arg.(*NtmDataStorage.NtmSalesConfigStorage)
		} else if tp.Name() == "NtmPaperStorage" {
			ps = arg.(*NtmDataStorage.NtmPaperStorage)
		} else if tp.Name() == "NtmSalesReportStorage" {
			srs = arg.(*NtmDataStorage.NtmSalesReportStorage)
		} else if tp.Name() == "NtmHospitalSalesReportStorage" {
			hsrs = arg.(*NtmDataStorage.NtmHospitalSalesReportStorage)
		}
	}
	return &NtmSalesConfigResource{
		NtmSalesConfigStorage:	sc,
		NtmPaperStorage: ps,
		NtmSalesReportStorage: srs,
		NtmHospitalSalesReportStorage: hsrs,
	}
}

// FindAll SalesConfigs
func (c NtmSalesConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*NtmModel.SalesConfig
	_, acok := r.QueryParams["account-id"]
	_, psok := r.QueryParams["proposal-id"]

	if acok && psok{
		result = c.NtmSalesConfigStorage.GetAll(r, -1, -1)
		paperModel := c.NtmPaperStorage.GetAll(r, -1, -1)


		if len(paperModel) > 0 {
			SalesReportIDs := paperModel[0].SalesReportIDs
			LastSalesReportID := SalesReportIDs[len(SalesReportIDs)-1:][0]
			SalesReport, _ := c.NtmSalesReportStorage.GetOne(LastSalesReportID)
			r.QueryParams["ids"] = SalesReport.HospitalSalesReportIDs
			HospitalSalesReports := c.NtmHospitalSalesReportStorage.GetAll(r, -1,-1)
			for _, salesConfigModel := range result {
				for _,  hospitalSalesReport:= range HospitalSalesReports {
					if salesConfigModel.DestConfigID == hospitalSalesReport.DestConfigID &&
						salesConfigModel.GoodsConfigID == hospitalSalesReport.GoodsConfigID {
						salesConfigModel.Sales = hospitalSalesReport.Sales
						salesConfigModel.Potential = hospitalSalesReport.Potential
					}
				}
			}
			//SalesReport.HospitalSalesReport = append(SalesReport.HospitalSalesReport, HospitalSalesReports...)
			//
			//for _, model := range result {
			//	model.SalesReportID = SalesReport.ID
			//	model.SalesReport = &SalesReport
			//}
		}
		return &Response{Res: result}, nil
	}


	result = c.NtmSalesConfigStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmSalesConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmSalesConfigStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmSalesConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.SalesConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmSalesConfigStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmSalesConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmSalesConfigStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmSalesConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.SalesConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmSalesConfigStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
