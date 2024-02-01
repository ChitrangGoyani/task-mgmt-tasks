package models

import "time"

type Tasks struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"` // bson is for mongodb
	UserID      string    `json:"userID"`
	Priority    string    `json:"priority"`
	Content     string    `json:"content"`
	Time        time.Time `json:"time"`
	Done        bool      `json:"done"`
	UpdatedTime time.Time `json:"updatedTime"`
}
