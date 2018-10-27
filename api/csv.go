package api

import (
	"net/http"
	"net/rpc"

	"github.com/darahayes/go-boom"
)

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
