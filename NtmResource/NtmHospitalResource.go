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

type NtmHospitalResource struct {
	NtmHospitalStorage *NtmDataStorage.NtmHospitalStorage
	NtmImageStorage    *NtmDataStorage.NtmImageStorage
	NtmHospitalConfigStorage *NtmDataStorage.NtmHospitalConfigStorage
}

func (s NtmHospitalResource) NewHospitalResource (args []BmDataStorage.BmStorage) *NtmHospitalResource {
	var is *NtmDataStorage.NtmImageStorage
	var hs *NtmDataStorage.NtmHospitalStorage
	var hcs *NtmDataStorage.NtmHospitalConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmImageStorage" {
			is = arg.(*NtmDataStorage.NtmImageStorage)
		} else if tp.Name() == "NtmHospitalStorage" {
			hs = arg.(*NtmDataStorage.NtmHospitalStorage)
		} else if tp.Name() == "NtmHospitalConfigStorage" {
			hcs = arg.(*NtmDataStorage.NtmHospitalConfigStorage)
		}
	}
	return &NtmHospitalResource{
		NtmImageStorage: is,
		NtmHospitalStorage: hs,
		NtmHospitalConfigStorage: hcs,
	}
}

func (s NtmHospitalResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*NtmModel.Hospital

	hospitalConfigsID, pciok := r.QueryParams["hospitalConfigsID"]

	if pciok {
		modelRootID := hospitalConfigsID[0]

		modelRoot, err := s.NtmHospitalConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		r.QueryParams["ids"] = []string{modelRoot.HospitalID} //这块有问题，待会想一下


		result = s.NtmHospitalStorage.GetAll(r, -1, -1)

		return &Response{Res: result}, nil
	}

	models := s.NtmHospitalStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmHospitalResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.Hospital
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
		for _, iter := range s.NtmHospitalStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmHospitalStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.Hospital{}
	count := s.NtmHospitalStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmHospitalResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmHospitalStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	model.Imgs = []*NtmModel.Image{}
	for _, kID := range model.ImagesIDs {
		choc, err := s.NtmImageStorage.GetOne(kID)
		if err != nil {
			return &Response{}, err
		}
		model.Imgs = append(model.Imgs, &choc)
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmHospitalResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Hospital)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmHospitalStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmHospitalResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmHospitalStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmHospitalResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Hospital)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmHospitalStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
