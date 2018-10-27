package dump

import "github.com/globalsign/mgo/bson"

//Format Format
type Format struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Serial  string
	Name    string
	Details string
}

//Comment Comment
type Comment struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Comment string
	CsvID   bson.ObjectId
}

//FormatWithComments FormatWithComments
type FormatWithComments struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Serial   string
	Name     string
	Details  string
	Comments []Comment
}
