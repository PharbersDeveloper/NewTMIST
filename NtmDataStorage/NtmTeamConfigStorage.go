package NtmDataStorage

import (
	"errors"
	"fmt"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
	"net/http"
)

type NtmTeamConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmTeamConfigStorage) NewTeamConfigStorage(args []BmDaemons.BmDaemon) *NtmTeamConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmTeamConfigStorage{mdb}
}

func (s NtmTeamConfigStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.TeamConfig {
	in := NtmModel.TeamConfig{}
	var out []NtmModel.TeamConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.TeamConfig
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil
	}
}

func (s NtmTeamConfigStorage) GetOne(id string) (NtmModel.TeamConfig, error) {
	in := NtmModel.TeamConfig{ID: id}
	out := NtmModel.TeamConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("TeamConfig for id %s not found", id)
	return NtmModel.TeamConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *NtmTeamConfigStorage) Insert(c NtmModel.TeamConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmTeamConfigStorage) Delete(id string) error {
	in := NtmModel.TeamConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("TeamConfig with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *NtmTeamConfigStorage) Update(c NtmModel.TeamConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("TeamConfig with id does not exist")
	}

	return nil
}

func (s *NtmTeamConfigStorage) Count(req api2go.Request, c NtmModel.TeamConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}