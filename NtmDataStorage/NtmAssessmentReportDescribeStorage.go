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

// NtmAssessmentReportDescribeStorage stores all of the tasty modelleaf, needs to be injected into
// AssessmentReportDescribe and AssessmentReportDescribe Resource. In the real world, you would use a database for that.
type NtmAssessmentReportDescribeStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmAssessmentReportDescribeStorage) NewAssessmentReportDescribeStorage(args []BmDaemons.BmDaemon) *NtmAssessmentReportDescribeStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmAssessmentReportDescribeStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmAssessmentReportDescribeStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.AssessmentReportDescribe {
	in := NtmModel.AssessmentReportDescribe{}
	var out []NtmModel.AssessmentReportDescribe
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
func (s NtmAssessmentReportDescribeStorage) GetOne(id string) (NtmModel.AssessmentReportDescribe, error) {
	in := NtmModel.AssessmentReportDescribe{ID: id}
	out := NtmModel.AssessmentReportDescribe{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("AssessmentReportDescribe for id %s not found", id)
	return NtmModel.AssessmentReportDescribe{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmAssessmentReportDescribeStorage) Insert(c NtmModel.AssessmentReportDescribe) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmAssessmentReportDescribeStorage) Delete(id string) error {
	in := NtmModel.AssessmentReportDescribe{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("AssessmentReportDescribe with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmAssessmentReportDescribeStorage) Update(c NtmModel.AssessmentReportDescribe) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("AssessmentReportDescribe with id does not exist")
	}

	return nil
}
