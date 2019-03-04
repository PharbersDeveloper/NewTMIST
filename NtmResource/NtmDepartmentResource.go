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

type NtmDepartmentResource struct {
	NtmDepartmentStorage *NtmDataStorage.NtmDepartmentStorage
}

func (c NtmDepartmentResource) NewDepartmentResource(args []BmDataStorage.BmStorage) *NtmDepartmentResource {
	var cs *NtmDataStorage.NtmDepartmentStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmDepartmentStorage" {
			cs = arg.(*NtmDataStorage.NtmDepartmentStorage)
		}
	}
	return &NtmDepartmentResource{NtmDepartmentStorage: cs}
}

// FindAll Departments
func (c NtmDepartmentResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := c.NtmDepartmentStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmDepartmentResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmDepartmentStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmDepartmentResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Department)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmDepartmentStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmDepartmentResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmDepartmentStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmDepartmentResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Department)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmDepartmentStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
