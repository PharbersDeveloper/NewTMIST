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

type NtmManagerinputResource struct {
	NtmManagerinputStorage *NtmDataStorage.NtmManagerinputStorage
	NtmPaperinputStorage   *NtmDataStorage.NtmPaperinputStorage
}

func (s NtmManagerinputResource) NewManagerinputResource(args []BmDataStorage.BmStorage) *NtmManagerinputResource {
	var bis *NtmDataStorage.NtmManagerinputStorage
	var pis *NtmDataStorage.NtmPaperinputStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmManagerinputStorage" {
			bis = arg.(*NtmDataStorage.NtmManagerinputStorage)
		} else if tp.Name() == "NtmPaperinputStorage" {
			pis = arg.(*NtmDataStorage.NtmPaperinputStorage)
		}
	}
	return &NtmManagerinputResource{
		NtmManagerinputStorage: bis,
		NtmPaperinputStorage:   pis,
	}
}

func (s NtmManagerinputResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	paperinputsID, piok := r.QueryParams["paperinputsID"]
	var result []*NtmModel.Managerinput

	if piok {
		modelRootID := paperinputsID[0]

		modelRoot, err := s.NtmPaperinputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}

		r.QueryParams["ids"] = modelRoot.ManagerinputIDs
		result = s.NtmManagerinputStorage.GetAll(r, -1, -1)
		return &Response{Res: result}, nil
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmManagerinputResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.Managerinput
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
		for _, iter := range s.NtmManagerinputStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmManagerinputStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.Managerinput{}
	count := s.NtmManagerinputStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmManagerinputResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmManagerinputStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmManagerinputResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Managerinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmManagerinputStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmManagerinputResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmManagerinputStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmManagerinputResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Managerinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmManagerinputStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
