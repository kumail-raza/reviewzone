package dump

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/globalsign/mgo/bson"
)

//Format Format
type Format struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Serial  string
	Name    string
	Details string
}

var (
	dbName             = "reviewzone"
	dbUserName         = getDBUser()
	dbPassword         = getDBPassword()
	dbConnectionString = fmt.Sprintf("mongodb:27017")
)

//Dumper Dumper
type Dumper struct{}

//DumpCSV DumpCSV
func (d *Dumper) dumpCSV(csvs [][]string) ([]string, error) {

	session, err := d.connectDB()
	if err != nil {
		return nil, err
	}
	c := session.DB(dbName).C("csv")
	var ids []string
	for _, row := range csvs {
		spew.Dump(row)
		f := Format{Serial: row[0], Name: row[1], ID: bson.NewObjectId()}
		err := c.Insert(f)
		if err != nil {
			return nil, err
		}
		ids = append(ids, f.ID.Hex())
	}
	return ids, nil

}
