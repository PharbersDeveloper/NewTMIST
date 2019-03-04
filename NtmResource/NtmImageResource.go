package NtmResource

import (
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmDataStorage"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"reflect"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type NtmImageResource struct {
	NtmImageStorage *NtmDataStorage.NtmImageStorage
	NtmProductStorage *NtmDataStorage.NtmProductStorage
}

func (c NtmImageResource) NewImageResource(args []BmDataStorage.BmStorage) NtmImageResource {
	var cs *NtmDataStorage.NtmImageStorage
	var ps *NtmDataStorage.NtmProductStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmImageStorage" {
			cs = arg.(*NtmDataStorage.NtmImageStorage)
		} else if tp.Name() == "NtmProductStorage" {
			ps = arg.(*NtmDataStorage.NtmProductStorage)
		}
	}
	return NtmImageResource{
		NtmImageStorage: cs,
		NtmProductStorage: ps,
	}
}

// FindAll images
func (c NtmImageResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	productsID, pok := r.QueryParams["productsID"]
	result := []NtmModel.Image{}
	if pok {
		modelRootID := productsID[0]

		modelRoot, err := c.NtmProductStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		for _, modelID := range modelRoot.ImagesIDs {
			model, err := c.NtmImageStorage.GetOne(modelID)
			if err != nil {
				return &Response{}, err
			}
			result = append(result, model)
		}

		return &Response{Res: result}, nil
	}
	result = c.NtmImageStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmImageResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmImageStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmImageResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Image)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmImageStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmImageResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmImageStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmImageResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Image)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmImageStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
