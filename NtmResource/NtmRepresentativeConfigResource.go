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

type NtmRepresentativeConfigResource struct {
	NtmRepresentativeConfigStorage *NtmDataStorage.NtmRepresentativeConfigStorage
	NtmResourceConfigStorage       *NtmDataStorage.NtmResourceConfigStorage
	NtmRepresentativeResource      *NtmRepresentativeResource
}

func (s NtmRepresentativeConfigResource) NewRepresentativeConfigResource(args []BmDataStorage.BmStorage) *NtmRepresentativeConfigResource {
	var repcs *NtmDataStorage.NtmRepresentativeConfigStorage
	var rcs *NtmDataStorage.NtmResourceConfigStorage
	var rr *NtmRepresentativeResource
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmRepresentativeConfigStorage" {
			repcs = arg.(*NtmDataStorage.NtmRepresentativeConfigStorage)
		} else if tp.Name() == "NtmResourceConfigStorage" {
			rcs = arg.(*NtmDataStorage.NtmResourceConfigStorage)
		} else if tp.Name() == "NtmRepresentativeResource" {
			rr = arg.(interface{}).(*NtmRepresentativeResource)
		}
	}
	return &NtmRepresentativeConfigResource{
		NtmRepresentativeConfigStorage: repcs,
		NtmResourceConfigStorage:       rcs,
		NtmRepresentativeResource:      rr,
	}
}

func (s NtmRepresentativeConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	resourceConfigsID, rcok := r.QueryParams["resourceConfigsID"]
	var result []NtmModel.RepresentativeConfig

	if rcok {
		modelRootID := resourceConfigsID[0]

		modelRoot, err := s.NtmResourceConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}

		model, err := s.NtmRepresentativeConfigStorage.GetOne(modelRoot.ResourceID)
		if err != nil {
			return &Response{}, err
		}
		result = append(result, model)

		return &Response{Res: result}, nil
	}

	models := s.NtmRepresentativeConfigStorage.GetAll(r, -1, -1)
	for _, model := range models {
		if model.RepresentativeID != "" {
			response, err := s.NtmRepresentativeResource.FindOne(model.RepresentativeID, api2go.Request{})
			item := response.Result()
			if err != nil {
				return &Response{}, err
			}
			model.Representative = item.(NtmModel.Representative)
		}

		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmRepresentativeConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.RepresentativeConfig
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
		for _, iter := range s.NtmRepresentativeConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmRepresentativeConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.RepresentativeConfig{}
	count := s.NtmRepresentativeConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmRepresentativeConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmRepresentativeConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	if model.RepresentativeID != "" {
		response, err := s.NtmRepresentativeResource.FindOne(model.RepresentativeID, api2go.Request{})
		item := response.Result()
		if err != nil {
			return &Response{}, err
		}
		model.Representative = item.(NtmModel.Representative)
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmRepresentativeConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.RepresentativeConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmRepresentativeConfigStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmRepresentativeConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmRepresentativeConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmRepresentativeConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.RepresentativeConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmRepresentativeConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
