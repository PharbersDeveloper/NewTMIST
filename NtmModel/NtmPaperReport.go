package NtmModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// 未完成
type PaperReport struct {
	ID      string        `json:"-"`
	Id_     bson.ObjectId `json:"-" bson:"_id"`
	PaperId string        `json:"paper-id" bson:"paper-id"`
	Phase   float64       `json:"phase" bson:"phase"`

	BusinessinputIDs []string         `json:"-" bson:"business-input-ids"`
	Businessinputs   []*Businessinput `json:"-"`

	RepresentativeinputIDs []string               `json:"-" bson:"representative-input-ids"`
	Representativeinputs   []*Representativeinput `json:"-"`

	ManagerinputIDs []string        `json:"-" bson:"manager-input-ids"`
	Managerinputs   []*Managerinput `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c PaperReport) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *PaperReport) SetID(id string) error {
	c.ID = id
	return nil
}

func (c PaperReport) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "Businessinputs",
			Name: "Businessinputs",
		},
		{
			Type: "Representativeinputs",
			Name: "Representativeinputs",
		},
		{
			Type: "Managerinputs",
			Name: "Managerinputs",
		},
	}
}

func (c PaperReport) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range c.BusinessinputIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "Businessinputs",
			Name: "Businessinputs",
		})
	}

	for _, kID := range c.RepresentativeinputIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "Representativeinputs",
			Name: "Representativeinputs",
		})
	}

	for _, kID := range c.ManagerinputIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "Managerinputs",
			Name: "Managerinputs",
		})
	}

	return result
}

func (c PaperReport) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range c.Businessinputs {
		result = append(result, c.Businessinputs[key])
	}

	for key := range c.Representativeinputs {
		result = append(result, c.Representativeinputs[key])
	}

	for key := range c.Managerinputs {
		result = append(result, c.Managerinputs[key])
	}

	return result
}

func (c *PaperReport) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "Businessinputs" {
		c.BusinessinputIDs = IDs
		return nil
	} else if name == "Representativeinputs" {
		c.RepresentativeinputIDs = IDs
		return nil
	} else if name == "Managerinputs" {
		c.ManagerinputIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *PaperReport) AddToManyIDs(name string, IDs []string) error {
	if name == "Businessinputs" {
		c.BusinessinputIDs = append(c.BusinessinputIDs, IDs...)
		return nil
	} else if name == "Representativeinputs" {
		c.RepresentativeinputIDs = append(c.RepresentativeinputIDs, IDs...)
		return nil
	} else if name == "Managerinputs" {
		c.ManagerinputIDs = append(c.ManagerinputIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *PaperReport) DeleteToManyIDs(name string, IDs []string) error {
	if name == "Businessinputs" {
		for _, ID := range IDs {
			for pos, oldID := range c.BusinessinputIDs {
				if ID == oldID {
					c.BusinessinputIDs = append(c.BusinessinputIDs[:pos], c.BusinessinputIDs[pos+1:]...)
				}
			}
		}
	} else if name == "Representativeinputs" {
		for _, ID := range IDs {
			for pos, oldID := range c.RepresentativeinputIDs {
				if ID == oldID {
					c.RepresentativeinputIDs = append(c.RepresentativeinputIDs[:pos], c.RepresentativeinputIDs[pos+1:]...)
				}
			}
		}
	} else if name == "Managerinputs" {
		for _, ID := range IDs {
			for pos, oldID := range c.ManagerinputIDs {
				if ID == oldID {
					c.ManagerinputIDs = append(c.ManagerinputIDs[:pos], c.ManagerinputIDs[pos+1:]...)
				}
			}
		}
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (u *PaperReport) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	r := make(map[string]interface{})
	var ids []bson.ObjectId
	for k, v := range parameters {
		switch k {
		case "ids":
			for i := 0; i < len(v); i++ {
				ids = append(ids, bson.ObjectIdHex(v[i]))
			}
			r["$in"] = ids
			rst["_id"] = r
		}
	}
	return rst
}
