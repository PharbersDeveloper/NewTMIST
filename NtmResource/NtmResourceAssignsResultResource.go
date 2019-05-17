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

type NtmResourceAssignsResultResource struct {
	NtmResourceAssignsResultStorage	*NtmDataStorage.NtmResourceAssignsResultStorage
	NtmAssessmentReportStorage			*NtmDataStorage.NtmAssessmentReportStorage
}

func (c NtmResourceAssignsResultResource) NewResourceAssignsResultResource(args []BmDataStorage.BmStorage) *NtmResourceAssignsResultResource {
	var rdr *NtmDataStorage.NtmResourceAssignsResultStorage
	var ar *NtmDataStorage.NtmAssessmentReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmResourceAssignsResultStorage" {
			rdr = arg.(*NtmDataStorage.NtmResourceAssignsResultStorage)
		} else if tp.Name() == "NtmAssessmentReportStorage" {
			ar = arg.(*NtmDataStorage.NtmAssessmentReportStorage)
		}
	}
	return &NtmResourceAssignsResultResource{
		NtmResourceAssignsResultStorage: rdr,
		NtmAssessmentReportStorage: ar,
	}
}

// FindAll images
func (c NtmResourceAssignsResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	assessmentReportsID, arOk := r.QueryParams["assessmentReportsID"]

	if arOk {
		modelRootID := assessmentReportsID[0]
		modelRoot, err := c.NtmAssessmentReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmResourceAssignsResultStorage.GetOne(modelRoot.ResourceAssignsResultID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []NtmModel.ResourceAssignsResult
	result = c.NtmResourceAssignsResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmResourceAssignsResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmResourceAssignsResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmResourceAssignsResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ResourceAssignsResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmResourceAssignsResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmResourceAssignsResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmResourceAssignsResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmResourceAssignsResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ResourceAssignsResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmResourceAssignsResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
