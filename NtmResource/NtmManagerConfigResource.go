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
	NtmManagerConfigStorage  *NtmDataStorage.NtmManagerConfigStorage
	NtmResourceConfigStorage *NtmDataStorage.NtmResourceConfigStorage
}

func (c NtmManagerConfigResource) NewManagerConfigResource(args []BmDataStorage.BmStorage) *NtmManagerConfigResource {
	var mcs *NtmDataStorage.NtmManagerConfigStorage
	var rcs *NtmDataStorage.NtmResourceConfigStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmManagerConfigStorage" {
			mcs = arg.(*NtmDataStorage.NtmManagerConfigStorage)
		} else if tp.Name() == "NtmResourceConfigStorage" {
			rcs = arg.(*NtmDataStorage.NtmResourceConfigStorage)
		}
	}
	return &NtmManagerConfigResource{
		NtmManagerConfigStorage:  mcs,
		NtmResourceConfigStorage: rcs,
	}
}

func (c NtmManagerConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	resourceConfigsID, rcok := r.QueryParams["resourceConfigsID"]
	result := []NtmModel.ManagerConfig{}
	if rcok {
		modelRootID := resourceConfigsID[0]

		modelRoot, err := c.NtmResourceConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}

		model, err := c.NtmManagerConfigStorage.GetOne(modelRoot.ResourceID)
		if err != nil {
			return &Response{}, err
		}
		result = append(result, model)

		return &Response{Res: result}, nil
	}
	result = c.NtmManagerConfigStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

func (c NtmManagerConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmManagerConfigStorage.GetOne(ID)
	return &Response{Res: res}, err
}

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

func (c NtmManagerConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmManagerConfigStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

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
