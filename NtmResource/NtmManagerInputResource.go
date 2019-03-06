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

type NtmManagerInputResource struct {
	NtmManagerInputStorage *NtmDataStorage.NtmManagerInputStorage
	NtmPaperInputStorage   *NtmDataStorage.NtmPaperInputStorage
}

func (s NtmManagerInputResource) NewManagerInputResource(args []BmDataStorage.BmStorage) *NtmManagerInputResource {
	var bis *NtmDataStorage.NtmManagerInputStorage
	var pis *NtmDataStorage.NtmPaperInputStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmManagerInputStorage" {
			bis = arg.(*NtmDataStorage.NtmManagerInputStorage)
		} else if tp.Name() == "NtmPaperInputStorage" {
			pis = arg.(*NtmDataStorage.NtmPaperInputStorage)
		}
	}
	return &NtmManagerInputResource{
		NtmManagerInputStorage: bis,
		NtmPaperInputStorage:   pis,
	}
}

func (s NtmManagerInputResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	paperInputsID, piok := r.QueryParams["paperInputsID"]
	var result []*NtmModel.ManagerInput

	if piok {
		modelRootID := paperInputsID[0]

		modelRoot, err := s.NtmPaperInputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}

		r.QueryParams["ids"] = modelRoot.ManagerInputIDs
		result = s.NtmManagerInputStorage.GetAll(r, -1, -1)
		return &Response{Res: result}, nil
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmManagerInputResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.ManagerInput
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
		for _, iter := range s.NtmManagerInputStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmManagerInputStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.ManagerInput{}
	count := s.NtmManagerInputStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmManagerInputResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmManagerInputStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmManagerInputResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.ManagerInput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmManagerInputStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmManagerInputResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmManagerInputStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmManagerInputResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.ManagerInput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmManagerInputStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
