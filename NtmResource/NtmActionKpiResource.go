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

type NtmActionKpiResource struct {
	NtmActionKpiStorage 			*NtmDataStorage.NtmActionKpiStorage
	//NtmRepresentativeConfigStorage	*NtmDataStorage.NtmRepresentativeConfigStorage
	//NtmPaperStorage					*NtmDataStorage.NtmPaperStorage
}

func (c NtmActionKpiResource) NewActionKpiResource(args []BmDataStorage.BmStorage) *NtmActionKpiResource {
	var cs *NtmDataStorage.NtmActionKpiStorage
	//var rcs *NtmDataStorage.NtmRepresentativeConfigStorage
	//var ps *NtmDataStorage.NtmPaperStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmActionKpiStorage" {
			cs = arg.(*NtmDataStorage.NtmActionKpiStorage)
		}
		//else if tp.Name() == "NtmRepresentativeConfigStorage" {
		//	rcs = arg.(*NtmDataStorage.NtmRepresentativeConfigStorage)
		//} else if tp.Name() == "NtmPaperStorage" {
		//	ps = arg.(*NtmDataStorage.NtmPaperStorage)
		//}
	}
	return &NtmActionKpiResource{
		NtmActionKpiStorage: cs,
	}
}

// FindAll ActionKpis
func (c NtmActionKpiResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*NtmModel.ActionKpi
	//papersID, rciok := r.QueryParams["papersID"]
	//
	//if rciok {
	//	modelRootID := papersID[0]
	//
	//	modelRoot, err := c.NtmPaperStorage.GetOne(modelRootID)
	//	if err != nil {
	//		return &Response{}, err
	//	}
	//
	//	r.QueryParams["ids"] = modelRoot.ActionKpiIDs
	//
	//	result = c.NtmActionKpiStorage.GetAll(r, -1,-1)
	//
	//	return &Response{Res: result}, nil
	//}

	result = c.NtmActionKpiStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmActionKpiResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmActionKpiStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmActionKpiResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ActionKpi)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmActionKpiStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmActionKpiResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmActionKpiStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmActionKpiResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ActionKpi)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmActionKpiStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
