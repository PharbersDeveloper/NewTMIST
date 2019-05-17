package NtmResource

import (
	"errors"
	"Ntm/NtmDataStorage"
	"Ntm/NtmModel"
	"reflect"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type NtmTitleResource struct {
	NtmTitleStorage	*NtmDataStorage.NtmTitleStorage
	NtmLevelConfigStorage 		*NtmDataStorage.NtmLevelConfigStorage
}

func (c NtmTitleResource) NewTitleResource(args []BmDataStorage.BmStorage) *NtmTitleResource {
	var rdr *NtmDataStorage.NtmTitleStorage
	var lcs *NtmDataStorage.NtmLevelConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmTitleStorage" {
			rdr = arg.(*NtmDataStorage.NtmTitleStorage)
		} else if tp.Name() == "NtmLevelConfigStorage" {
			lcs = arg.(*NtmDataStorage.NtmLevelConfigStorage)
		}
	}
	return &NtmTitleResource{
		NtmTitleStorage: rdr,
		NtmLevelConfigStorage: lcs,
	}
}

// FindAll images
func (c NtmTitleResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	levelConfigsID, lcsOk := r.QueryParams["levelConfigsID"]

	if lcsOk {
		modelRootID := levelConfigsID[0]
		modelRoot, err := c.NtmLevelConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmTitleStorage.GetOne(modelRoot.TitleID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []NtmModel.Title
	result = c.NtmTitleStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmTitleResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmTitleStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmTitleResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Title)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmTitleStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmTitleResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmTitleStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmTitleResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Title)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmTitleStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
