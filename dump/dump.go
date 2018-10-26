package dump

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

//Format Format
type Format struct {
	Serial   string
	Name     string
	Details  string
	ID       bson.ObjectId
	Comments []string
}

var (
	dbName             = "test"
	dbUserName         = getDBUser()
	dbPassword         = getDBPassword()
	dbConnectionString = fmt.Sprintf("localhost:27017")
)

//Dumper Dumper
type Dumper struct{}

//DumpCSV DumpCSV
func (d *Dumper) DumpCSV(csvs [][]string) ([]string, error) {

	session, err := d.ConnectDB()
	if err != nil {
		return nil, err
	}
	c := session.DB(dbName).C("csv")
	var ids []string
	for _, row := range csvs {
		f := Format{Serial: row[0], ID: bson.NewObjectId()}
		err := c.Insert(f)
		if err != nil {
			return nil, err
		}
		ids = append(ids, f.ID.Hex())
	}
	return ids, nil

}
