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

type NtmProductResource struct {
	NtmProductStorage       *NtmDataStorage.NtmProductStorage
	NtmImageStorage         *NtmDataStorage.NtmImageStorage
	NtmProductConfigStorage *NtmDataStorage.NtmProductConfigStorage
}

func (s NtmProductResource) NewProductResource(args []BmDataStorage.BmStorage) *NtmProductResource {
	var is *NtmDataStorage.NtmImageStorage
	var ps *NtmDataStorage.NtmProductStorage
	var pcs *NtmDataStorage.NtmProductConfigStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmImageStorage" {
			is = arg.(*NtmDataStorage.NtmImageStorage)
		} else if tp.Name() == "NtmProductStorage" {
			ps = arg.(*NtmDataStorage.NtmProductStorage)
		} else if tp.Name() == "NtmProductConfigStorage" {
			pcs = arg.(*NtmDataStorage.NtmProductConfigStorage)
		}
	}
	return &NtmProductResource{
		NtmImageStorage:         is,
		NtmProductStorage:       ps,
		NtmProductConfigStorage: pcs,
	}
}

func (s NtmProductResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	productConfigID, pcok := r.QueryParams["productConfigsID"]
	var result []NtmModel.Product

	if pcok {
		modelRootID := productConfigID[0]

		modelRoot, err := s.NtmProductConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}

		model, err := s.NtmProductStorage.GetOne(modelRoot.ProductID)
		if err != nil {
			return &Response{}, err
		}
		result = append(result, model)

		return &Response{Res: result}, nil
	}

	models := s.NtmProductStorage.GetAll(r, -1, -1)
	for _, model := range models {
		// get all sweets for the model
		model.Imgs = []*NtmModel.Image{}
		for _, kID := range model.ImagesIDs {
			choc, err := s.NtmImageStorage.GetOne(kID)
			if err != nil {
				return &Response{}, err
			}
			model.Imgs = append(model.Imgs, &choc)
		}

		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmProductResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.Product
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
		for _, iter := range s.NtmProductStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmProductStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.Product{}
	count := s.NtmProductStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmProductResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmProductStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	model.Imgs = []*NtmModel.Image{}
	for _, kID := range model.ImagesIDs {
		choc, err := s.NtmImageStorage.GetOne(kID)
		if err != nil {
			return &Response{}, err
		}
		model.Imgs = append(model.Imgs, &choc)
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmProductResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Product)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmProductStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmProductResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmProductStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmProductResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Product)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmProductStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
