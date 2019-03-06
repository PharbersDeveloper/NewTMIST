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

type NtmScenarioResource struct {
	NtmScenarioStorage *NtmDataStorage.NtmScenarioStorage
}

func (c NtmScenarioResource) NewScenarioResource(args []BmDataStorage.BmStorage) *NtmScenarioResource {
	var cs *NtmDataStorage.NtmScenarioStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmScenarioStorage" {
			cs = arg.(*NtmDataStorage.NtmScenarioStorage)
		}
	}
	return &NtmScenarioResource{NtmScenarioStorage: cs}
}

// FindAll Scenarios
func (c NtmScenarioResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := c.NtmScenarioStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmScenarioResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmScenarioStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmScenarioResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Scenario)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmScenarioStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmScenarioResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmScenarioStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmScenarioResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Scenario)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmScenarioStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
