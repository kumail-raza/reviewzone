package api

import (
	"net/http"
	"net/rpc"
	"sync"

	boom "github.com/darahayes/go-boom"
	"github.com/gorilla/mux"

	"github.com/globalsign/mgo/bson"
)

type ResponseCSV struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Serial  string
	Name    string
	Details string
}
type ResponseComment struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Comment string
	CsvID   bson.ObjectId
}

type ResponseCSVWithComments struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Serial   string
	Name     string
	Details  string
	Comments []ResponseComment
}

//DumpCSV DumpCSV
func DumpCSV(readerService, dumpService *rpc.Client) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var csvs [][]string
		err := readerService.Call("Service.ReadCSVFile", "text.csv", &csvs)
		if err != nil {
			panic(err)
		}

		var csvIDs []string
		err = dumpService.Call("Service.DumpCSV", csvs, &csvIDs)
		if err != nil {
			boom.BadGateway(w, err.Error())
			return
		}
		Respond(w, csvIDs)
		return
	}
}

//ReadCSV ReadCSV
func ReadCSV(readerService *rpc.Client) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var csvs [][]string

		err := readerService.Call("Service.ReadCSVFile", "text.csv", &csvs)
		if err != nil {
			panic(err)
		}
		Respond(w, struct {
			CSVS [][]string
		}{csvs})
	}
}

//ReadCSVFromDB ReadCSVFromDB
func ReadCSVFromDB(dumpService *rpc.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var csvs []ResponseCSV
		err := dumpService.Call("Service.ReadCSVFromDB", "", &csvs)
		if err != nil {
			boom.BadGateway(w, err.Error())
		}
		Respond(w, csvs)
	}
}

//ReadCSVWithComments ReadCSVWithComments
func ReadCSVWithComments(dumpService, commentService *rpc.Client) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		csvID := mux.Vars(r)["csvid"]

		var wg sync.WaitGroup

		var csv ResponseCSV
		var comments []ResponseComment
		var rpcError error

		wg.Add(2)
		go func() {
			err := dumpService.Call("Service.ReadOneCSV", csvID, &csv)
			if err != nil {
				rpcError = err
			}
			wg.Done()

		}()

		go func() {
			err := commentService.Call("Service.GetComments", csvID, &comments)
			if err != nil {
				rpcError = err
			}
			wg.Done()
		}()

		wg.Wait()

		if rpcError != nil {
			boom.BadGateway(w, rpcError.Error())
			return
		}

		Respond(w, ResponseCSVWithComments{
			ID:       csv.ID,
			Serial:   csv.Serial,
			Name:     csv.Name,
			Comments: comments,
		})

	}

}
