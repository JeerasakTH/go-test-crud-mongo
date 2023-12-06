package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Address  string             `bson:"address" json:"address"`
}
