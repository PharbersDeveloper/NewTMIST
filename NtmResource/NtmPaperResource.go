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

type NtmPaperResource struct {
	NtmPaperStorage			*NtmDataStorage.NtmPaperStorage
	NtmPaperinputStorage	*NtmDataStorage.NtmPaperinputStorage
	NtmSalesReportStorage	*NtmDataStorage.NtmSalesReportStorage
	NtmPersonnelAssessmentStorage	*NtmDataStorage.NtmPersonnelAssessmentStorage
}

func (s NtmPaperResource) NewPaperResource (args []BmDataStorage.BmStorage) *NtmPaperResource {
	var ps *NtmDataStorage.NtmPaperStorage
	var pis *NtmDataStorage.NtmPaperinputStorage
	var srs *NtmDataStorage.NtmSalesReportStorage
	var pas *NtmDataStorage.NtmPersonnelAssessmentStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "NtmPaperStorage" {
			ps = arg.(*NtmDataStorage.NtmPaperStorage)
		} else if tp.Name() == "NtmPaperinputStorage" {
			pis = arg.(*NtmDataStorage.NtmPaperinputStorage)
		} else if tp.Name() == "NtmSalesReportStorage" {
			srs = arg.(*NtmDataStorage.NtmSalesReportStorage)
		} else if tp.Name() == "NtmPersonnelAssessmentStorage" {
			pas = arg.(*NtmDataStorage.NtmPersonnelAssessmentStorage)
		}
	}
	return &NtmPaperResource{
		NtmPaperinputStorage: pis,
		NtmPaperStorage: ps,
		NtmSalesReportStorage: srs,
		NtmPersonnelAssessmentStorage: pas,
	}
}

func (s NtmPaperResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	//var result []NtmModel.Paper
	result := s.NtmPaperStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s NtmPaperResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []NtmModel.Paper
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
		for _, iter := range s.NtmPaperStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.NtmPaperStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := NtmModel.Paper{}
	count := s.NtmPaperStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s NtmPaperResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.NtmPaperStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	model.Paperinputs = []*NtmModel.Paperinput{}
	r.QueryParams["ids"] = model.InputIDs
	paperInputModels := s.NtmPaperinputStorage.GetAll(r, -1,-1)
	model.Paperinputs = paperInputModels

	model.SalesReports = []*NtmModel.SalesReport{}
	r.QueryParams["ids"] = model.SalesReportIDs
	salesReportModels := s.NtmSalesReportStorage.GetAll(r, -1,-1)
	model.SalesReports = salesReportModels

	model.PersonnelAssessment = []*NtmModel.PersonnelAssessment{}
	r.QueryParams["ids"] = model.PersonnelAssessmentIDs
	personnelAssessmentModels := s.NtmPersonnelAssessmentStorage.GetAll(r, -1,-1)
	model.PersonnelAssessment = personnelAssessmentModels

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s NtmPaperResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Paper)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.NtmPaperStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s NtmPaperResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.NtmPaperStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s NtmPaperResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(NtmModel.Paper)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.NtmPaperStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}