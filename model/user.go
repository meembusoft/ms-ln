package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct
type User struct {
	ID       primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Phone    string             `json:"phone" bson:"phone,omitempty"`
	Password string             `json:"password" binding:"required,min=8" bson:"password"`
}
