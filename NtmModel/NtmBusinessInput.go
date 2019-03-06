package NtmModel

import "gopkg.in/mgo.v2/bson"

type BusinessInput struct {
	ID               string        `json:"-"`
	Id_              bson.ObjectId `json:"-" bson:"_id"`
	HospitalId       string        `json:"hospital-id" bson:"hospital-id"`
	RepresentativeId string        `json:"representative-id" bson:"representative-id"`
	SalesTarget      float64       `json:"sales-target" bson:"sales-target"`
	Budget           float64       `json:"budget" bson:"budget"`
	MeetingPlaces    float64       `json:"meeting-places" bson:"meeting-places"`
	VisitTime        float64       `json:"vivit-time" bson:"vivit-time"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c BusinessInput) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *BusinessInput) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *BusinessInput) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
