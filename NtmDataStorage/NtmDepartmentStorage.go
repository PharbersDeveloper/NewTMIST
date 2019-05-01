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

// NtmDepartmentStorage stores all of the tasty modelleaf, needs to be injected into
// Department and Department Resource. In the real world, you would use a database for that.
type NtmDepartmentStorage struct {
	Departments map[string]*NtmModel.Department
	idCount     int

	db *BmMongodb.BmMongodb
}

func (s NtmDepartmentStorage) NewDepartmentStorage(args []BmDaemons.BmDaemon) *NtmDepartmentStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmDepartmentStorage{make(map[string]*NtmModel.Department), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmDepartmentStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.Department {
	in := NtmModel.Department{}
	var out []*NtmModel.Department
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
func (s NtmDepartmentStorage) GetOne(id string) (NtmModel.Department, error) {
	in := NtmModel.Department{ID: id}
	out := NtmModel.Department{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Department for id %s not found", id)
	return NtmModel.Department{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmDepartmentStorage) Insert(c NtmModel.Department) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmDepartmentStorage) Delete(id string) error {
	in := NtmModel.Department{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Department with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmDepartmentStorage) Update(c NtmModel.Department) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Department with id does not exist")
	}

	return nil
}
