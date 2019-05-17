package NtmDataStorage

import (
	"fmt"
	"errors"
	"Ntm/NtmModel"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
)

// NtmAssessmentReportStorage stores all of the tasty modelleaf, needs to be injected into
// AssessmentReport and AssessmentReport Resource. In the real world, you would use a database for that.
type NtmAssessmentReportStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmAssessmentReportStorage) NewAssessmentReportStorage(args []BmDaemons.BmDaemon) *NtmAssessmentReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmAssessmentReportStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmAssessmentReportStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.AssessmentReport {
	in := NtmModel.AssessmentReport{}
	var out []NtmModel.AssessmentReport
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(&iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s NtmAssessmentReportStorage) GetOne(id string) (NtmModel.AssessmentReport, error) {
	in := NtmModel.AssessmentReport{ID: id}
	out := NtmModel.AssessmentReport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("AssessmentReport for id %s not found", id)
	return NtmModel.AssessmentReport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmAssessmentReportStorage) Insert(c NtmModel.AssessmentReport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmAssessmentReportStorage) Delete(id string) error {
	in := NtmModel.AssessmentReport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("AssessmentReport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmAssessmentReportStorage) Update(c NtmModel.AssessmentReport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("AssessmentReport with id does not exist")
	}

	return nil
}
