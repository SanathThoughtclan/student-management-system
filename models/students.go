package models

import "time"

type Student struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	FirstName string    `json:"first_name" bson:"first_name"`
	LastName  string    `json:"last_name" bson:"last_name"`
	Course    string    `json:"course" bson:"course"`
	Grade     string    `json:"grade" bson:"grade"`
	CreatedBy string    `json:"created_by" bson:"created_by"`
	CreatedOn time.Time `json:"created_on" bson:"created_on"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
	UpdatedOn time.Time `json:"updated_on" bson:"updated_on"`
}
