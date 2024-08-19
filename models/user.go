package models

import "time"

type User struct {
	UserID    string    `bson:"user_id"`
	Username  string    `json:"username"`
	Password  string    `bson:"password"`
	CreatedOn time.Time `bson:"created_on"`
}
