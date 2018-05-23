package types

import (
	"time"
)

// Group is a type to transform group data to database
type Group struct {
	ID          uint64    `json:"id" bson:"_id,omitempty"`
	GroupName   string    `bson:"groupName" json:"groupName"`
	StartYear   int       `bson:"startYear" json:"startYear"`
	EndYear     int       `bson:"endYear" json:"endYear"`
	LeaderIDs   []uint64  `bson:"leaderIDs" json:"leaderIDs"`
	FromGroupID uint64    `bson:"fromGroupID" json:"fromGroupID"`
	ToGroupIDs  []uint64  `bson:"toGroupIDs" json:"toGroupIDs"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
	MemberIDs   []uint64  `bson:"memberIDs" json:"memberIDs"`
}
