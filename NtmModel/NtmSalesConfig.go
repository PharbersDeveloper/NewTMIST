package NtmModel

import (
	"gopkg.in/mgo.v2/bson"
)

// SalesConfig Info
type SalesConfig struct {
	ID         		string        `json:"-"`
	Id_        		bson.ObjectId `json:"-" bson:"_id"`
	ScenarioId 		string        `json:"scenario-id" bson:"scenario-id"`
	DestID     		string        `json:"dest-id" bson:"dest-id"`
	GoodsIDs		[]string 	  `json:"goods-ids" bson:"goods-ids"`
	AccessStatus   	string  `json:"access-status" bson:"access-status"`
	LastYearSales 	float64 `json:"last-year-sales" bson:"last-year-sales"`
	Potential     	float64 `json:"potential" bson:"potential"`
	//ReportID		string	`json"-" bson:"report-id"` //预留字段
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c SalesConfig) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *SalesConfig) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *SalesConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
