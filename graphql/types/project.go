package types

import (
	"time"
)

// Project is a type to transform project data to database
type Project struct {
	ID          uint64    `json:"id" bson:"_id,omitempty"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	Year        int       `bson:"year" json:"year"`
	StartedTime time.Time `bson:"startedTime" json:"startedTime "`
	EndedTime   time.Time `bson:"endedTime" json:"endedTime "`
	AdminID     string    `bson:"adminID" json:"adminID"`
	Logo        string    `bson:"logo" json:"logo"`
	Images      []string  `bson:"image" json:"image"`
	MemberIDs   []uint64  `bson:"memberIDs" json:"memberIDs"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
}
