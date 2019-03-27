package NtmModel

import "gopkg.in/mgo.v2/bson"

type Managerinput struct {
	ID                   string        `json:"-"`
	Id_                  bson.ObjectId `json:"-" bson:"_id"`
	StrategyAnalysisTime float64       `json:"strategy-analysis-time" bson:"strategy-analysis-time"`
	AdminWorkTime        float64       `json:"admin-work-time" bson:"admin-work-time"`
	ClientManagementTime float64       `json:"client-management-time" bson:"client-management-time"`
	KPIAnalysisTime      float64       `json:"KPI-analysis-time" bson:"KPI-analysis-time"`
	TeamMeetingTime      float64       `json:"team-meeting-time" bson:"team-meeting-time"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Managerinput) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Managerinput) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Managerinput) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
