package BmModel

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmLoginSucceed struct {
	ID  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	// ApplicantID string     `json:"-" bson:"applicant-id"`
	// Applicant   *Applicant `json:"-"`
	Applicant Applicant `json:"Applicant" jsonapi:"relationships"`
	Token     string    `json:"token"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmLoginSucceed) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmLoginSucceed) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmLoginSucceed) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmLoginSucceed) QueryId() string {
	return bd.ID
}

func (bd *BmLoginSucceed) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmLoginSucceed) SetId(id string) {
	bd.ID = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmLoginSucceed) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Applicant":
		bd.Applicant = v.(Applicant)
	}
	return bd
}

func (bd BmLoginSucceed) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmLoginSucceed) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmLoginSucceed) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmLoginSucceed) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
