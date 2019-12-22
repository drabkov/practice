package domain

type PublishingHouse struct {
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Adress string `json:"address,omitempty" bson:"address,omitempty"`
}
