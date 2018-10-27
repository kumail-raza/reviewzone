package comments

import (
	"errors"
	"fmt"

	"github.com/globalsign/mgo/bson"
)

//Comment Comment
type Comment struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Comment string
	CsvID   bson.ObjectId
}

const (
	collection = "comments"
)

//AddOnCSV AddOnCSV
func (c *Comment) addOnCSV(csvID string, comments ...string) ([]string, error) {

	session, err := c.ConnectDB()
	defer session.Close()
	if err != nil {
		return nil, err
	}
	col := session.DB(dBName).C(collection)
	var commentIds []string
	doneCh := make(chan string, len(comments))
	errCh := make(chan error)

	for i := range comments {
		go func(comment string) {
			cmt := Comment{bson.NewObjectId(), comment, bson.ObjectIdHex(csvID)}
			err = col.Insert(cmt)
			if err != nil {
				errCh <- err
			}
			doneCh <- cmt.ID.Hex()
		}(comments[i])
	}

	for i := 0; i < len(comments); i++ {
		select {
		case cid := <-doneCh:
			commentIds = append(commentIds, cid)
		case err = <-errCh:
			return nil, err
		}

	}

	return commentIds, nil

}

//GetComments GetComments
func (c *Comment) getComments(csvID string) ([]Comment, error) {

	if !bson.IsObjectIdHex(csvID) {
		return nil, errors.New("Invalid object id")
	}
	s, err := c.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer s.Close()

	var comments []Comment
	fmt.Println("incoming", csvID)

	err = s.DB(dBName).C(collection).Find(bson.M{"csvid": bson.ObjectIdHex(csvID)}).All(&comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
