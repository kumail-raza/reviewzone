package admin

import (
	"github.com/globalsign/mgo/bson"
)

type Admin struct{}

type Format struct {
	ID      bson.ObjectId
	Serial  string
	Name    string
	Details string
}

const (
	DUMPS_PORT    = ":4000"
	COMMENTS_PORT = ":3000"
)

//GetCSVs GetCSVs
func (a *Admin) GetCSVs() ([][]string, error) {

	// s, err := a.ConnectDB()
	// if err != nil {
	// 	return nil, err
	// }
	// c := s.DB(dbName).C("csv")
	// var f []Format

	// err = c.Find(nil).All(&f)
	// if err != nil {
	// 	return nil, err
	// }
	// spew.Dump(f)

	// var csvs [][]string
	// for _, v := range f {
	// 	csvs = append(csvs, []string{v.Serial, v.ID.Hex()})
	// }
	return csvs, nil
}
