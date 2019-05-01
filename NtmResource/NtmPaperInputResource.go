package NtmResource

import (
	"errors"
	"Ntm/NtmDataStorage"
	"Ntm/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type NtmPaperinputResource struct {
	NtmPaperinputStorage          *NtmDataStorage.NtmPaperinputStorage
	NtmBusinessinputStorage       *NtmDataStorage.NtmBusinessinputStorage
	NtmRepresentativeinputStorage *NtmDataStorage.NtmRepresentativeinputStorage
	NtmManagerinputStorage        *NtmDataStorage.NtmManagerinputStorage
	NtmPaperStorage				  *NtmDataStorage.NtmPaperStorage
	NtmScenarioStorage			  *NtmDataStorage.NtmScenarioStorage
}

func (s NtmPaperinputResource) NewPaperinputResource(args []BmDataStorage.BmStorage) *NtmPaperinputResource {
	var pis *NtmDataStorage.NtmPaperinputStorage
	var bis *NtmDataStorage.NtmBusinessinputStorage
	var ris *NtmDataStorage.NtmRepresentativeinputStorage
	var mis *NtmDataStorage.NtmManagerinputStorage
	var ps *NtmDataStorage.NtmPaperStorage
	var ss *NtmDataStorage.NtmScenarioStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmPaperinputStorage" {
			pis = arg.(*NtmDataStorage.NtmPaperinputStorage)
		} else if tp.Name() == "NtmBusinessinputStorage" {
			bis = arg.(*NtmDataStorage.NtmBusinessinputStorage)
		} else if tp.Name() == "NtmRepresentativeinputStorage" {
			ris = arg.(*NtmDataStorage.NtmRepresentativeinputStorage)
		} else if tp.Name() == "NtmManagerinputStorage" {
			mis = arg.(*NtmDataStorage.NtmManagerinputStorage)
		} else if tp.Name() == "NtmPaperStorage" {
			ps = arg.(*NtmDataStorage.NtmPaperStorage)
		} else if tp.Name() == "NtmScenarioStorage" {
			ss = arg.(*NtmDataStorage.NtmScenarioStorage)
		}
	}
	return &NtmPaperinputResource{
		NtmPaperinputStorage:          	pis,
		NtmBusinessinputStorage:       	bis,
		NtmRepresentativeinputStorage: 	ris,
		NtmManagerinputStorage:        	mis,
		NtmPaperStorage: 				ps,
		NtmScenarioStorage: 			ss,
	}
}

func (s NtmPaperinputResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	papersID, dcok := r.QueryParams["papersID"]
	var result []NtmModel.Paperinput

	if dcok {
		modelRootID := papersID[0]
		modelRoot, err := s.NtmPaperStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.InputIDs

		result := s.NtmPaperinputStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	models := s.NtmPaperinputStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmPaperinputResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.Paperinput
		number, size, offset, limit string
	)

	numberQuery, ok := r.QueryParams["page[number]"]
	if ok {
		number = numberQuery[0]
	}
	sizeQuery, ok := r.QueryParams["page[size]"]
	if ok {
		size = sizeQuery[0]
	}
	offsetQuery, ok := r.QueryParams["page[offset]"]
	if ok {
		offset = offsetQuery[0]
	}
	limitQuery, ok := r.QueryParams["page[limit]"]
	if ok {
		limit = limitQuery[0]
	}

	if size != "" {
		sizeI, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		numberI, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		start := sizeI * (numberI - 1)
		for _, iter := range s.NtmPaperinputStorage.GetAll(r, int(start), int(sizeI)) {
			result = append(result, *iter)
		}

	} else {
		limitI, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		offsetI, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		for _, iter := range s.NtmPaperinputStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.Paperinput{}
	count := s.NtmPaperinputStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmPaperinputResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmPaperinputStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	err = s.ResetReferencedModel(&model, &r)

	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmPaperinputResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Paperinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmPaperinputStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmPaperinputResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmPaperinputStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmPaperinputResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Paperinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmPaperinputStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}

func (s NtmPaperinputResource) ResetReferencedModel(model *NtmModel.Paperinput, r *api2go.Request) error {
	model.Businessinputs = []*NtmModel.Businessinput{}
	r.QueryParams["ids"] = model.BusinessinputIDs
	for _, Businessinput := range s.NtmBusinessinputStorage.GetAll(*r, -1, -1) {
		model.Businessinputs = append(model.Businessinputs, Businessinput)
	}

	model.Representativeinputs = []*NtmModel.Representativeinput{}
	r.QueryParams["ids"] = model.RepresentativeinputIDs
	for _, Representativeinput := range s.NtmRepresentativeinputStorage.GetAll(*r, -1, -1) {
		model.Representativeinputs = append(model.Representativeinputs, Representativeinput)
	}

	model.Managerinputs = []*NtmModel.Managerinput{}
	r.QueryParams["ids"] = model.ManagerinputIDs
	for _, manageInput := range s.NtmManagerinputStorage.GetAll(*r, -1, -1) {
		model.Managerinputs = append(model.Managerinputs, manageInput)
	}

	result, _ := s.NtmScenarioStorage.GetOne(model.ScenarioID)
	model.Scenario = &result

	return nil
}
