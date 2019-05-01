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

// NtmRepresentativeAbilityStorage stores all of the tasty modelleaf, needs to be injected into
// RepresentativeAbility and RepresentativeAbility Resource. In the real world, you would use a database for that.
type NtmRepresentativeAbilityStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmRepresentativeAbilityStorage) NewRepresentativeAbilityStorage(args []BmDaemons.BmDaemon) *NtmRepresentativeAbilityStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmRepresentativeAbilityStorage{mdb}
}

// GetAll of the modelleaf
func (s NtmRepresentativeAbilityStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.RepresentativeAbility {
	in := NtmModel.RepresentativeAbility{}
	var out []*NtmModel.RepresentativeAbility
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
func (s NtmRepresentativeAbilityStorage) GetOne(id string) (NtmModel.RepresentativeAbility, error) {
	in := NtmModel.RepresentativeAbility{ID: id}
	out := NtmModel.RepresentativeAbility{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("RepresentativeAbility for id %s not found", id)
	return NtmModel.RepresentativeAbility{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmRepresentativeAbilityStorage) Insert(c NtmModel.RepresentativeAbility) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmRepresentativeAbilityStorage) Delete(id string) error {
	in := NtmModel.RepresentativeAbility{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("RepresentativeAbility with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmRepresentativeAbilityStorage) Update(c NtmModel.RepresentativeAbility) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("RepresentativeAbility with id does not exist")
	}

	return nil
}
