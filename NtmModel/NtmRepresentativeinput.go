package NtmModel

import "gopkg.in/mgo.v2/bson"

type Representativeinput struct {
	ID                       string        `json:"-"`
	Id_                      bson.ObjectId `json:"-" bson:"_id"`
	ResourceConfigId         string        `json:"resource-config-id" bson:"resource-config-id"`
	ProductKnowledgeTraining float64       `json:"product-knowledge-training" bson:"product-knowledge-training"`
	SalesAbilityTraining     float64       `json:"sales-ability-training" bson:"sales-ability-training"`
	RegionTraining           float64       `json:"region-training" bson:"region-training"`
	PerformanceTraining      float64       `json:"performance-training" bson:"performance-training"`
	VocationalDevelopment    float64       `json:"vocational-development" bson:"vocational-development"`
	AssistAccessTime         float64       `json:"assist-access-time" bson:"assist-access-time"`
	TeamMeeting              float64       `json:"team-meeting" bson:"team-meeting"`
	AbilityCoach             float64       `json:"ability-coach" bson:"ability-coach"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Representativeinput) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Representativeinput) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Representativeinput) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
