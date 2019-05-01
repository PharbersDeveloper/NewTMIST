package NtmResource

import (
	"errors"
	"Ntm/NtmDataStorage"
	"Ntm/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type NtmPersonnelAssessmentResource struct {
	NtmPersonnelAssessmentStorage 			*NtmDataStorage.NtmPersonnelAssessmentStorage
	NtmPaperStorage							*NtmDataStorage.NtmPaperStorage
}

func (s NtmPersonnelAssessmentResource) NewPersonnelAssessmentResource(args []BmDataStorage.BmStorage) *NtmPersonnelAssessmentResource {
	var pas *NtmDataStorage.NtmPersonnelAssessmentStorage
	var ps  *NtmDataStorage.NtmPaperStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmPersonnelAssessmentStorage" {
			pas = arg.(*NtmDataStorage.NtmPersonnelAssessmentStorage)
		} else if tp.Name() == "NtmPaperStorage" {
			ps = arg.(*NtmDataStorage.NtmPaperStorage)
		}
	}
	return &NtmPersonnelAssessmentResource{
		NtmPersonnelAssessmentStorage:		pas,
		NtmPaperStorage: 					ps,
	}
}

func (s NtmPersonnelAssessmentResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.PersonnelAssessment
	papersID, pok := r.QueryParams["papersID"]

	if pok {
		modelRootID := papersID[0]
		modelRoot, err := s.NtmPaperStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.PersonnelAssessmentIDs

		result := s.NtmPersonnelAssessmentStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	models := s.NtmPersonnelAssessmentStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmPersonnelAssessmentResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	modelRoot, err := s.NtmPersonnelAssessmentStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: modelRoot}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmPersonnelAssessmentResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.PersonnelAssessment)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmPersonnelAssessmentStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmPersonnelAssessmentResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmPersonnelAssessmentStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmPersonnelAssessmentResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.PersonnelAssessment)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmPersonnelAssessmentStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
