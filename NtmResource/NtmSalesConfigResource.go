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

type NtmSalesConfigResource struct {
	NtmSalesConfigStorage       *NtmDataStorage.NtmSalesConfigStorage
}

func (c NtmSalesConfigResource) NewSalesConfigResource(args []BmDataStorage.BmStorage) *NtmSalesConfigResource {
	var sc *NtmDataStorage.NtmSalesConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmSalesConfigStorage" {
			sc = arg.(*NtmDataStorage.NtmSalesConfigStorage)
		}
	}
	return &NtmSalesConfigResource{
		NtmSalesConfigStorage:	sc,
	}
}

// FindAll SalesConfigs
func (c NtmSalesConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.SalesConfig
	result = c.NtmSalesConfigStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmSalesConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmSalesConfigStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmSalesConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.SalesConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmSalesConfigStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmSalesConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmSalesConfigStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmSalesConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.SalesConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmSalesConfigStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
