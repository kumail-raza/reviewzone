package admin

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/globalsign/mgo/bson"
)

type Admin struct{}

type Format struct {
	Serial   string
	Name     string
	Details  string
	ID       bson.ObjectId
	Comments []string
}

//GetCSVs GetCSVs
func (a *Admin) GetCSVs() ([][]string, error) {

	s, err := a.ConnectDB()
	if err != nil {
		return nil, err
	}
	c := s.DB(dbName).C("csv")
	var f []Format

	err = c.Find(nil).All(&f)
	if err != nil {
		return nil, err
	}
	spew.Dump(f)

	var csvs [][]string
	for _, v := range f {
		csvs = append(csvs, []string{v.Serial, v.ID.Hex()})
	}
	return csvs, nil
}
