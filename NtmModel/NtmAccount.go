package NtmModel

import (
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Nickname string `json:"nickname" bson:"nickname"`
	Phone    string `json:"phone" bson:"phone"`

	Token    string `json:"token"`
}

func (u *Account) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
