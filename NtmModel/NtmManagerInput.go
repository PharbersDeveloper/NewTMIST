package NtmModel

import "gopkg.in/mgo.v2/bson"

type ManagerInput struct {
	ID           string        `json:"-"`
	Id_          bson.ObjectId `json:"-" bson:"_id"`
	StrategyTime float64       `json:"strategy-time" bson:"strategy-time"`
	AdminTime    float64       `json:"admin-time" bson:"admin-time"`
	KPI          float64       `json:"KPI" bson:"KPI"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ManagerInput) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ManagerInput) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *ManagerInput) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
