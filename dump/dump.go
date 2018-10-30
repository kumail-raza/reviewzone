package dump

import (
	"github.com/globalsign/mgo/bson"
)

const (
	collection = "csv"
)

//Dumper Dumper
type Dumper struct{}

//DumpCSV DumpCSV
func (d *Dumper) dumpCSV(csvs [][]string) ([]string, error) {

	session, err := d.connectDB()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(dbName).C(collection)
	var ids []string
	for _, row := range csvs {
		f := Format{Serial: row[0], Name: row[1], ID: bson.NewObjectId()}
		err := c.Insert(f)
		if err != nil {
			return nil, err
		}
		ids = append(ids, f.ID.Hex())
	}
	return ids, nil

}

//ReadCSVFromDB ReadCSVFromDB
func (d *Dumper) readCSVFromDB() ([]Format, error) {
	s, err := d.connectDB()
	if err != nil {
		return nil, err
	}
	defer s.Close()
	c := s.DB(dbName).C(collection)

	var f []Format
	err = c.Find(nil).All(&f)
	if err != nil {
		return nil, err
	}
	return f, nil

}

func (d *Dumper) readOneCSV(csvID string) (Format, error) {

	f := Format{}
	s, err := d.connectDB()
	if err != nil {
		return f, err
	}
	defer s.Close()
	c := s.DB(dbName).C(collection)
	err = c.FindId(bson.ObjectIdHex(csvID)).One(&f)
	if err != nil {
		return f, err
	}
	return f, nil
}
