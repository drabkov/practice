package domain

type PublishingHouse struct {
	Name   string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Adress string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
