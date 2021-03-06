package NtmResource

import (
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmDataStorage"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type NtmPaperInputResource struct {
	NtmPaperInputStorage          *NtmDataStorage.NtmPaperInputStorage
	NtmBusinessInputStorage       *NtmDataStorage.NtmBusinessInputStorage
	NtmRepresentativeInputStorage *NtmDataStorage.NtmRepresentativeInputStorage
	NtmManagerInputStorage        *NtmDataStorage.NtmManagerInputStorage
	NtmPaperStorage				  *NtmDataStorage.NtmPaperStorage
}

func (s NtmPaperInputResource) NewPaperInputResource(args []BmDataStorage.BmStorage) *NtmPaperInputResource {
	var pis *NtmDataStorage.NtmPaperInputStorage
	var bis *NtmDataStorage.NtmBusinessInputStorage
	var ris *NtmDataStorage.NtmRepresentativeInputStorage
	var mis *NtmDataStorage.NtmManagerInputStorage
	var ps *NtmDataStorage.NtmPaperStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmPaperInputStorage" {
			pis = arg.(*NtmDataStorage.NtmPaperInputStorage)
		} else if tp.Name() == "NtmBusinessInputStorage" {
			bis = arg.(*NtmDataStorage.NtmBusinessInputStorage)
		} else if tp.Name() == "NtmRepresentativeInputStorage" {
			ris = arg.(*NtmDataStorage.NtmRepresentativeInputStorage)
		} else if tp.Name() == "NtmManagerInputStorage" {
			mis = arg.(*NtmDataStorage.NtmManagerInputStorage)
		} else if tp.Name() == "NtmPaperStorage" {
			ps = arg.(*NtmDataStorage.NtmPaperStorage)
		}
	}
	return &NtmPaperInputResource{
		NtmPaperInputStorage:          pis,
		NtmBusinessInputStorage:       bis,
		NtmRepresentativeInputStorage: ris,
		NtmManagerInputStorage:        mis,
		NtmPaperStorage: ps,
	}
}

func (s NtmPaperInputResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	papersID, dcok := r.QueryParams["papersID"]
	var result []NtmModel.PaperInput

	if dcok {
		modelRootID := papersID[0]
		modelRoot, err := s.NtmPaperStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.InputIDs

		result := s.NtmPaperInputStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	models := s.NtmPaperInputStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmPaperInputResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.PaperInput
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
		for _, iter := range s.NtmPaperInputStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmPaperInputStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.PaperInput{}
	count := s.NtmPaperInputStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmPaperInputResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmPaperInputStorage.GetOne(ID)
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
func (s NtmPaperInputResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.PaperInput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmPaperInputStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmPaperInputResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmPaperInputStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmPaperInputResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.PaperInput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmPaperInputStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}

func (s NtmPaperInputResource) ResetReferencedModel(model *NtmModel.PaperInput, r *api2go.Request) error {
	model.BusinessInputs = []*NtmModel.BusinessInput{}
	r.QueryParams["ids"] = model.BusinessInputIDs
	for _, businessInput := range s.NtmBusinessInputStorage.GetAll(*r, -1, -1) {
		model.BusinessInputs = append(model.BusinessInputs, businessInput)
	}

	model.RepresentativeInputs = []*NtmModel.RepresentativeInput{}
	r.QueryParams["ids"] = model.RepresentativeInputIDs
	for _, representativeInput := range s.NtmRepresentativeInputStorage.GetAll(*r, -1, -1) {
		model.RepresentativeInputs = append(model.RepresentativeInputs, representativeInput)
	}

	model.ManagerInputs = []*NtmModel.ManagerInput{}
	r.QueryParams["ids"] = model.ManagerInputIDs
	for _, manageInput := range s.NtmManagerInputStorage.GetAll(*r, -1, -1) {
		model.ManagerInputs = append(model.ManagerInputs, manageInput)
	}

	return nil
}
