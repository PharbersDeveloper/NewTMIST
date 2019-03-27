package NtmModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type Paperinput struct {
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
func (c Paperinput) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Paperinput) SetID(id string) error {
	c.ID = id
	return nil
}

func (c Paperinput) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "businessinputs",
			Name: "businessinputs",
		},
		{
			Type: "representativeinputs",
			Name: "representativeinputs",
		},
		{
			Type: "managerinputs",
			Name: "managerinputs",
		},
	}
}

func (c Paperinput) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range c.BusinessinputIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "businessinputs",
			Name: "businessinputs",
		})
	}

	for _, kID := range c.RepresentativeinputIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "representativeinputs",
			Name: "representativeinputs",
		})
	}

	for _, kID := range c.ManagerinputIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "managerinputs",
			Name: "managerinputs",
		})
	}

	return result
}

func (c Paperinput) GetReferencedStructs() []jsonapi.MarshalIdentifier {
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

func (c *Paperinput) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "businessinputs" {
		c.BusinessinputIDs = IDs
		return nil
	} else if name == "representativeinputs" {
		c.RepresentativeinputIDs = IDs
		return nil
	} else if name == "managerinputs" {
		c.ManagerinputIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Paperinput) AddToManyIDs(name string, IDs []string) error {
	if name == "businessinputs" {
		c.BusinessinputIDs = append(c.BusinessinputIDs, IDs...)
		return nil
	} else if name == "representativeinputs" {
		c.RepresentativeinputIDs = append(c.RepresentativeinputIDs, IDs...)
		return nil
	} else if name == "managerinputs" {
		c.ManagerinputIDs = append(c.ManagerinputIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Paperinput) DeleteToManyIDs(name string, IDs []string) error {
	if name == "businessinputs" {
		for _, ID := range IDs {
			for pos, oldID := range c.BusinessinputIDs {
				if ID == oldID {
					c.BusinessinputIDs = append(c.BusinessinputIDs[:pos], c.BusinessinputIDs[pos+1:]...)
				}
			}
		}
	} else if name == "representativeinputs" {
		for _, ID := range IDs {
			for pos, oldID := range c.RepresentativeinputIDs {
				if ID == oldID {
					c.RepresentativeinputIDs = append(c.RepresentativeinputIDs[:pos], c.RepresentativeinputIDs[pos+1:]...)
				}
			}
		}
	} else if name == "managerinputs" {
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

func (u *Paperinput) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
