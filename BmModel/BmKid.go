package BmModel

import (
	"gopkg.in/mgo.v2/bson"
)

// Kid is the Kid that a user consumes in order to get fat and happy
type Kid struct {
	ID           string        `json:"id"`
	Id_          bson.ObjectId `json:"-" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	NickName     string        `json:"nickname" bson:"nickname"`
	Gender       float64       `json:"gender" bson:"gender"`
	Dob          float64       `json:"dob" bson:"dob"`
	GuardianRole string        `json:"guardian-role" bson:"guardian-role"`

	ApplicantID string    `json:"applicant-id" bson:"applicant-id"`
	Applicant   Applicant `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Kid) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Kid) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Kid) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "applicant-id":
			rst[k] = v[0]
		}
	}

	return rst
}
