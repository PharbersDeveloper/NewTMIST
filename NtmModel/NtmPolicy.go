package NtmModel

import "gopkg.in/mgo.v2/bson"

type Policy struct {
	ID			string        `json:"-"`
	Id_			bson.ObjectId `json:"-" bson:"_id"`
	Name		string        `json:"name" bson:"name"`
	Describe	string    	  `json:"describe" bson:"describe"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Policy) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Policy) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Policy) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
