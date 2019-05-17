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

type NtmAssessmentReportDescribeResource struct {
	NtmAssessmentReportDescribeStorage          *NtmDataStorage.NtmAssessmentReportDescribeStorage
	NtmRegionalDivisionResultStorage	*NtmDataStorage.NtmRegionalDivisionResultStorage
	NtmTargetAssignsResultStorage		*NtmDataStorage.NtmTargetAssignsResultStorage
	NtmResourceAssignsResultStorage		*NtmDataStorage.NtmResourceAssignsResultStorage
	NtmManageTimeResultStorage			*NtmDataStorage.NtmManageTimeResultStorage
	NtmManageTeamResultStorage			*NtmDataStorage.NtmManageTeamResultStorage

}

func (c NtmAssessmentReportDescribeResource) NewAssessmentReportDescribeResource(args []BmDataStorage.BmStorage) *NtmAssessmentReportDescribeResource {
	var ard *NtmDataStorage.NtmAssessmentReportDescribeStorage
	var rdr	*NtmDataStorage.NtmRegionalDivisionResultStorage
	var tar	*NtmDataStorage.NtmTargetAssignsResultStorage
	var rar *NtmDataStorage.NtmResourceAssignsResultStorage
	var mtr *NtmDataStorage.NtmManageTimeResultStorage
	var mtrs *NtmDataStorage.NtmManageTeamResultStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmAssessmentReportDescribeStorage" {
			ard = arg.(*NtmDataStorage.NtmAssessmentReportDescribeStorage)
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
		}
	}
	return &NtmAssessmentReportDescribeResource{
		NtmAssessmentReportDescribeStorage: ard,
		NtmRegionalDivisionResultStorage: rdr,
		NtmTargetAssignsResultStorage: tar,
		NtmResourceAssignsResultStorage: rar,
		NtmManageTimeResultStorage: mtr,
		NtmManageTeamResultStorage: mtrs,
	}
}

// FindAll images
func (c NtmAssessmentReportDescribeResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.AssessmentReportDescribe

	regionalDivisionResultsID, rdrOk := r.QueryParams["regionalDivisionResultsID"]
	targetAssignsResultsID, tarOk := r.QueryParams["targetAssignsResultsID"]
	resourceAssignsResultsID, rarOk := r.QueryParams["resourceAssignsResultsID"]
	manageTimeResultsID, mtrOk := r.QueryParams["manageTimeResultsID"]
	manageTeamResultsID, mtrsOk := r.QueryParams["manageTeamResultsID"]

	if rdrOk {
		modelRootID := regionalDivisionResultsID[0]
		modelRoot, err := c.NtmRegionalDivisionResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportDescribeIDs
		model := c.NtmAssessmentReportDescribeStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}

	if tarOk {
		modelRootID := targetAssignsResultsID[0]
		modelRoot, err := c.NtmTargetAssignsResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportDescribeIDs
		model := c.NtmAssessmentReportDescribeStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}

	if rarOk {
		modelRootID := resourceAssignsResultsID[0]
		modelRoot, err := c.NtmResourceAssignsResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportDescribeIDs
		model := c.NtmAssessmentReportDescribeStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}

	if mtrOk {
		modelRootID := manageTimeResultsID[0]
		modelRoot, err := c.NtmManageTimeResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportDescribeIDs
		model := c.NtmAssessmentReportDescribeStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}

	if mtrsOk {
		modelRootID := manageTeamResultsID[0]
		modelRoot, err := c.NtmManageTeamResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportDescribeIDs
		model := c.NtmAssessmentReportDescribeStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}


	result = c.NtmAssessmentReportDescribeStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmAssessmentReportDescribeResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmAssessmentReportDescribeStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmAssessmentReportDescribeResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.AssessmentReportDescribe)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmAssessmentReportDescribeStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmAssessmentReportDescribeResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmAssessmentReportDescribeStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmAssessmentReportDescribeResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.AssessmentReportDescribe)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmAssessmentReportDescribeStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
