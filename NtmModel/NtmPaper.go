package NtmModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type Paper struct {
	ID         string        `json:"-"`
	Id_        bson.ObjectId `json:"-" bson:"_id"`
	AccountID  string        `json:"account-id" bson:"account-id"`
	ProposalID string        `json:"proposal-id" bson:"proposal-id"`
	Name       string        `json:"name" bson:"name"`
	Describe   string        `json:"describe" bson:"describe"`
	StartTime  int64         `json:"start-time" bson:"start-time"`
	EndTime    int64         `json:"end-time" bson:"end-time"`
	InputState string        `json:"state" bson:"input-state"`

	InputIDs    []string      `json:"-" bson:"input-ids"`
	PaperInputs []*PaperInput `json:"-"`

	ReportIDs    []string       `json:"-" bson:"report-ids"`
	PaperReports []*PaperReport `json:"-"`
}

func (c Paper) GetID() string {
	return c.ID
}

func (c Paper) SetID(id string) error {
	c.ID = id
	return nil
}

func (c Paper) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "paperInputs",
			Name: "paperInputs",
		},
	}
}

func (c Paper) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range c.InputIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "paperInputs",
			Name: "paperInputs",
		})
	}
	return result
}

func (c Paper) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range c.PaperInputs {
		result = append(result, c.PaperInputs[key])
	}
	return result
}

func (c *Paper) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "PaperInput" {
		c.InputIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Paper) AddToManyIDs(name string, IDs []string) error {
	if name == "PaperInput" {
		c.InputIDs = append(c.InputIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Paper) DeleteToManyIDs(name string, IDs []string) error {
	if name == "PaperInput" {
		for _, ID := range IDs {
			for pos, oldID := range c.InputIDs {
				if ID == oldID {
					c.InputIDs = append(c.InputIDs[:pos], c.InputIDs[pos+1:]...)
				}
			}
		}
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Paper) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "ids":
			r := make(map[string]interface{})
			var ids []bson.ObjectId
			for i := 0; i < len(v); i++ {
				ids = append(ids, bson.ObjectIdHex(v[i]))
			}
			r["$in"] = ids
			rst["_id"] = r
		}
	}
	return rst
}
