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

type NtmPolicyResource struct {
	NtmPolicyStorage *NtmDataStorage.NtmPolicyStorage
}

func (c NtmPolicyResource) NewPolicyResource(args []BmDataStorage.BmStorage) NtmPolicyResource {
	var cs *NtmDataStorage.NtmPolicyStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmPolicyStorage" {
			cs = arg.(*NtmDataStorage.NtmPolicyStorage)
		}
	}
	return NtmPolicyResource{NtmPolicyStorage: cs}
}

// FindAll Policys
func (c NtmPolicyResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := c.NtmPolicyStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmPolicyResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmPolicyStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmPolicyResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Policy)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmPolicyStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmPolicyResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmPolicyStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmPolicyResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Policy)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmPolicyStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
