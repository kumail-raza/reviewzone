package dump

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/rpc"
	"time"

	"github.com/globalsign/mgo/bson"
)

var (
	dbName             = "reviewzone"
	dbUserName         = getDBUser()
	dbPassword         = getDBPassword()
	dbConnectionString = fmt.Sprintf("mongo:27017")
	collection         = "csv"
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

func (d *Dumper) readWithComments(csvID string) (FormatWithComments, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	s, err := d.connectDB()
	var f FormatWithComments
	if err != nil {
		return f, err
	}
	defer s.Close()
	c := s.DB(dbName).C(collection)

	serviceCH := make(chan *rpc.Client)
	errCh := make(chan error)

	go func() {
		service, err := rpc.DialHTTP("tcp", "comments:3000")
		if err != nil {
			errCh <- err
			log.Fatal(err)
		}
		serviceCH <- service

	}()

	fmt.Println("cool here", csvID)
	err = c.FindId(bson.ObjectIdHex(csvID)).One(&f)
	if err != nil {
		return f, err
	}
	fmt.Println(f.Serial)
	var comments []Comment
	select {
	case s := <-serviceCH:
		err = s.Call("Service.GetComments", f.ID, &comments)
		if err != nil {
			fmt.Println("Coudn't get comments")
		}
		f.Comments = comments
		return f, nil

	case err := <-errCh:
		return f, err
	case <-ctx.Done():
		return f, errors.New("dump service timed out")
	}

}
