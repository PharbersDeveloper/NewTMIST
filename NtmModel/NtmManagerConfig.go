package NtmModel

import "gopkg.in/mgo.v2/bson"

// ManagerConfig Info
type ManagerConfig struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	ManagerKPI     float64 `json:"manager-kpi" bson:"manager-kpi"`
	ManagerTime    float64 `json:"manager-time" bson:"manager-time"`
	VisitTotalTime float64 `json:"visit-total-time" bson:"visit-total-time"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ManagerConfig) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ManagerConfig) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *ManagerConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
