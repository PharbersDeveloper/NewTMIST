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

type NtmProposalResource struct {
	NtmProposalStorage        *NtmDataStorage.NtmProposalStorage
	NtmUseableProposalStorage *NtmDataStorage.NtmUseableProposalStorage
}

func (c NtmProposalResource) NewProposalResource(args []BmDataStorage.BmStorage) *NtmProposalResource {
	var ps *NtmDataStorage.NtmProposalStorage
	var ups *NtmDataStorage.NtmUseableProposalStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmProposalStorage" {
			ps = arg.(*NtmDataStorage.NtmProposalStorage)
		} else if tp.Name() == "NtmUseableProposalStorage" {
			ups = arg.(*NtmDataStorage.NtmUseableProposalStorage)
		}
	}
	return &NtmProposalResource{
		NtmProposalStorage:        ps,
		NtmUseableProposalStorage: ups,
	}
}

// FindAll Proposals
func (c NtmProposalResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	useableProposalsID, upiok := r.QueryParams["useableProposalsID"]

	if upiok {
		modelRootID := useableProposalsID[0]
		modelRoot, err := c.NtmUseableProposalStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := c.NtmProposalStorage.GetOne(modelRoot.ProposalID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	result := c.NtmProposalStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmProposalResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmProposalStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmProposalResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Proposal)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmProposalStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmProposalResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmProposalStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmProposalResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Proposal)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmProposalStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
