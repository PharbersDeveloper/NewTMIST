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

type NtmAssessmentReportResource struct {
	NtmAssessmentReportStorage          *NtmDataStorage.NtmAssessmentReportStorage
	NtmRegionalDivisionResultStorage	*NtmDataStorage.NtmRegionalDivisionResultStorage
	NtmTargetAssignsResultStorage		*NtmDataStorage.NtmTargetAssignsResultStorage
	NtmResourceAssignsResultStorage		*NtmDataStorage.NtmResourceAssignsResultStorage
	NtmManageTimeResultStorage			*NtmDataStorage.NtmManageTimeResultStorage
	NtmManageTeamResultStorage			*NtmDataStorage.NtmManageTeamResultStorage
	NtmGeneralPerformanceResultStorage	*NtmDataStorage.NtmGeneralPerformanceResultStorage
	NtmPaperStorage 					*NtmDataStorage.NtmPaperStorage

}

func (c NtmAssessmentReportResource) NewAssessmentReportResource(args []BmDataStorage.BmStorage) *NtmAssessmentReportResource {
	var ard *NtmDataStorage.NtmAssessmentReportStorage
	var rdr	*NtmDataStorage.NtmRegionalDivisionResultStorage
	var tar	*NtmDataStorage.NtmTargetAssignsResultStorage
	var rar *NtmDataStorage.NtmResourceAssignsResultStorage
	var mtr *NtmDataStorage.NtmManageTimeResultStorage
	var mtrs *NtmDataStorage.NtmManageTeamResultStorage
	var gpr *NtmDataStorage.NtmGeneralPerformanceResultStorage
	var p *NtmDataStorage.NtmPaperStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmAssessmentReportStorage" {
			ard = arg.(*NtmDataStorage.NtmAssessmentReportStorage)
		} else if tp.Name() == "NtmRegionalDivisionResultStorage" {
			rdr = arg.(*NtmDataStorage.NtmRegionalDivisionResultStorage)
		} else if tp.Name() == "NtmTargetAssignsResultStorage" {
			tar = arg.(*NtmDataStorage.NtmTargetAssignsResultStorage)
		} else if tp.Name() == "NtmResourceAssignsResultStorage" {
			rar = arg.(*NtmDataStorage.NtmResourceAssignsResultStorage)
		} else if tp.Name() == "NtmManageTimeResultStorage" {
			mtr = arg.(*NtmDataStorage.NtmManageTimeResultStorage)
		} else if tp.Name() == "NtmManageTeamResultStorage" {
			mtrs = arg.(*NtmDataStorage.NtmManageTeamResultStorage)
		} else if tp.Name() == "NtmPaperStorage" {
			p = arg.(*NtmDataStorage.NtmPaperStorage)
		} else if tp.Name() == "NtmGeneralPerformanceResultStorage" {
			gpr = arg.(*NtmDataStorage.NtmGeneralPerformanceResultStorage)
		}
	}
	return &NtmAssessmentReportResource{
		NtmAssessmentReportStorage: ard,
		NtmRegionalDivisionResultStorage: rdr,
		NtmTargetAssignsResultStorage: tar,
		NtmResourceAssignsResultStorage: rar,
		NtmManageTimeResultStorage: mtr,
		NtmManageTeamResultStorage: mtrs,
		NtmPaperStorage: p,
		NtmGeneralPerformanceResultStorage: gpr,
	}
}

// FindAll images
func (c NtmAssessmentReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.AssessmentReport

	papersID, pOk := r.QueryParams["papersID"]

	if pOk {
		modelRootID := papersID[0]
		modelRoot, err := c.NtmPaperStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportIDs

		models := c.NtmAssessmentReportStorage.GetAll(r, -1,-1)

		return &Response{Res: models}, nil
	}

	result = c.NtmAssessmentReportStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmAssessmentReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmAssessmentReportStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmAssessmentReportResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.AssessmentReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmAssessmentReportStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmAssessmentReportResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmAssessmentReportStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmAssessmentReportResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.AssessmentReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmAssessmentReportStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
