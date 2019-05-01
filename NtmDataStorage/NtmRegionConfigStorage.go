package NtmDataStorage

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"Ntm/NtmModel"
	"github.com/manyminds/api2go"
	"net/http"
)

// NtmRegionConfigStorage stores all of the tasty chocolate, needs to be injected into
// RegionConfig Resource. In the real world, you would use a database for that.
type NtmRegionConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmRegionConfigStorage) NewRegionConfigStorage(args []BmDaemons.BmDaemon) *NtmRegionConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmRegionConfigStorage{mdb}
}

// GetAll of the chocolate
func (s NtmRegionConfigStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.RegionConfig {
	in := NtmModel.RegionConfig{}
	var out []NtmModel.RegionConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.RegionConfig
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
func (s NtmRegionConfigStorage) GetOne(id string) (NtmModel.RegionConfig, error) {
	in := NtmModel.RegionConfig{ID: id}
	out := NtmModel.RegionConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		// TODO: 双重绑定没明白啥意思
		//双重绑定
		//if out.RegionID != "" {
		//	item, err := NtmRegionStorage{db: s.db}.GetOne(out.RegionID)
		//	if err != nil {
		//		return NtmModel.RegionConfig{}, err
		//	}
		//	out.Region = item
		//}

		return out, nil
	}
	errMessage := fmt.Sprintf("RegionConfig for id %s not found", id)
	return NtmModel.RegionConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmRegionConfigStorage) Insert(c NtmModel.RegionConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmRegionConfigStorage) Delete(id string) error {
	in := NtmModel.RegionConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("RegionConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *NtmRegionConfigStorage) Update(c NtmModel.RegionConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("RegionConfig with id does not exist")
	}

	return nil
}

func (s *NtmRegionConfigStorage) Count(req api2go.Request, c NtmModel.RegionConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
