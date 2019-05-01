package NtmResource

import (
	"errors"
	"Ntm/NtmDataStorage"
	"Ntm/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type NtmGoodsConfigResource struct {
	NtmGoodsConfigStorage   			*NtmDataStorage.NtmGoodsConfigStorage
	NtmProductConfigStorage 			*NtmDataStorage.NtmProductConfigStorage
	NtmProductSalesReportStorage   		*NtmDataStorage.NtmProductSalesReportStorage
	NtmHospitalSalesReportStorage		*NtmDataStorage.NtmHospitalSalesReportStorage
	NtmRepresentativeSalesReportStorage	*NtmDataStorage.NtmRepresentativeSalesReportStorage
	NtmSalesConfigStorage 				*NtmDataStorage.NtmSalesConfigStorage
	NtmBusinessinputStorage				*NtmDataStorage.NtmBusinessinputStorage
}

func (s NtmGoodsConfigResource) NewGoodsConfigResource(args []BmDataStorage.BmStorage) *NtmGoodsConfigResource {
	var gcs *NtmDataStorage.NtmGoodsConfigStorage
	var pcs *NtmDataStorage.NtmProductConfigStorage
	var psr *NtmDataStorage.NtmProductSalesReportStorage
	var hsr *NtmDataStorage.NtmHospitalSalesReportStorage
	var rsr *NtmDataStorage.NtmRepresentativeSalesReportStorage
	var sc *NtmDataStorage.NtmSalesConfigStorage
	var bis *NtmDataStorage.NtmBusinessinputStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmGoodsConfigStorage" {
			gcs = arg.(*NtmDataStorage.NtmGoodsConfigStorage)
		} else if tp.Name() == "NtmProductConfigStorage" {
			pcs = arg.(interface{}).(*NtmDataStorage.NtmProductConfigStorage)
		} else if tp.Name() == "NtmProductSalesReportStorage" {
			psr = arg.(*NtmDataStorage.NtmProductSalesReportStorage)
		} else if tp.Name() == "NtmHospitalSalesReportStorage" {
			hsr = arg.(*NtmDataStorage.NtmHospitalSalesReportStorage)
		} else if tp.Name() == "NtmRepresentativeSalesReportStorage" {
			rsr = arg.(*NtmDataStorage.NtmRepresentativeSalesReportStorage)
		} else if tp.Name() == "NtmSalesConfigStorage" {
			sc = arg.(*NtmDataStorage.NtmSalesConfigStorage)
		} else if tp.Name() == "NtmBusinessinputStorage" {
			bis = arg.(*NtmDataStorage.NtmBusinessinputStorage)
		}
	}
	return &NtmGoodsConfigResource{
		NtmGoodsConfigStorage:   gcs,
		NtmProductConfigStorage: pcs,
		NtmProductSalesReportStorage: psr,
		NtmHospitalSalesReportStorage : hsr,
		NtmRepresentativeSalesReportStorage: rsr,
		NtmSalesConfigStorage: sc,
		NtmBusinessinputStorage: bis,
	}
}

func (s NtmGoodsConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	productSalesReportsID, psrok := r.QueryParams["productSalesReportsID"]
	hospitalSalesReportsID, hsrok := r.QueryParams["hospitalSalesReportsID"]
	representativeSalesReportsID, rsrok := r.QueryParams["representativeSalesReportsID"]
	salesConfigsID, scok := r.QueryParams["salesConfigsID"]
	businessinputsID, bok := r.QueryParams["businessinputsID"]

	if psrok {
		modelRootID := productSalesReportsID[0]
		modelRoot, err := s.NtmProductSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err:= s.NtmGoodsConfigStorage.GetOne(modelRoot.GoodsConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if hsrok {
		modelRootID := hospitalSalesReportsID[0]
		modelRoot, err := s.NtmHospitalSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err:= s.NtmGoodsConfigStorage.GetOne(modelRoot.GoodsConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if rsrok {
		modelRootID := representativeSalesReportsID[0]
		modelRoot, err := s.NtmRepresentativeSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err:= s.NtmGoodsConfigStorage.GetOne(modelRoot.GoodsConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if scok {
		modelRootID := salesConfigsID[0]
		modelRoot, err := s.NtmSalesConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err:= s.NtmGoodsConfigStorage.GetOne(modelRoot.GoodsConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if bok {
		modelRootID := businessinputsID[0]
		modelRoot, err := s.NtmBusinessinputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		result, err := s.NtmGoodsConfigStorage.GetOne(modelRoot.GoodsConfigId)

		if err != nil {
			return &Response{}, nil
		}

		return &Response{Res: result}, nil
	}

	var result []*NtmModel.GoodsConfig
	models := s.NtmGoodsConfigStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmGoodsConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.GoodsConfig
		number, size, offset, limit string
	)

	numberQuery, ok := r.QueryParams["page[number]"]
	if ok {
		number = numberQuery[0]
	}
	sizeQuery, ok := r.QueryParams["page[size]"]
	if ok {
		size = sizeQuery[0]
	}
	offsetQuery, ok := r.QueryParams["page[offset]"]
	if ok {
		offset = offsetQuery[0]
	}
	limitQuery, ok := r.QueryParams["page[limit]"]
	if ok {
		limit = limitQuery[0]
	}

	if size != "" {
		sizeI, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		numberI, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		start := sizeI * (numberI - 1)
		for _, iter := range s.NtmGoodsConfigStorage.GetAll(r, int(start), int(sizeI)) {
			result = append(result, *iter)
		}

	} else {
		limitI, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		offsetI, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		for _, iter := range s.NtmGoodsConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.GoodsConfig{}
	count := s.NtmGoodsConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmGoodsConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmGoodsConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	if model.GoodsType == 0 {
		resp, err := s.NtmProductConfigStorage.GetOne(model.GoodsID)
		if err != nil {
			return &Response{}, err
		}
		model.ProductConfig = &resp
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmGoodsConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.GoodsConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}
	id := s.NtmGoodsConfigStorage.Insert(model)
	model.ID = id
	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmGoodsConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmGoodsConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmGoodsConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.GoodsConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}
	err := s.NtmGoodsConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
