package api

import (
	"fmt"
	"net/http"
	"net/rpc"

	"github.com/gorilla/mux"

	"github.com/darahayes/go-boom"
	"github.com/globalsign/mgo/bson"
)

type ResponseCSV struct {
	ID      bson.ObjectId
	Serial  string
	Name    string
	Details string
}
type ResponseComent struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Comment string
	CsvID   bson.ObjectId
}
type ResponseCSVWithComments struct {
	ID       bson.ObjectId
	Serial   string
	Name     string
	Details  string
	Comments []ResponseComent
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
func ReadCSVWithComments(dumpService *rpc.Client) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("HERE")
		v := mux.Vars(r)
		var fc ResponseCSVWithComments
		err := dumpService.Call("Service.ReadCSVWithComments", v["csvid"], &fc)
		if err != nil {
			boom.BadGateway(w, err.Error())
		}
		Respond(w, fc)

	}

}
