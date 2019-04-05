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

type NtmRepresentativeResource struct {
	NtmRepresentativeStorage       *NtmDataStorage.NtmRepresentativeStorage
	NtmRepresentativeConfigStorage *NtmDataStorage.NtmRepresentativeConfigStorage
	NtmImageStorage                *NtmDataStorage.NtmImageStorage
	NtmActionKpiStorage			   *NtmDataStorage.NtmActionKpiStorage
	NtmRepresentativeAbilityStorage *NtmDataStorage.NtmRepresentativeAbilityStorage
}

func (s NtmRepresentativeResource) NewRepresentativeResource(args []BmDataStorage.BmStorage) *NtmRepresentativeResource {
	var is *NtmDataStorage.NtmImageStorage
	var reps *NtmDataStorage.NtmRepresentativeStorage
	var repcs *NtmDataStorage.NtmRepresentativeConfigStorage
	var aks	*NtmDataStorage.NtmActionKpiStorage
	var ras *NtmDataStorage.NtmRepresentativeAbilityStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmImageStorage" {
			is = arg.(*NtmDataStorage.NtmImageStorage)
		} else if tp.Name() == "NtmRepresentativeStorage" {
			reps = arg.(*NtmDataStorage.NtmRepresentativeStorage)
		} else if tp.Name() == "NtmRepresentativeConfigStorage" {
			repcs = arg.(*NtmDataStorage.NtmRepresentativeConfigStorage)
		} else if tp.Name() == "NtmActionKpiStorage" {
			aks = arg.(*NtmDataStorage.NtmActionKpiStorage)
		} else if tp.Name() == "NtmRepresentativeAbilityStorage" {
			ras = arg.(*NtmDataStorage.NtmRepresentativeAbilityStorage)
		}
	}
	return &NtmRepresentativeResource{
		NtmImageStorage:                is,
		NtmRepresentativeStorage:       reps,
		NtmRepresentativeConfigStorage: repcs,
		NtmActionKpiStorage: 			aks,
		NtmRepresentativeAbilityStorage: ras,
	}
}

func (s NtmRepresentativeResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	representativeConfigsID, rcok := r.QueryParams["representativeConfigsID"]
	actionKpisID, akok := r.QueryParams["actionKpisID"]
	representativeAbilitiesID, raok := r.QueryParams["representativeAbilitiesID"]
	var result []NtmModel.Representative

	if rcok {
		modelRootID := representativeConfigsID[0]
		modelRoot, err := s.NtmRepresentativeConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := s.NtmRepresentativeStorage.GetOne(modelRoot.RepresentativeID)
		if err != nil {
			return &Response{}, err
		}
		result = append(result, model)
		return &Response{Res: result}, nil
	}

	if akok {
		modelRootID := actionKpisID[0]
		modelRoot, err := s.NtmActionKpiStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := s.NtmRepresentativeStorage.GetOne(modelRoot.RepresentativeID)
		if err != nil {
			return &Response{}, err
		}
		result = append(result, model)
		return &Response{Res: result}, nil
	}

	if raok {
		modelRootID := representativeAbilitiesID[0]
		modelRoot, err := s.NtmRepresentativeAbilityStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := s.NtmRepresentativeStorage.GetOne(modelRoot.RepresentativeID)
		if err != nil {
			return &Response{}, err
		}
		result = append(result, model)
		return &Response{Res: result}, nil
	}

	models := s.NtmRepresentativeStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmRepresentativeResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.Representative
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
		for _, iter := range s.NtmRepresentativeStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmRepresentativeStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.Representative{}
	count := s.NtmRepresentativeStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmRepresentativeResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmRepresentativeStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	r.QueryParams["ids"] = model.ImagesIDs
	images := s.NtmImageStorage.GetAll(r, -1, -1)
	for _, image := range images {
		model.Imgs = append(model.Imgs, &image)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmRepresentativeResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Representative)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmRepresentativeStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmRepresentativeResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmRepresentativeStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmRepresentativeResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Representative)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmRepresentativeStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
