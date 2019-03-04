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

type NtmProductConfigResource struct {
	NtmProductConfigStorage *NtmDataStorage.NtmProductConfigStorage
	NtmProductResource      *NtmProductResource
}

func (s NtmProductConfigResource) NewProductConfigResource(args []BmDataStorage.BmStorage) *NtmProductConfigResource {
	var hs *NtmDataStorage.NtmProductConfigStorage
	var pr *NtmProductResource
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmProductConfigStorage" {
			hs = arg.(*NtmDataStorage.NtmProductConfigStorage)
		} else if tp.Name() == "NtmProductResource" {
			pr = arg.(interface{}).(*NtmProductResource)
		}
	}
	return &NtmProductConfigResource{
		NtmProductConfigStorage: hs,
		NtmProductResource:      pr,
	}
}

func (s NtmProductConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.ProductConfig
	models := s.NtmProductConfigStorage.GetAll(r, -1, -1)

	for _, model := range models {
		if model.ProductID != "" {
			response, err := s.NtmProductResource.FindOne(model.ProductID, api2go.Request{})
			item := response.Result()
			if err != nil {
				return &Response{}, err
			}
			model.Product = item.(NtmModel.Product)
		}

		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmProductConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.ProductConfig
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
		for _, iter := range s.NtmProductConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmProductConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.ProductConfig{}
	count := s.NtmProductConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmProductConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmProductConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	if model.ProductID != "" {
		response, err := s.NtmProductResource.FindOne(model.ProductID, api2go.Request{})
		item := response.Result()
		if err != nil {
			return &Response{}, err
		}
		model.Product = item.(NtmModel.Product)
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmProductConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.ProductConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmProductConfigStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmProductConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmProductConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmProductConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.ProductConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmProductConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
