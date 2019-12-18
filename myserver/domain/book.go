package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Book struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn            string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Title           string             `json:"title" bson:"title,omitempty"`
	Author          *Author            `json:"author" bson:"author,omitempty"`
	PublishingHouse *PublishingHouse   `json:"publishinghouse" bson:"author,omitempty"`
}
