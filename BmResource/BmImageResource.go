package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/NtmPods/BmDataStorage"
	"github.com/PharbersDeveloper/NtmPods/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmImageResource struct {
	BmImageStorage       *BmDataStorage.BmImageStorage
	BmSessioninfoStorage *BmDataStorage.BmSessioninfoStorage
	BmBrandStorage       *BmDataStorage.BmBrandStorage
	BmYardStorage        *BmDataStorage.BmYardStorage
}

func (c BmImageResource) NewImageResource(args []BmDataStorage.BmStorage) BmImageResource {
	var us *BmDataStorage.BmSessioninfoStorage
	var cs *BmDataStorage.BmImageStorage
	var bs *BmDataStorage.BmBrandStorage
	var ys *BmDataStorage.BmYardStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmSessioninfoStorage" {
			us = arg.(*BmDataStorage.BmSessioninfoStorage)
		} else if tp.Name() == "BmImageStorage" {
			cs = arg.(*BmDataStorage.BmImageStorage)
		} else if tp.Name() == "BmBrandStorage" {
			bs = arg.(*BmDataStorage.BmBrandStorage)
		} else if tp.Name() == "BmYardStorage" {
			ys = arg.(*BmDataStorage.BmYardStorage)
		}
	}
	return BmImageResource{BmSessioninfoStorage: us, BmImageStorage: cs, BmBrandStorage: bs, BmYardStorage: ys}
}

// FindAll images
func (c BmImageResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	sessioninfosID, ok := r.QueryParams["sessioninfosID"]
	brandsID, brdok := r.QueryParams["brandsID"]
	yardsID, ydok := r.QueryParams["yardsID"]
	result := []BmModel.Image{}
	if ok {
		modelRootID := sessioninfosID[0]

		modelRoot, err := c.BmSessioninfoStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		for _, modelID := range modelRoot.ImagesIDs {
			model, err := c.BmImageStorage.GetOne(modelID)
			if err != nil {
				return &Response{}, err
			}
			result = append(result, model)
		}

		return &Response{Res: result}, nil
	} else if brdok {
		modelRootID := brandsID[0]

		modelRoot, err := c.BmBrandStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		for _, modelID := range modelRoot.ImagesIDs {
			model, err := c.BmImageStorage.GetOne(modelID)
			if err != nil {
				return &Response{}, err
			}
			result = append(result, model)
		}

		return &Response{Res: result}, nil
	} else if ydok {
		modelRootID := yardsID[0]

		modelRoot, err := c.BmYardStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		for _, modelID := range modelRoot.ImagesIDs {
			model, err := c.BmImageStorage.GetOne(modelID)
			if err != nil {
				return &Response{}, err
			}
			result = append(result, model)
		}

		return &Response{Res: result}, nil
	}
	//result = c.BmImageStorage.GetAll(r)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmImageResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmImageStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmImageResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Image)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmImageStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmImageResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmImageStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmImageResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Image)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmImageStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
