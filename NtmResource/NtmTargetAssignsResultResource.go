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

type NtmTargetAssignsResultResource struct {
	NtmTargetAssignsResultStorage	*NtmDataStorage.NtmTargetAssignsResultStorage
	NtmAssessmentReportStorage			*NtmDataStorage.NtmAssessmentReportStorage
}

func (c NtmTargetAssignsResultResource) NewTargetAssignsResultResource(args []BmDataStorage.BmStorage) *NtmTargetAssignsResultResource {
	var rdr *NtmDataStorage.NtmTargetAssignsResultStorage
	var ar *NtmDataStorage.NtmAssessmentReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmTargetAssignsResultStorage" {
			rdr = arg.(*NtmDataStorage.NtmTargetAssignsResultStorage)
		} else if tp.Name() == "NtmAssessmentReportStorage" {
			ar = arg.(*NtmDataStorage.NtmAssessmentReportStorage)
		}
	}
	return &NtmTargetAssignsResultResource{
		NtmTargetAssignsResultStorage: rdr,
		NtmAssessmentReportStorage: ar,
	}
}

// FindAll images
func (c NtmTargetAssignsResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	assessmentReportsID, arOk := r.QueryParams["assessmentReportsID"]

	if arOk {
		modelRootID := assessmentReportsID[0]
		modelRoot, err := c.NtmAssessmentReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmTargetAssignsResultStorage.GetOne(modelRoot.TargetAssignsResultID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []NtmModel.TargetAssignsResult
	result = c.NtmTargetAssignsResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmTargetAssignsResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmTargetAssignsResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmTargetAssignsResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.TargetAssignsResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmTargetAssignsResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmTargetAssignsResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmTargetAssignsResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmTargetAssignsResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.TargetAssignsResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmTargetAssignsResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
