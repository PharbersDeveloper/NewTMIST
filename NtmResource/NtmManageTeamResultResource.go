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

type NtmManageTeamResultResource struct {
	NtmManageTeamResultStorage	*NtmDataStorage.NtmManageTeamResultStorage
	NtmAssessmentReportStorage			*NtmDataStorage.NtmAssessmentReportStorage
}

func (c NtmManageTeamResultResource) NewManageTeamResultResource(args []BmDataStorage.BmStorage) *NtmManageTeamResultResource {
	var rdr *NtmDataStorage.NtmManageTeamResultStorage
	var ar *NtmDataStorage.NtmAssessmentReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmManageTeamResultStorage" {
			rdr = arg.(*NtmDataStorage.NtmManageTeamResultStorage)
		} else if tp.Name() == "NtmAssessmentReportStorage" {
			ar = arg.(*NtmDataStorage.NtmAssessmentReportStorage)
		}
	}
	return &NtmManageTeamResultResource{
		NtmManageTeamResultStorage: rdr,
		NtmAssessmentReportStorage: ar,
	}
}

// FindAll images
func (c NtmManageTeamResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	assessmentReportsID, arOk := r.QueryParams["assessmentReportsID"]

	if arOk {
		modelRootID := assessmentReportsID[0]
		modelRoot, err := c.NtmAssessmentReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmManageTeamResultStorage.GetOne(modelRoot.ManageTeamResultID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []NtmModel.ManageTeamResult
	result = c.NtmManageTeamResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmManageTeamResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmManageTeamResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmManageTeamResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ManageTeamResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmManageTeamResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmManageTeamResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmManageTeamResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmManageTeamResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.ManageTeamResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmManageTeamResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
