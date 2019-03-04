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

// NtmProposalStorage stores all of the tasty modelleaf, needs to be injected into
// Proposal and Proposal Resource. In the real world, you would use a database for that.
type NtmProposalStorage struct {
	Policies  map[string]*NtmModel.Proposal
	idCount int

	db *BmMongodb.BmMongodb
}

func (s NtmProposalStorage) NewProposalStorage(args []BmDaemons.BmDaemon) *NtmProposalStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &NtmProposalStorage{make(map[string]*NtmModel.Proposal), 1, mdb}
}

// GetAll of the modelleaf
func (s NtmProposalStorage) GetAll(r api2go.Request, skip int, take int) []NtmModel.Proposal {
	in := NtmModel.Proposal{}
	var out []NtmModel.Proposal
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
func (s NtmProposalStorage) GetOne(id string) (NtmModel.Proposal, error) {
	in := NtmModel.Proposal{ID: id}
	out := NtmModel.Proposal{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Proposal for id %s not found", id)
	return NtmModel.Proposal{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *NtmProposalStorage) Insert(c NtmModel.Proposal) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *NtmProposalStorage) Delete(id string) error {
	in := NtmModel.Proposal{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Proposal with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *NtmProposalStorage) Update(c NtmModel.Proposal) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Proposal with id does not exist")
	}

	return nil
}
