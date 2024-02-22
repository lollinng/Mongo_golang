package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
omitemptignore empty fields
y and validate:required  make the field required
*/

// json will send as id but in mondb in bson it saved as _id
type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required"`
}
