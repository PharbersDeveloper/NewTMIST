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

type NtmManageTimeResultResource struct {
	NtmManageTimeResultStorage	*NtmDataStorage.NtmManageTimeResultStorage
	NtmAssessmentReportStorage			*NtmDataStorage.NtmAssessmentReportStorage
}

func (c NtmManageTimeResultResource) NewManageTimeResultResource(args []BmDataStorage.BmStorage) *NtmManageTimeResultResource {
	var rdr *NtmDataStorage.NtmManageTimeResultStorage
	var ar *NtmDataStorage.NtmAssessmentReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmManageTimeResultStorage" {
			rdr = arg.(*NtmDataStorage.NtmManageTimeResultStorage)
		} else if tp.Name() == "NtmAssessmentReportStorage" {
			ar = arg.(*NtmDataStorage.NtmAssessmentReportStorage)
		}
	}
	return &NtmManageTimeResultResource{
		NtmManageTimeResultStorage: rdr,
		NtmAssessmentReportStorage: ar,
	}
}

// FindAll images
func (c NtmManageTimeResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	assessmentReportsID, arOk := r.QueryParams["assessmentReportsID"]

	if arOk {
		modelRootID := assessmentReportsID[0]
		modelRoot, err := c.NtmAssessmentReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmManageTimeResultStorage.GetOne(modelRoot.ManageTimeResultID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []NtmModel.ManageTimeResult
	result = c.NtmManageTimeResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmManageTimeResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmManageTimeResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmManageTimeResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ManageTimeResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmManageTimeResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmManageTimeResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmManageTimeResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmManageTimeResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ManageTimeResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmManageTimeResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
