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

type NtmUseableProposalStorage struct {
	db *BmMongodb.BmMongodb
}

func (s NtmUseableProposalStorage) NewUseableProposalStorage(args []BmDaemons.BmDaemon) *NtmUseableProposalStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmUseableProposalStorage{mdb}
}

func (s NtmUseableProposalStorage) GetAll(r api2go.Request, skip int, take int) []*NtmModel.UseableProposal {
	in := NtmModel.UseableProposal{}
	var out []NtmModel.UseableProposal
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*NtmModel.UseableProposal
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

func (s NtmUseableProposalStorage) GetOne(id string) (NtmModel.UseableProposal, error) {
	in := NtmModel.UseableProposal{ID: id}
	out := NtmModel.UseableProposal{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("UseableProposal for id %s not found", id)
	return NtmModel.UseableProposal{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *NtmUseableProposalStorage) Insert(c NtmModel.UseableProposal) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmUseableProposalStorage) Delete(id string) error {
	in := NtmModel.UseableProposal{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("UseableProposal with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *NtmUseableProposalStorage) Update(c NtmModel.UseableProposal) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("UseableProposal with id does not exist")
	}

	return nil
}

func (s *NtmUseableProposalStorage) Count(req api2go.Request, c NtmModel.UseableProposal) int {
	r, _ := s.db.Count(req, &c)
	return r
}
