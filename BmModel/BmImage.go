package BmModel

import "gopkg.in/mgo.v2/bson"

// Image is the Image that a user consumes in order to get fat and happy
type Image struct {
	ID   string        `json:"-"`
	Id_  bson.ObjectId `json:"-" bson:"_id"`
	Img  string        `json:"img" bson:"img"`
	Tag  string        `json:"tag" bson:"tag"`
	Flag float64       `json:"flag" bson:"flag"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Image) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Image) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Image) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
