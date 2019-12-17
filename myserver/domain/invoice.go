package domain

import "gopkg.in/mgo.v2/bson"

type Invoices []Invoice

type Invoice struct {
	ID       bson.ObjectId `bson:"_id"`
	Partner  *Partner      `json:"partner"`
	Products []Product     `json:"products"`
}
