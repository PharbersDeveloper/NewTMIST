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

// NtmGoodsConfigStorage stores all of the tasty chocolate, needs to be injected into
// GoodsConfig Goods. In the real world, you would use a database for that.
type NtmGoodsConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmGoodsConfigStorage) NewGoodsConfigStorage(args []BmDaemons.BmDaemon) *NtmGoodsConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmGoodsConfigStorage{mdb}
}

// GetAll of the chocolate
func (s NtmGoodsConfigStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.GoodsConfig {
	in := NtmModel.GoodsConfig{}
	var out []NtmModel.GoodsConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.GoodsConfig
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
func (s NtmGoodsConfigStorage) GetOne(id string) (NtmModel.GoodsConfig, error) {
	in := NtmModel.GoodsConfig{ID: id}
	out := NtmModel.GoodsConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("GoodsConfig for id %s not found", id)
	return NtmModel.GoodsConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmGoodsConfigStorage) Insert(c NtmModel.GoodsConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmGoodsConfigStorage) Delete(id string) error {
	in := NtmModel.GoodsConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("GoodsConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *NtmGoodsConfigStorage) Update(c NtmModel.GoodsConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("GoodsConfig with id does not exist")
	}

	return nil
}

func (s *NtmGoodsConfigStorage) Count(req api2go.Request, c NtmModel.GoodsConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
