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

type NtmResourceConfigResource struct {
	NtmResourceConfigStorage       *NtmDataStorage.NtmResourceConfigStorage
	NtmManagerConfigStorage        *NtmDataStorage.NtmManagerConfigStorage
	NtmRepresentativeConfigStorage *NtmDataStorage.NtmRepresentativeConfigStorage
}

func (s NtmResourceConfigResource) NewResourceConfigResource(args []BmDataStorage.BmStorage) NtmResourceConfigResource {
	var rcs *NtmDataStorage.NtmResourceConfigStorage
	var mcs *NtmDataStorage.NtmManagerConfigStorage
	var repcs *NtmDataStorage.NtmRepresentativeConfigStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmResourceConfigStorage" {
			rcs = arg.(*NtmDataStorage.NtmResourceConfigStorage)
		} else if tp.Name() == "NtmManagerConfigStorage" {
			mcs = arg.(*NtmDataStorage.NtmManagerConfigStorage)
		} else if tp.Name() == "NtmRepresentativeConfigStorage" {
			repcs = arg.(*NtmDataStorage.NtmRepresentativeConfigStorage)
		}
	}
	return NtmResourceConfigResource{
		NtmResourceConfigStorage: rcs,
		NtmManagerConfigStorage: mcs,
		NtmRepresentativeConfigStorage: repcs,
	}
}

func (s NtmResourceConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.ResourceConfig
	models := s.NtmResourceConfigStorage.GetAll(r, -1, -1)

	for _, model := range models {
		if model.ResourceType == 0 {
			r, err := s.NtmManagerConfigStorage.GetOne(model.ResourceID)
			if err != nil {
				return &Response{}, err
			}
			model.ManagerConfig = r
		} else if model.ResourceType == 1 {
			r, err := s.NtmRepresentativeConfigStorage.GetOne(model.ResourceID)
			if err != nil {
				return &Response{}, err
			}
			model.RepresentativeConfig = r
		}

		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmResourceConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.ResourceConfig
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
		for _, iter := range s.NtmResourceConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmResourceConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.ResourceConfig{}
	count := s.NtmResourceConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmResourceConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmResourceConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	if model.ResourceType == 0 {
		r, err := s.NtmManagerConfigStorage.GetOne(model.ResourceID)
		if err != nil {
			return &Response{}, err
		}
		model.ManagerConfig = r
	} else if model.ResourceType == 1 {
		r, err := s.NtmRepresentativeConfigStorage.GetOne(model.ResourceID)
		if err != nil {
			return &Response{}, err
		}
		model.RepresentativeConfig = r
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmResourceConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.ResourceConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmResourceConfigStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmResourceConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmResourceConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmResourceConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.ResourceConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmResourceConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
