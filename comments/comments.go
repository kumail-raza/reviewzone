package comments

import "github.com/globalsign/mgo/bson"

type Comments struct {
	Comments []string
}

//Format Format
type Format struct {
	Serial   string
	Name     string
	Details  string
	ID       bson.ObjectId
	Comments []string
}

//AddOnCSV AddOnCSV
func (c *Comments) AddOnCSV(csvID string) error {

	session, err := c.ConnectDB()
	defer session.Close()
	if err != nil {
		return err
	}

	who := bson.M{"_id": bson.ObjectIdHex(csvID)}
	pushToArray := bson.M{"$push": bson.M{"comments": c.Comments}}
	if err != nil {
		return err
	}
	err = session.DB(DBName).C("csv").Update(who, pushToArray)
	if err != nil {
		return err
	}
	return nil
}
