package types

// Company is a type to transform company data to database
type Company struct {
	ID          uint64   `json:"id" bson:"_id,omitempty"`
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Logo        string   `bson:"logo" json:"logo"`
	CreatorID   uint64   `bson:"creatorID" json:"creatorID"`
	Images      []string `bson:"images" json:"images"`
	MemberIDs   []uint64 `bson:"memberIDs" json:"memberIDs"`
}
