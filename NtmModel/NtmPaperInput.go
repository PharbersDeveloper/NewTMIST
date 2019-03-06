package NtmModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type PaperInput struct {
	ID      string        `json:"-"`
	Id_     bson.ObjectId `json:"-" bson:"_id"`
	PaperId string        `json:"paper-id" bson:"paper-id"`
	Phase   float64       `json:"phase" bson:"phase"`

	BusinessInputIDs []string         `json:"-" bson:"business-input-ids"`
	BusinessInputs   []*BusinessInput `json:"-"`

	RepresentativeInputIDs []string               `json:"-" bson:"representative-input-ids"`
	RepresentativeInputs   []*RepresentativeInput `json:"-"`

	ManagerInputIDs []string        `json:"-" bson:"manager-input-ids"`
	ManagerInputs   []*ManagerInput `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c PaperInput) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *PaperInput) SetID(id string) error {
	c.ID = id
	return nil
}

func (c PaperInput) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "businessInputs",
			Name: "businessInputs",
		},
		{
			Type: "representativeInputs",
			Name: "representativeInputs",
		},
		{
			Type: "managerInputs",
			Name: "managerInputs",
		},
	}
}

func (c PaperInput) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range c.BusinessInputIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "businessInputs",
			Name: "businessInputs",
		})
	}

	for _, kID := range c.RepresentativeInputIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "representativeInputs",
			Name: "representativeInputs",
		})
	}

	for _, kID := range c.ManagerInputIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "managerInputs",
			Name: "managerInputs",
		})
	}

	return result
}

func (c PaperInput) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range c.BusinessInputs {
		result = append(result, c.BusinessInputs[key])
	}

	for key := range c.RepresentativeInputs {
		result = append(result, c.RepresentativeInputs[key])
	}

	for key := range c.ManagerInputs {
		result = append(result, c.ManagerInputs[key])
	}

	return result
}

func (c *PaperInput) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "businessInputs" {
		c.BusinessInputIDs = IDs
		return nil
	} else if name == "representativeInputs" {
		c.RepresentativeInputIDs = IDs
		return nil
	} else if name == "managerInputs" {
		c.ManagerInputIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *PaperInput) AddToManyIDs(name string, IDs []string) error {
	if name == "businessInputs" {
		c.BusinessInputIDs = append(c.BusinessInputIDs, IDs...)
		return nil
	} else if name == "representativeInputs" {
		c.RepresentativeInputIDs = append(c.RepresentativeInputIDs, IDs...)
		return nil
	} else if name == "managerInputs" {
		c.ManagerInputIDs = append(c.ManagerInputIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *PaperInput) DeleteToManyIDs(name string, IDs []string) error {
	if name == "businessInputs" {
		for _, ID := range IDs {
			for pos, oldID := range c.BusinessInputIDs {
				if ID == oldID {
					c.BusinessInputIDs = append(c.BusinessInputIDs[:pos], c.BusinessInputIDs[pos+1:]...)
				}
			}
		}
	} else if name == "representativeInputs" {
		for _, ID := range IDs {
			for pos, oldID := range c.RepresentativeInputIDs {
				if ID == oldID {
					c.RepresentativeInputIDs = append(c.RepresentativeInputIDs[:pos], c.RepresentativeInputIDs[pos+1:]...)
				}
			}
		}
	} else if name == "managerInputs" {
		for _, ID := range IDs {
			for pos, oldID := range c.ManagerInputIDs {
				if ID == oldID {
					c.ManagerInputIDs = append(c.ManagerInputIDs[:pos], c.ManagerInputIDs[pos+1:]...)
				}
			}
		}
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (u *PaperInput) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
