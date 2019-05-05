package NtmResource

import (
	"Ntm/NtmDataStorage"
	"Ntm/NtmModel"
	"errors"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type NtmBusinessinputResource struct {
	NtmBusinessinputStorage *NtmDataStorage.NtmBusinessinputStorage
	NtmPaperinputStorage    *NtmDataStorage.NtmPaperinputStorage
}

func (s NtmBusinessinputResource) NewBusinessinputResource(args []BmDataStorage.BmStorage) *NtmBusinessinputResource {
	var bis *NtmDataStorage.NtmBusinessinputStorage
	var pis *NtmDataStorage.NtmPaperinputStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmBusinessinputStorage" {
			bis = arg.(*NtmDataStorage.NtmBusinessinputStorage)
		} else if tp.Name() == "NtmPaperinputStorage" {
			pis = arg.(*NtmDataStorage.NtmPaperinputStorage)
		}
	}
	return &NtmBusinessinputResource{
		NtmBusinessinputStorage: bis,
		NtmPaperinputStorage:    pis,
	}
}

func (s NtmBusinessinputResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	paperinputsID, piok := r.QueryParams["paperinputsID"]
	var result []*NtmModel.Businessinput

	if piok {
		modelRootID := paperinputsID[0]

		modelRoot, err := s.NtmPaperinputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		r.QueryParams["ids"] = modelRoot.BusinessinputIDs


		result = s.NtmBusinessinputStorage.GetAll(r, -1, -1)

		return &Response{Res: result}, nil
	}

	models := s.NtmBusinessinputStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, model)
	}
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmBusinessinputResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.Businessinput
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
		for _, iter := range s.NtmBusinessinputStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmBusinessinputStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.Businessinput{}
	count := s.NtmBusinessinputStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmBusinessinputResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmBusinessinputStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmBusinessinputResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Businessinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmBusinessinputStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmBusinessinputResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmBusinessinputStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmBusinessinputResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Businessinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmBusinessinputStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
