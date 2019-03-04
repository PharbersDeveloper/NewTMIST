package NtmModel

import "gopkg.in/mgo.v2/bson"

// Scenario Info
type Scenario struct {
	ID   		string        `json:"-"`
	Id_  		bson.ObjectId `json:"-" bson:"_id"`
	ProposalID	string        `json:"proposal-id" bson:"proposal-id"`
	Phase  		int       	  `json:"phase" bson:"phase"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Scenario) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Scenario) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Scenario) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}