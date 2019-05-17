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

type NtmLevelResource struct {
	NtmLevelStorage          	*NtmDataStorage.NtmLevelStorage
	NtmLevelConfigStorage 		*NtmDataStorage.NtmLevelConfigStorage
}

func (c NtmLevelResource) NewLevelResource(args []BmDataStorage.BmStorage) *NtmLevelResource {
	var ls *NtmDataStorage.NtmLevelStorage
	var lcs *NtmDataStorage.NtmLevelConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmLevelStorage" {
			ls = arg.(*NtmDataStorage.NtmLevelStorage)
		} else if tp.Name() == "NtmLevelConfigStorage" {
			lcs = arg.(*NtmDataStorage.NtmLevelConfigStorage)
		}
	}
	return &NtmLevelResource{
		NtmLevelStorage: ls,
		NtmLevelConfigStorage: lcs,
	}
}

// FindAll images
func (c NtmLevelResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	levelConfigsID, lcsOk := r.QueryParams["levelConfigsID"]

	if lcsOk {
		modelRootID := levelConfigsID[0]
		modelRoot, err := c.NtmLevelConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmLevelStorage.GetOne(modelRoot.LevelID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []NtmModel.Level
	result = c.NtmLevelStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmLevelResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmLevelStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmLevelResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Level)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmLevelStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmLevelResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmLevelStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmLevelResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Level)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmLevelStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
