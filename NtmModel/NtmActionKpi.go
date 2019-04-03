package NtmModel

import "gopkg.in/mgo.v2/bson"

type ActionKpi struct {
	ID						string        `json:"-"`
	Id_						bson.ObjectId `json:"-" bson:"_id"`
	ScenarioID				string	`json:"-" bson:"scenario-id"`
	RepresentativeID		string	`json:"representative-id" bson:"representative-id"`
	TargetNumber			float64	`json:"target-number" bson:"target-number"`
	TargetCoverage			float64	`json:"target-coverage" bson:"target-coverage"`
	HighLevelFrequency		float64	`json:"high-level-frequency" bson:"high-level-frequency"`
	MiddleLevelFrequency	float64	`json:"middle-level-frequency" bson:"middle-level-frequency"`
	LowLevelFrequency		float64	`json:"low-level-frequency" bson:"low-level-frequency"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a ActionKpi) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *ActionKpi) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *ActionKpi) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "ids":
			r := make(map[string]interface{})
			var ids []bson.ObjectId
			for i := 0; i < len(v); i++ {
				ids = append(ids, bson.ObjectIdHex(v[i]))
			}
			r["$in"] = ids
			rst["_id"] = r
		case "scenario-id":
			rst[k] = v[0]
		}
	}
	return rst
}
