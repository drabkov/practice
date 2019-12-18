package domain

type PublishingHouse struct {
	Name   string `json:"name,omitempty" bson:"firstname,omitempty"`
	Adress string `json:"address,omitempty" bson:"lastname,omitempty"`
}
