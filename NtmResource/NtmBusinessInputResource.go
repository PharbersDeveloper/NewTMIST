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

type NtmBusinessInputResource struct {
	NtmBusinessInputStorage *NtmDataStorage.NtmBusinessInputStorage
	NtmPaperInputStorage    *NtmDataStorage.NtmPaperInputStorage
}

func (s NtmBusinessInputResource) NewBusinessInputResource(args []BmDataStorage.BmStorage) *NtmBusinessInputResource {
	var bis *NtmDataStorage.NtmBusinessInputStorage
	var pis *NtmDataStorage.NtmPaperInputStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmBusinessInputStorage" {
			bis = arg.(*NtmDataStorage.NtmBusinessInputStorage)
		} else if tp.Name() == "NtmPaperInputStorage" {
			pis = arg.(*NtmDataStorage.NtmPaperInputStorage)
		}
	}
	return &NtmBusinessInputResource{
		NtmBusinessInputStorage: bis,
		NtmPaperInputStorage:    pis,
	}
}

func (s NtmBusinessInputResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	paperInputsID, piok := r.QueryParams["paperInputsID"]
	var result []*NtmModel.BusinessInput

	if piok {
		modelRootID := paperInputsID[0]

		modelRoot, err := s.NtmPaperInputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		r.QueryParams["ids"] = modelRoot.BusinessInputIDs


		result = s.NtmBusinessInputStorage.GetAll(r, -1, -1)

		return &Response{Res: result}, nil
	}

	models := s.NtmBusinessInputStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, model)
	}
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmBusinessInputResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.BusinessInput
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
		for _, iter := range s.NtmBusinessInputStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmBusinessInputStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.BusinessInput{}
	count := s.NtmBusinessInputStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmBusinessInputResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmBusinessInputStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmBusinessInputResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.BusinessInput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmBusinessInputStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmBusinessInputResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmBusinessInputStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmBusinessInputResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.BusinessInput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmBusinessInputStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
