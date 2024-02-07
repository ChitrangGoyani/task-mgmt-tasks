package models

type Tasks struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"` // bson is for mongodb
	UserID      string `json:"userID"`
	Priority    string `json:"priority"`
	Content     string `json:"content"`
	Time        string `json:"time"`
	Done        bool   `json:"done"`
	UpdatedTime string `json:"updatedTime"`
}
