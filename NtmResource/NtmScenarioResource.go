package NtmResource

import (
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmDataStorage"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"net/http"
	"reflect"
	"strconv"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type NtmScenarioResource struct {
	NtmScenarioStorage 		*NtmDataStorage.NtmScenarioStorage
	NtmProposalStorage		*NtmDataStorage.NtmProposalStorage
	NtmPaperStorage			*NtmDataStorage.NtmPaperStorage
	NtmPaperinputStorage	*NtmDataStorage.NtmPaperinputStorage
	NtmPersonnelAssessmentStorage *NtmDataStorage.NtmPersonnelAssessmentStorage
}

func (c NtmScenarioResource) NewScenarioResource(args []BmDataStorage.BmStorage) *NtmScenarioResource {
	var cs *NtmDataStorage.NtmScenarioStorage
	var ps *NtmDataStorage.NtmProposalStorage
	var pas *NtmDataStorage.NtmPaperStorage
	var pis *NtmDataStorage.NtmPaperinputStorage
	var pass *NtmDataStorage.NtmPersonnelAssessmentStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmScenarioStorage" {
			cs = arg.(*NtmDataStorage.NtmScenarioStorage)
		} else if tp.Name() == "NtmProposalStorage" {
			ps = arg.(*NtmDataStorage.NtmProposalStorage)
		} else if tp.Name() == "NtmPaperStorage" {
			pas = arg.(*NtmDataStorage.NtmPaperStorage)
		} else if tp.Name() == "NtmPaperinputStorage" {
			pis = arg.(*NtmDataStorage.NtmPaperinputStorage)
		} else if tp.Name() == "NtmPersonnelAssessmentStorage" {
			pass = arg.(*NtmDataStorage.NtmPersonnelAssessmentStorage)
		}
	}
	return &NtmScenarioResource{
		NtmScenarioStorage: cs,
		NtmProposalStorage: ps,
		NtmPaperStorage: pas,
		NtmPaperinputStorage: pis,
		NtmPersonnelAssessmentStorage: pass,
	}
}

// FindAll Scenarios
// TODO @Alex 这边后续必须重构，太难看了 自己留
func (c NtmScenarioResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.Scenario
	proposalsID, psok := r.QueryParams["proposal-id"]
	_, acok := r.QueryParams["account-id"]

	paperinputsID, piok := r.QueryParams["paperinputsID"]
	personnelAssessmentsID, paok := r.QueryParams["personnelAssessmentsID"]


	if psok && acok {

		proposalModel, _ := c.NtmProposalStorage.GetOne(proposalsID[0])
		paperModel := c.NtmPaperStorage.GetAll(r, -1,-1)[0]
		r.QueryParams["ids"] = paperModel.InputIDs
		r.QueryParams["orderby"] = []string{"time"}
		paperInputModel := c.NtmPaperinputStorage.GetAll(r, -1,-1)
		lastPaperInputModel := paperInputModel[len(paperInputModel)-1:][0]
		lastPhase := lastPaperInputModel.Phase
		totalPhase := proposalModel.TotalPhase

		if paperModel.InputState == 1 {
			r.QueryParams["phase"] = []string{strconv.Itoa(lastPhase)}
			result = c.NtmScenarioStorage.GetAll(r, -1, -1)
		} else if paperModel.InputState == 2 && lastPaperInputModel.Phase != totalPhase {
			r.QueryParams["phase"] = []string{strconv.Itoa(lastPhase + 1)}
			result = c.NtmScenarioStorage.GetAll(r, -1, -1)
		} else {
			r.QueryParams["phase"] = []string{"1"}
			result = c.NtmScenarioStorage.GetAll(r, -1, -1)
		}
		return &Response{Res: result}, nil
	}

	if piok {
		modelRootID := paperinputsID[0]
		modelRoot, err := c.NtmPaperinputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := c.NtmScenarioStorage.GetOne(modelRoot.ScenarioID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	if paok {
		modelRootID := personnelAssessmentsID[0]
		modelRoot, err := c.NtmPersonnelAssessmentStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := c.NtmScenarioStorage.GetOne(modelRoot.ScenarioID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	result = c.NtmScenarioStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c NtmScenarioResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.NtmScenarioStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c NtmScenarioResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Scenario)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.NtmScenarioStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c NtmScenarioResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.NtmScenarioStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c NtmScenarioResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(NtmModel.Scenario)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.NtmScenarioStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
