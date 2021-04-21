package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RatingItem struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Ip     string             `bson:"ip,omitempty"`
	Rating int                `bson:"rating,omitempty"`
}
