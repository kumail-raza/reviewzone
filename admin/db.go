package admin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/globalsign/mgo"
)

var (
	dbName             = "reviewzone"
	dbUserName         = getDBUser()
	dbPassword         = getDBPassword()
	dbConnectionString = fmt.Sprintf("mongodb:27017")
)

func getDBUser() string {
	user := os.Getenv("MONGODB_USER")
	if len(user) == 0 {
		return "user"
	}
	return user
}

func getDBPassword() string {
	pwd := os.Getenv("MONGODB_PASS")
	if len(pwd) == 0 {
		return "pass"
	}
	return pwd
}

//ConnectDB ConnectDB
func (a *Admin) ConnectDB() (*mgo.Session, error) {

	sessionCh := make(chan *mgo.Session)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		fmt.Println(dbConnectionString)
		session, err := mgo.Dial(dbConnectionString + "/" + dbName)
		if err != nil {
			errCh <- err
		}
		sessionCh <- session

	}()
	select {
	case s := <-sessionCh:
		return s, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		return nil, errors.New("Timeout! Coudn't make database connection")
	}
}
