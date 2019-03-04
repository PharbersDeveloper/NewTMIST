package NtmModel

import "gopkg.in/mgo.v2/bson"

type RepresentativeInput struct {
	ID                       string        `json:"-"`
	Id_                      bson.ObjectId `json:"-" bson:"_id"`
	RepresentativeId         string        `json:"representative-id" bson:"representative-id"`
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
func (c RepresentativeInput) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *RepresentativeInput) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *RepresentativeInput) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
