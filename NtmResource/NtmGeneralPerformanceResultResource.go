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

type NtmGeneralPerformanceResultResource struct {
	NtmGeneralPerformanceResultStorage	*NtmDataStorage.NtmGeneralPerformanceResultStorage
	NtmAssessmentReportStorage			*NtmDataStorage.NtmAssessmentReportStorage
}

func (c NtmGeneralPerformanceResultResource) NewGeneralPerformanceResultResource(args []BmDataStorage.BmStorage) *NtmGeneralPerformanceResultResource {
	var rdr *NtmDataStorage.NtmGeneralPerformanceResultStorage
	var ar *NtmDataStorage.NtmAssessmentReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmGeneralPerformanceResultStorage" {
			rdr = arg.(*NtmDataStorage.NtmGeneralPerformanceResultStorage)
		} else if tp.Name() == "NtmAssessmentReportStorage" {
			ar = arg.(*NtmDataStorage.NtmAssessmentReportStorage)
		}
	}
	return &NtmGeneralPerformanceResultResource{
		NtmGeneralPerformanceResultStorage: rdr,
		NtmAssessmentReportStorage: ar,
	}
}

// FindAll images
func (c NtmGeneralPerformanceResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	assessmentReportsID, arOk := r.QueryParams["assessmentReportsID"]

	if arOk {
		modelRootID := assessmentReportsID[0]
		modelRoot, err := c.NtmAssessmentReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.NtmGeneralPerformanceResultStorage.GetOne(modelRoot.GeneralPerformanceResultID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []NtmModel.GeneralPerformanceResult
	result = c.NtmGeneralPerformanceResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmGeneralPerformanceResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmGeneralPerformanceResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmGeneralPerformanceResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.GeneralPerformanceResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmGeneralPerformanceResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmGeneralPerformanceResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmGeneralPerformanceResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmGeneralPerformanceResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.GeneralPerformanceResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmGeneralPerformanceResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
