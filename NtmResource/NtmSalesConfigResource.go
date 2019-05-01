package NtmResource

import (
	"errors"
	"Ntm/NtmDataStorage"
	"Ntm/NtmModel"
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
	NtmSalesReportResource 			*NtmSalesReportResource
	NtmDestConfigStorage 		*NtmDataStorage.NtmDestConfigStorage
	NtmHospitalConfigStorage 	*NtmDataStorage.NtmHospitalConfigStorage
	NtmGoodsConfigStorage 		*NtmDataStorage.NtmGoodsConfigStorage
}

func (c NtmSalesConfigResource) NewSalesConfigResource(args []BmDataStorage.BmStorage) *NtmSalesConfigResource {
	var sc *NtmDataStorage.NtmSalesConfigStorage
	var ps *NtmDataStorage.NtmPaperStorage
	var srr	*NtmSalesReportResource


	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmSalesConfigStorage" {
			sc = arg.(*NtmDataStorage.NtmSalesConfigStorage)
		} else if tp.Name() == "NtmPaperStorage" {
			ps = arg.(*NtmDataStorage.NtmPaperStorage)
		} else if tp.Name() == "NtmSalesReportResource" {
			srr = arg.(interface{}).(*NtmSalesReportResource)
		}
	}
	return &NtmSalesConfigResource{
		NtmSalesConfigStorage:	sc,
		NtmPaperStorage: ps,
		NtmSalesReportResource: srr,
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
			// 获取这个用户在关卡下最新的报告
			SalesReportIDs := paperModel[0].SalesReportIDs
			LastSalesReportID := SalesReportIDs[len(SalesReportIDs)-1:][0]
			SalesReportResponse, err := c.NtmSalesReportResource.FindOne(LastSalesReportID, r)
			if err != nil {
				return &Response{}, err
			}
			response := SalesReportResponse.Result()
			item := response.(NtmModel.SalesReport)

			for _, salesConfigModel := range result {
				salesConfigModel.SalesReportID = item.ID
				salesConfigModel.SalesReport = &item
			}
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
