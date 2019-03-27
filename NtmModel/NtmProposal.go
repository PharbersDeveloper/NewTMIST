package NtmModel

import "gopkg.in/mgo.v2/bson"

// Proposal Info
type Proposal struct {
	ID         string        `json:"-"`
	Id_        bson.ObjectId `json:"-" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Describe   string        `json:"describe" bson:"describe"`
	TotalPhase int           `json:"total-phase" bson:"total-phase"`
	InputIDs   []string      `json:"input-ids" bson:"input-ids"`
	ReportIDs  []string      `json:"report-ids" bson:"report-ids"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Proposal) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Proposal) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Proposal) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
