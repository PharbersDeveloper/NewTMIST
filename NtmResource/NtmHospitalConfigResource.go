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

type NtmHospitalConfigResource struct {
	NtmHospitalConfigStorage 	*NtmDataStorage.NtmHospitalConfigStorage
	NtmHospitalStorage          *NtmDataStorage.NtmHospitalStorage
	NtmPolicyStorage			*NtmDataStorage.NtmPolicyStorage
	NtmDepartmentStorage		*NtmDataStorage.NtmDepartmentStorage
}

func (s NtmHospitalConfigResource) NewHospitalConfigResource(args []BmDataStorage.BmStorage) NtmHospitalConfigResource {
	var is *NtmDataStorage.NtmHospitalStorage
	var hs *NtmDataStorage.NtmHospitalConfigStorage
	var ps *NtmDataStorage.NtmPolicyStorage
	var ds *NtmDataStorage.NtmDepartmentStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmHospitalStorage" {
			is = arg.(*NtmDataStorage.NtmHospitalStorage)
		} else if tp.Name() == "NtmHospitalConfigStorage" {
			hs = arg.(*NtmDataStorage.NtmHospitalConfigStorage)
		} else if tp.Name() == "NtmPolicyStorage" {
			ps = arg.(*NtmDataStorage.NtmPolicyStorage)
		} else if tp.Name() == "NtmDepartmentStorage" {
			ds = arg.(*NtmDataStorage.NtmDepartmentStorage)
		}
	}
	return NtmHospitalConfigResource{NtmHospitalStorage: is, NtmHospitalConfigStorage: hs, NtmPolicyStorage: ps, NtmDepartmentStorage: ds}
}

func (s NtmHospitalConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []NtmModel.HospitalConfig
	models := s.NtmHospitalConfigStorage.GetAll(r, -1, -1)

	for _, model := range models {
		err := s.ResetReferencedModel(model)

		if err != nil {
			return &Response{}, err
		}

		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmHospitalConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.HospitalConfig
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
		for _, iter := range s.NtmHospitalConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmHospitalConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.HospitalConfig{}
	count := s.NtmHospitalConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmHospitalConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmHospitalConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	err = s.ResetReferencedModel(&model)

	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmHospitalConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.HospitalConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmHospitalConfigStorage.Insert(model)
	model.ID = id

	s.ResetReferencedModel(&model)

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmHospitalConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmHospitalConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmHospitalConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.HospitalConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmHospitalConfigStorage.Update(model)

	s.ResetReferencedModel(&model)

	return &Response{Res: model, Code: http.StatusNoContent}, err
}

func (s NtmHospitalConfigResource) ResetReferencedModel(model *NtmModel.HospitalConfig) error {
	model.Policies = []*NtmModel.Policy{}
	for _, tmpID := range model.PolicyIDs {
		choc, err := s.NtmPolicyStorage.GetOne(tmpID)
		if err != nil {
			return err
		}
		model.Policies = append(model.Policies, &choc)
	}
	model.Departments = []*NtmModel.Department{}
	for _, tmpID := range model.DepartmentIDs {
		choc, err := s.NtmDepartmentStorage.GetOne(tmpID)
		if err != nil {
			return err
		}
		model.Departments = append(model.Departments, &choc)
	}

	if model.HospitalID != "" {
		hospital, err := s.NtmHospitalStorage.GetOne(model.HospitalID)
		if err != nil {
			return err
		}
		model.Hospital = hospital
	}

	return nil
}
