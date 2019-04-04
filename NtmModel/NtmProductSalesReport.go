package NtmModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// ProductSalesReport Info
type ProductSalesReport struct {
	ID         		string        `json:"-"`
	Id_        		bson.ObjectId `json:"-" bson:"_id"`
	GoodsConfigID	string  `json:"-" bson:"goods-config-id"`

	GoodsConfig 	*GoodsConfig `json:"-"`

	ProductName		string `json:"product-name" bson:"product-name"`

	Sales			float64 `json:"sales" bson:"sales"`
	SalesQuota 		float64	`json:"sales-quota" bson:"sales-quota"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interfac
func (c ProductSalesReport) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ProductSalesReport) SetID(id string) error {
	c.ID = id
	return nil
}


// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u ProductSalesReport) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "goodsConfigs",
			Name: "goodsConfig",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u ProductSalesReport) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if u.GoodsConfigID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.GoodsConfigID,
			Type: "goodsConfigs",
			Name: "goodsConfig",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u ProductSalesReport) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.GoodsConfigID != "" && u.GoodsConfig != nil {
		result = append(result, u.GoodsConfig)
	}

	return result
}

func (u *ProductSalesReport) SetToOneReferenceID(name, ID string) error {
	if name == "goodsConfig" {
		u.GoodsConfigID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *ProductSalesReport) GetConditionsBsonM(parameters map[string][]string) bson.M {
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