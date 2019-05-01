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

type NtmPolicyResource struct {
	NtmPolicyStorage *NtmDataStorage.NtmPolicyStorage
	NtmHospitalConfigStorage *NtmDataStorage.NtmHospitalConfigStorage
}

func (c NtmPolicyResource) NewPolicyResource(args []BmDataStorage.BmStorage) *NtmPolicyResource {
	var cs *NtmDataStorage.NtmPolicyStorage
	var hcs *NtmDataStorage.NtmHospitalConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmPolicyStorage" {
			cs = arg.(*NtmDataStorage.NtmPolicyStorage)
	 	} else if tp.Name() == "NtmHospitalConfigStorage" {
	 		hcs = arg.(*NtmDataStorage.NtmHospitalConfigStorage)
		}
	}
	return &NtmPolicyResource{
		NtmPolicyStorage: cs,
		NtmHospitalConfigStorage: hcs,
	}
}

// FindAll Policys
func (c NtmPolicyResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*NtmModel.Policy
	hospitalConfigsID, pciok := r.QueryParams["hospitalConfigsID"]

	if pciok {
		modelRootID := hospitalConfigsID[0]

		modelRoot, err := c.NtmHospitalConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		r.QueryParams["ids"] = modelRoot.PolicyIDs

		result = c.NtmPolicyStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	result = c.NtmPolicyStorage.GetAll(r, -1, -1)
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
