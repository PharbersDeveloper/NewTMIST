package NtmModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type ProductConfig struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	ProductType     string `json:"product-type" bson:"product-type"`
	PriceType       string  `json:"price-type" bson:"price-type"`
	Price           float64  `json:"price" bson:"price"`
	LifeCycle       string  `json:"life-cycle" bson:"life-cycle"`
	LaunchTime      float64  `json:"launch-time" bson:"launch-time"`
	ProductCategory string  `json:"product-category" bson:"product-category"`
	TreatmentArea   string  `json:"treatment-area" bson:"treatment-area"`
	ProductFeature  string  `json:"product-feature" bson:"product-feature"`

	ProductID string `json:"-" bson:"product-id"`
	Product   Product `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u ProductConfig) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *ProductConfig) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u ProductConfig) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "products",
			Name: "product",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u ProductConfig) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if u.ProductID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.ProductID,
			Type: "products",
			Name: "product",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u ProductConfig) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.ProductID != "" {
		result = append(result, u.Product)
	}

	return result
}

func (u *ProductConfig) SetToOneReferenceID(name, ID string) error {
	if name == "product" {
		u.ProductID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *ProductConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
