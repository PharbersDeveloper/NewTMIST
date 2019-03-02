package NtmDataStorage

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"github.com/manyminds/api2go"
	"net/http"
)

// NtmHospitalConfigStorage stores all of the tasty chocolate, needs to be injected into
// HospitalConfig Resource. In the real world, you would use a database for that.
type NtmHospitalConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmHospitalConfigStorage) NewHospitalConfigStorage(args []BmDaemons.BmDaemon) *NtmHospitalConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmHospitalConfigStorage{mdb}
}

// GetAll of the chocolate
func (s NtmHospitalConfigStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.HospitalConfig {
	in := NtmModel.HospitalConfig{}
	var out []NtmModel.HospitalConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.HospitalConfig
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Student)
	}
}

// GetOne
func (s NtmHospitalConfigStorage) GetOne(id string) (NtmModel.HospitalConfig, error) {
	in := NtmModel.HospitalConfig{ID: id}
	out := NtmModel.HospitalConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		// TODO: 双重绑定没明白啥意思
		//双重绑定
		//if out.HospitalID != "" {
		//	item, err := NtmHospitalStorage{db: s.db}.GetOne(out.HospitalID)
		//	if err != nil {
		//		return NtmModel.HospitalConfig{}, err
		//	}
		//	out.Hospital = item
		//}

		return out, nil
	}
	errMessage := fmt.Sprintf("HospitalConfig for id %s not found", id)
	return NtmModel.HospitalConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmHospitalConfigStorage) Insert(c NtmModel.HospitalConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmHospitalConfigStorage) Delete(id string) error {
	in := NtmModel.HospitalConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("HospitalConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *NtmHospitalConfigStorage) Update(c NtmModel.HospitalConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("HospitalConfig with id does not exist")
	}

	return nil
}

func (s *NtmHospitalConfigStorage) Count(req api2go.Request, c NtmModel.HospitalConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
