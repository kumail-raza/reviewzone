package comments

import "github.com/globalsign/mgo/bson"

//Comment Comment
type Comment struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Comment string
	CsvID   bson.ObjectId
}

type Format struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Serial  string
	Name    string
	Details string
}

type FormatWithComments struct {
	Format
	Comments []Comment
}
