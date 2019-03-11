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

type NtmRegionResource struct {
	NtmRegionStorage *NtmDataStorage.NtmRegionStorage
	NtmImageStorage  *NtmDataStorage.NtmImageStorage
	NtmRegionConfigStorage *NtmDataStorage.NtmRegionConfigStorage
}

func (s NtmRegionResource) NewRegionResource(args []BmDataStorage.BmStorage) *NtmRegionResource {
	var is *NtmDataStorage.NtmImageStorage
	var hs *NtmDataStorage.NtmRegionStorage
	var rcs *NtmDataStorage.NtmRegionConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmImageStorage" {
			is = arg.(*NtmDataStorage.NtmImageStorage)
		} else if tp.Name() == "NtmRegionStorage" {
			hs = arg.(*NtmDataStorage.NtmRegionStorage)
		} else if tp.Name() == "NtmRegionConfigStorage" {
			rcs = arg.(*NtmDataStorage.NtmRegionConfigStorage)
		}
	}
	return &NtmRegionResource{
		NtmImageStorage: is,
		NtmRegionStorage: hs,
		NtmRegionConfigStorage: rcs,
	}
}

func (s NtmRegionResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	regionConfigsID, pciok := r.QueryParams["regionConfigsID"]

	if pciok {
		modelRootID := regionConfigsID[0]
		modelRoot, err := s.NtmRegionConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err := s.NtmRegionConfigStorage.GetOne(modelRoot.RegionID)
		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []NtmModel.Region
	models := s.NtmRegionStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmRegionResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.Region
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
		for _, iter := range s.NtmRegionStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmRegionStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.Region{}
	count := s.NtmRegionStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmRegionResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmRegionStorage.GetOne(ID)
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
func (s NtmRegionResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Region)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmRegionStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmRegionResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmRegionStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmRegionResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Region)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmRegionStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
