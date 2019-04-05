package NtmResource

import (
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmDataStorage"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type NtmRepresentativeAbilityResource struct {
	NtmRepresentativeAbilityStorage *NtmDataStorage.NtmRepresentativeAbilityStorage
	NtmPersonnelAssessmentStorage	*NtmDataStorage.NtmPersonnelAssessmentStorage
}

func (s NtmRepresentativeAbilityResource) NewRepresentativeAbilityResource(args []BmDataStorage.BmStorage) *NtmRepresentativeAbilityResource {
	var ras *NtmDataStorage.NtmRepresentativeAbilityStorage
	var pas *NtmDataStorage.NtmPersonnelAssessmentStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmRepresentativeAbilityStorage" {
			ras = arg.(*NtmDataStorage.NtmRepresentativeAbilityStorage)
		} else if tp.Name() == "NtmPersonnelAssessmentStorage" {
			pas = arg.(*NtmDataStorage.NtmPersonnelAssessmentStorage)
		}
	}
	return &NtmRepresentativeAbilityResource{
		NtmRepresentativeAbilityStorage: ras,
		NtmPersonnelAssessmentStorage:   pas,
	}
}

func (s NtmRepresentativeAbilityResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.RepresentativeAbility

	personnelAssessmentsID, pasok := r.QueryParams["personnelAssessmentsID"]

	if pasok {
		modelRootID := personnelAssessmentsID[0]

		modelRoot, err := s.NtmPersonnelAssessmentStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}

		r.QueryParams["ids"] = modelRoot.RepresentativeAbilityIDs

		result := s.NtmRepresentativeAbilityStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	models := s.NtmRepresentativeAbilityStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmRepresentativeAbilityResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	modelRoot, err := s.NtmRepresentativeAbilityStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: modelRoot}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmRepresentativeAbilityResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.RepresentativeAbility)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmRepresentativeAbilityStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmRepresentativeAbilityResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmRepresentativeAbilityStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmRepresentativeAbilityResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.RepresentativeAbility)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmRepresentativeAbilityStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
