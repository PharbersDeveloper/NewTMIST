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

type NtmManagerConfigResource struct {
	NtmManagerConfigStorage *NtmDataStorage.NtmManagerConfigStorage
}

func (c NtmManagerConfigResource) NewManagerConfigResource(args []BmDataStorage.BmStorage) NtmManagerConfigResource {
	var cs *NtmDataStorage.NtmManagerConfigStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmManagerConfigStorage" {
			cs = arg.(*NtmDataStorage.NtmManagerConfigStorage)
		}
	}
	return NtmManagerConfigResource{NtmManagerConfigStorage: cs}
}

// FindAll images
func (c NtmManagerConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := []NtmModel.ManagerConfig{}
	result = c.NtmManagerConfigStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmManagerConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmManagerConfigStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmManagerConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ManagerConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmManagerConfigStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmManagerConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmManagerConfigStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmManagerConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ManagerConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmManagerConfigStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
