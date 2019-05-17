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

// NtmLevelConfigStorage stores all of the tasty modelleaf, needs to be injected into
// LevelConfig and LevelConfig Resource. In the real world, you would use a database for that.
type NtmLevelConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmLevelConfigStorage) NewLevelConfigStorage(args []BmDaemons.BmDaemon) *NtmLevelConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmLevelConfigStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmLevelConfigStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.LevelConfig {
	in := NtmModel.LevelConfig{}
	var out []NtmModel.LevelConfig
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
func (s NtmLevelConfigStorage) GetOne(id string) (NtmModel.LevelConfig, error) {
	in := NtmModel.LevelConfig{ID: id}
	out := NtmModel.LevelConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("LevelConfig for id %s not found", id)
	return NtmModel.LevelConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmLevelConfigStorage) Insert(c NtmModel.LevelConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmLevelConfigStorage) Delete(id string) error {
	in := NtmModel.LevelConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("LevelConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmLevelConfigStorage) Update(c NtmModel.LevelConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("LevelConfig with id does not exist")
	}

	return nil
}
