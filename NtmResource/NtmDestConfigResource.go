package NtmResource

import (
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmDataStorage"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type NtmDestConfigResource struct {
	NtmDestConfigStorage    		*NtmDataStorage.NtmDestConfigStorage
	NtmRegionConfigResource       	*NtmRegionConfigResource
	NtmHospitalConfigResource		*NtmHospitalConfigResource
}

func (s NtmDestConfigResource) NewDestConfigResource(args []BmDataStorage.BmStorage) *NtmDestConfigResource {
	var dcs *NtmDataStorage.NtmDestConfigStorage
	var rcr *NtmRegionConfigResource
	var hcr *NtmHospitalConfigResource
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmDestConfigStorage" {
			dcs = arg.(*NtmDataStorage.NtmDestConfigStorage)
		} else if tp.Name() == "NtmRegionConfigResource" {
			rcr = arg.(interface{}).(*NtmRegionConfigResource)
		} else if tp.Name() == "NtmHospitalConfigResource" {
			hcr = arg.(interface{}).(*NtmHospitalConfigResource)
		}
	}
	return &NtmDestConfigResource{
		NtmDestConfigStorage:    	dcs,
		NtmRegionConfigResource: 	rcr,
		NtmHospitalConfigResource: 	hcr,
	}
}

func (s NtmDestConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.DestConfig
	models := s.NtmDestConfigStorage.GetAll(r, -1, -1)

	for _, model := range models {
		if model.DestType == 0 {
			rcr, err := s.NtmRegionConfigResource.FindOne(model.DestID, api2go.Request{})
			if err != nil {
				return &Response{}, err
			}
			model.RegionConfig = rcr.Result().(NtmModel.RegionConfig)
		} else if model.DestType == 1 {
			hcr, err := s.NtmHospitalConfigResource.FindOne(model.DestID, api2go.Request{})
			if err != nil {
				return &Response{}, err
			}
			model.HospitalConfig = hcr.Result().(NtmModel.HospitalConfig)
		}

		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmDestConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.DestConfig
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
		for _, iter := range s.NtmDestConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmDestConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.DestConfig{}
	count := s.NtmDestConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmDestConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmDestConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	if model.DestType == 0 {
		rcr, err := s.NtmRegionConfigResource.FindOne(model.DestID, api2go.Request{})
		if err != nil {
			return &Response{}, err
		}
		model.RegionConfig = rcr.Result().(NtmModel.RegionConfig)
	} else if model.DestType == 1 {
		hcr, err := s.NtmHospitalConfigResource.FindOne(model.DestID, api2go.Request{})
		if err != nil {
			return &Response{}, err
		}
		model.HospitalConfig = hcr.Result().(NtmModel.HospitalConfig)
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmDestConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.DestConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmDestConfigStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmDestConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmDestConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmDestConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.DestConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmDestConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
