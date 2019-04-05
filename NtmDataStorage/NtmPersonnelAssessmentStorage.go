package NtmDataStorage

import (
	"fmt"
	"errors"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
)

// NtmPersonnelAssessmentStorage stores all of the tasty modelleaf, needs to be injected into
// PersonnelAssessment and PersonnelAssessment Resource. In the real world, you would use a database for that.
type NtmPersonnelAssessmentStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmPersonnelAssessmentStorage) NewPersonnelAssessmentStorage(args []BmDaemons.BmDaemon) *NtmPersonnelAssessmentStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmPersonnelAssessmentStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmPersonnelAssessmentStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.PersonnelAssessment {
	in := NtmModel.PersonnelAssessment{}
	var out []*NtmModel.PersonnelAssessment
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s NtmPersonnelAssessmentStorage) GetOne(id string) (NtmModel.PersonnelAssessment, error) {
	in := NtmModel.PersonnelAssessment{ID: id}
	out := NtmModel.PersonnelAssessment{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("PersonnelAssessment for id %s not found", id)
	return NtmModel.PersonnelAssessment{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmPersonnelAssessmentStorage) Insert(c NtmModel.PersonnelAssessment) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmPersonnelAssessmentStorage) Delete(id string) error {
	in := NtmModel.PersonnelAssessment{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("PersonnelAssessment with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmPersonnelAssessmentStorage) Update(c NtmModel.PersonnelAssessment) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("PersonnelAssessment with id does not exist")
	}

	return nil
}

